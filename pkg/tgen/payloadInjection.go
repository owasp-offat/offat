package tgen

import (
	c "github.com/dmdhrumilmistry/fasthttpclient/client"
	"github.com/owasp-offat/offat/pkg/parser"
	"github.com/owasp-offat/offat/pkg/utils"
	"github.com/rs/zerolog/log"
)

// injects payload in HTTP parser.param based on type/value
// It's being used in `injectParamIntoApiTest` function
func injectParamInParam(params []parser.Param, payload, injectIn string) []parser.Param {
	// inject payload in key
	injectedParams := append(params, parser.Param{
		Name: payload,
		In:   injectIn,
		Type: []string{"string"},
	})

	for _, param := range injectedParams {
		var paramType string
		if len(param.Type) == 0 && param.Value == nil {
			log.Warn().Msgf("skipping payload %s injection for %v since type/value is missing", payload, param)
			continue
		} else if len(param.Type) == 0 {
			log.Warn().Msgf("injecting payload %s in %v with missing type", payload, param)
			paramType = "random"
		} else {
			paramType = param.Type[0]
		}

		switch paramType {
		case "string", "random":
			param.Value = payload
		}
	}

	return injectedParams
}

// generates Api tests by injecting payloads in values
func injectParamIntoApiTest(url string, docParams []*parser.DocHttpParams, queryParams map[string]string, headers map[string]string, testName string, injectionConfig InjectionConfig) []*ApiTest {
	var tests []*ApiTest
	docPrms := docParams
	// TODO: only inject payloads if any payload is accepted by the endpoint, else ignore injection
	// as this will reduce number of tests generated and increase efficiency
	for _, payload := range injectionConfig.Payloads {
		// TODO: implement injection in both key or value at a time
		for _, docParam := range docPrms {
			// inject payloads into string before converting it to map[string]string
			if injectionConfig.InBody {
				docParam.BodyParams = injectParamInParam(docParam.BodyParams, payload.InjText, Body)

			}
			if injectionConfig.InQuery {
				docParam.QueryParams = injectParamInParam(docParam.QueryParams, payload.InjText, Query)
			}
			if injectionConfig.InCookie {
				docParam.CookieParams = injectParamInParam(docParam.CookieParams, payload.InjText, Cookie)
			}
			if injectionConfig.InHeader {
				docParam.HeaderParams = injectParamInParam(docParam.HeaderParams, payload.InjText, Header)
			}

			// parse maps
			url, headersMap, queryMap, bodyData, pathWithParams, err := httpParamToRequest(url, docParam, queryParams, headers, utils.JSON)
			if err != nil {
				log.Error().Err(err).Msgf("failed to generate request params from DocHttpParams, skipping test for this case %v due to error %v", *docParam, err)
				continue
			}

			// check for uri endpoint injection in query param for vulnerable endpoint detection/backtracking
			// this is required since all endpoints will make call to same ssrf payload
			// so in order to detect vulnerable endpoint inject its uri path in query param
			// example: https://ssrf-website.com?offat_test_endpoint=/api/v1/users

			if injectionConfig.InjectUriInQuery {
				queryMap["offat_test_endpoint"] = docParam.Path
			}

			request := c.NewRequest(url, docParam.HttpMethod, queryMap, headersMap, bodyData)

			test := ApiTest{
				TestName:                testName,
				Request:                 request,
				Path:                    docParam.Path,
				PathWithParams:          pathWithParams,
				VulnerableResponseCodes: payload.VulnerableResponseCodes,
				ImmuneResponseCodes:     payload.ImmuneResponseCodes,
				MatchRegex:              payload.Regex,
			}
			tests = append(tests, &test)
		}
	}

	return tests
}
