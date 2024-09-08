package tgen

import (
	_ "github.com/owasp-offat/offat/pkg/logging"
	"github.com/owasp-offat/offat/pkg/parser"
	"github.com/rs/zerolog/log"
)

// generates very basic SSRF API tests by injecting provided URL
func BasicSsrfTest(ssrfUrl, baseUrl string, docParams []*parser.DocHttpParams, queryParams map[string]string, headers map[string]string, injectionConfig InjectionConfig) []*ApiTest {
	testName := "Basic SSRF Test"

	payloads := []Payload{
		{InjText: ssrfUrl},
	}

	injectionConfig.Payloads = payloads

	tests := injectParamIntoApiTest(baseUrl, docParams, queryParams, headers, testName, injectionConfig)
	log.Info().Msg("Check SSRF server for calls. Vulnerable endpoint path will be available in query param")

	return tests
}
