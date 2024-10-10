package http

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/dmdhrumilmistry/fasthttpclient/client"
	"github.com/li-jin-gou/http2curl"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/time/rate"
)

type ThrottledTransport struct {
	roundTripperWrap http.RoundTripper
	ratelimiter      *rate.Limiter
}

type CustomClient struct {
	Requests  []*client.Request
	Responses []*client.ConcurrentResponse
	Client    *http.Client
}

func GetResponseHeaders(resp *http.Response) map[string]string {
	headers := make(map[string]string)
	for header, value := range resp.Header {
		headers[header] = value[0]
	}
	return headers
}

func SetRequestBody(body interface{}, req *http.Request) error {
	if body == nil {
		return nil
	}

	bodyBytes, ok := body.([]byte)
	if !ok {
		return fmt.Errorf("body only supports []byte type")
	} else {
		req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	}
	return nil
}

func (c *CustomClient) Do(uri string, method string, queryParams any, headers any, reqBody any) (*client.Response, error) {
	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return nil, err
	}

	queryParamsMap, ok := queryParams.(map[string]string)
	if !ok && queryParams != nil {
		return nil, fmt.Errorf("queryParams must be a map[string]string")
	} else {
		for key, value := range queryParamsMap {
			req.Header.Add(key, value)
		}
	}

	err = SetRequestBody(reqBody, req)
	if err != nil {
		return nil, err
	}

	curlCmd := "error generating curl command"
	curlCmdObj, err := http2curl.GetCurlCommand(req)
	if err == nil {
		curlCmd = curlCmdObj.String()
	}

	now := time.Now()
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	elapsed := time.Since(now)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	respHeaders := GetResponseHeaders(resp)
	statusCode := resp.StatusCode

	return &client.Response{StatusCode: statusCode, Body: body, Headers: respHeaders, CurlCommand: curlCmd, TimeElapsed: elapsed}, nil
}

// RoundTrip method implements the rate limiting logic for HTTP requests
func (c *ThrottledTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	err := c.ratelimiter.Wait(r.Context()) // This is a blocking call. Honors the rate limit
	if err != nil {
		return nil, err
	}
	return c.roundTripperWrap.RoundTrip(r)
}

// NewThrottledTransport wraps the provided transport with a rate limiter
func NewThrottledTransport(limitPeriod time.Duration, requestCount int, transportWrap http.RoundTripper) http.RoundTripper {
	return &ThrottledTransport{
		roundTripperWrap: transportWrap,
		ratelimiter:      rate.NewLimiter(rate.Every(limitPeriod), requestCount),
	}
}

func NewConfigHttp2(requestsPerSecond *int, skipTlsVerification *bool, proxy *string) *CustomClient {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: *skipTlsVerification,
		MinVersion:         tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_AES_128_GCM_SHA256,       // TLS 1.3
			tls.TLS_AES_256_GCM_SHA384,       // TLS 1.3
			tls.TLS_CHACHA20_POLY1305_SHA256, // TLS 1.3
		},
		PreferServerCipherSuites: true,
	}

	var transport *http.Transport
	if *proxy != "" {
		proxy_url, _ := url.Parse(*proxy)
		transport = &http.Transport{
			TLSClientConfig: tlsConfig,
			Proxy:           http.ProxyURL(proxy_url),
		}
	} else {
		transport = &http.Transport{
			TLSClientConfig: tlsConfig,
		}
	}

	err := http2.ConfigureTransport(transport)
	if err != nil {
		log.Error().Msgf("failed to configure transport: %v", err)
	}

	rateLimitedTransport := NewThrottledTransport(1*time.Second, *requestsPerSecond, transport)

	client := http.Client{
		Transport: rateLimitedTransport,
	}

	return &CustomClient{
		Client: &client,
	}
}
