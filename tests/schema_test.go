package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"tokoin/schema"
)

func TestOrganizationSchemaToString(t *testing.T) {
	testcases := []TestCase{
		{"Organization schema to string", nil, 1, false},
	}

	orgSchema := schema.Organization{}
	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			str := orgSchema.ToString()
			assert.NotEmpty(t, str)
		})
	}
}

func TestOrganizationsSchemaToString(t *testing.T) {
	testcases := []TestCase{
		{"Organizations schema to string", nil, 1, false},
	}

	orgSchema := schema.Organizations{}
	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			str := orgSchema.ToString()
			assert.Empty(t, str)
		})
	}
}

func TestTicketSchemaToString(t *testing.T) {
	testcases := []TestCase{
		{"Ticket schema to string", nil, 1, false},
	}

	ticketSchema := schema.Ticket{}
	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			str := ticketSchema.ToString()
			assert.NotEmpty(t, str)
		})
	}
}

func TestTicketsSchemaToString(t *testing.T) {
	testcases := []TestCase{
		{"Tickets schema to string", nil, 1, false},
	}

	ticketSchema := schema.Tickets{}
	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			str := ticketSchema.ToString()
			assert.Empty(t, str)
		})
	}
}

func TestUserSchemaToString(t *testing.T) {
	testcases := []TestCase{
		{"User schema to string", nil, 1, false},
	}

	userSchema := schema.User{}
	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			str := userSchema.ToString()
			assert.NotEmpty(t, str)
		})
	}
}

func TestUsersSchemaToString(t *testing.T) {
	testcases := []TestCase{
		{"Users schema to string", nil, 1, false},
	}

	userSchema := schema.Users{}
	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			str := userSchema.ToString()
			assert.Empty(t, str)
		})
	}
}
