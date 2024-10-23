package http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dmdhrumilmistry/fasthttpclient/client"
	"github.com/li-jin-gou/http2curl"
	"github.com/valyala/fasthttp"
	"golang.org/x/time/rate"
)

type Config struct {
	HttpClient        *fasthttp.Client
	RequestsPerSecond *int
}

type Http struct {
	Requests  []*client.Request
	Responses []*client.ConcurrentResponse
	Config    *Config
	Client    *client.RateLimitedClient
}

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

	// fmt.Println("Response Protocol: ", resp.Proto)
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
