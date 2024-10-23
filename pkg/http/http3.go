package http

import (
	"crypto/tls"
	"net/http"
	"os"
	"time"

	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
	"github.com/quic-go/quic-go/qlog"
	"github.com/rs/zerolog/log"
)

func NewConfigHttp3(requestsPerSecond *int, skipTlsVerification *bool, proxy *string) *CustomClient {
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

	var transport *http3.Transport
	if *proxy != "" {
		log.Error().Msgf("Cannot use proxy with HTTP/3")
		os.Exit(1)
	} else {
		transport = &http3.Transport{
			TLSClientConfig: tlsConfig,
			QUICConfig: &quic.Config{
				Tracer: qlog.DefaultConnectionTracer,
			},
		}
	}

	rateLimitedTransport := NewThrottledTransport(1*time.Second, *requestsPerSecond, transport)

	client := http.Client{
		Transport: rateLimitedTransport,
	}

	return &CustomClient{
		Client: &client,
	}
}
