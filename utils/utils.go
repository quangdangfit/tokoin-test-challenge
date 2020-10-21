package utils

import (
	"strings"

	"github.com/pkg/errors"
)

func StringToBoolean(s string) (bool, error) {
	v := strings.ToLower(s)
	if v == "true" {
		return true, nil
	}
	if v == "false" {
		return false, nil
	}

	return false, errors.New("invalid boolean string")
}
