package http

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"

	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
)

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
