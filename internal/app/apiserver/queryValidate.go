package apiserver

import (
	"errors"
	"net/url"
	"regexp"
)

func validateQuery(vars url.Values, queryName string) error {
	err := errors.New("invalid queries")
	if len(vars) == 0 {
		return nil
	}
	if len(vars) > 1 {
		return err
	}
	if len(vars[queryName]) != 1 {
		return err
	}
	if match, _ := regexp.MatchString("[1-9]+", vars[queryName][0]); match == false {
		return err
	}
	return nil
}
