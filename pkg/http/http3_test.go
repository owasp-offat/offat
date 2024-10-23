package http_test

import (
	"testing"

	c "github.com/dmdhrumilmistry/fasthttpclient/client"
	"github.com/owasp-offat/offat/pkg/http"
	"github.com/valyala/fasthttp"
)

func TestHttp3Client(t *testing.T) {
	// http2 client
	requestsPerSecond := 10
	skipTlsVerification := false
	proxy := ""
	hc := http.NewConfigHttp3(&requestsPerSecond, &skipTlsVerification, &proxy)

	url := "https://cloudflare-quic.com/" // This is a QUIC enabled website
	hc.Requests = append(hc.Requests, c.NewRequest(url, fasthttp.MethodGet, nil, nil, nil))
	hc.Requests = append(hc.Requests, c.NewRequest(url, fasthttp.MethodGet, nil, nil, nil))

	t.Run("Concurrent Requests Test", func(t *testing.T) {
		hc.Responses = c.MakeConcurrentRequests(hc.Requests, hc)

		for _, connResp := range hc.Responses {
			if connResp.Error != nil {
				t.Fatalf("failed to make concurrent request: %v\n", connResp.Error)
			}
		}
	})

}
