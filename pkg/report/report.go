package report

import (
	"encoding/json"
	"fmt"

	"github.com/owasp-offat/offat/pkg/tgen"
	"github.com/owasp-offat/offat/pkg/utils"
)

func Report(apiTests []*tgen.ApiTest, contentType string) ([]byte, error) {
	switch contentType {
	case utils.JSON:
		return json.Marshal(&apiTests)
	default:
		return nil, fmt.Errorf("invalid report content type")
	}
}
