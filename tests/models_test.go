package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"tokoin/models"
)

func TestOrganizationModelToString(t *testing.T) {
	testcases := []TestCase{
		{"Organization model to string", nil, 1, false},
	}

	orgModel := models.Organization{}
	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			str := orgModel.ToString()
			assert.NotEmpty(t, str)
		})
	}
}

func TestTicketModelToString(t *testing.T) {
	testcases := []TestCase{
		{"Ticket model to string", nil, 1, false},
	}

	ticketModel := models.Ticket{}
	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			str := ticketModel.ToString()
			assert.NotEmpty(t, str)
		})
	}
}

func TestUserModelToString(t *testing.T) {
	testcases := []TestCase{
		{"User model to string", nil, 1, false},
	}

	userModel := models.User{}
	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			str := userModel.ToString()
			assert.NotEmpty(t, str)
		})
	}
}
