package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"tokoin/utils"
)

func TestUtilsStringToBoolean(t *testing.T) {
	testcases := []TestCase{
		{"Test convert valid lower value", "true", true, false},
		{"Test convert valid value", "True", true, false},
		{"Test convert invalid value", "input", false, true},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			result, err := utils.StringToBoolean(testcase.Args.(string))
			assert.Equal(t, testcase.ExpectedResult, result, err)
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}
