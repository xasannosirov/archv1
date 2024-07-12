package utils

import (
	"archv1/internal/pkg/config"
	"archv1/internal/pkg/tokens"
	"net/http"
	"strconv"
	"strings"
)

// QueryParams ...
type QueryParams struct {
	Page  int64
	Limit int64
}

// ParseQueryParams parse page, limit queries from param
func ParseQueryParams(queryParams map[string][]string) (*QueryParams, []string) {
	params := QueryParams{
		Page:  1,
		Limit: 10,
	}
	var errStr []string
	var err error

	for key, value := range queryParams {
		if key == "page" {
			params.Page, err = strconv.ParseInt(value[0], 10, 64)
			if err != nil {
				errStr = append(errStr, "Invalid `page` param")
			}
			continue
		}

		if key == "limit" {
			params.Limit, err = strconv.ParseInt(value[0], 10, 64)
			if err != nil {
				errStr = append(errStr, "Invalid `limit` param")
			}
			continue
		}
	}

	return &params, errStr
}

func GetTokenClaimsFromHeader(request *http.Request, config *config.Config) (map[string]interface{}, error) {
	token := request.Header.Get("Authorization")
	var softToken string
	if strings.Contains(token, "Bearer ") {
		softToken = strings.Split(token, "Bearer ")[1]
	} else {
		softToken = token
	}

	claims, err := tokens.ExtractClaim(softToken, []byte(config.JWTSecret))
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func GetLanguageFromHeader(request *http.Request) string {
	lang := strings.ToLower(request.Header.Get("Accept-Language"))

	if lang == "" {
		return "en"
	}

	return lang
}
