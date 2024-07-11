package utils

import "strconv"

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
