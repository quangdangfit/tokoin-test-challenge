package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"tokoin/repositories/files"
)

const (
	TestDataTicketFilePath = "data/tickets.json"
)

func TestTicketLoadData(t *testing.T) {
	testcases := []TestCase{
		{"Load from existed file", TestDataTicketFilePath, nil, false},
		{"Load from not existed file", "", nil, true},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			repo := &files.TicketRepo{}
			err := repo.LoadData(testcase.Args.(string))
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}

func TestTicketSearch(t *testing.T) {
	orgRepo := &files.TicketRepo{}
	assert.Nil(t, orgRepo.LoadData(TestDataTicketFilePath))

	testcases := []SearchTestCase{
		// search existed record.
		{"Search by existed _id", SearchArgs{"_id", "27c447d9-cfda-4415-9a72-d5aa12942cf1"}, 1, false},
		{"Search by existed url", SearchArgs{"url", "http://initech.tokoin.io.com/api/v2/tickets/27c447d9-cfda-4415-9a72-d5aa12942cf1.json"}, 1, false},
		{"Search by existed external_id", SearchArgs{"external_id", "385ac1f0-e1e9-4bed-ba06-2f3013d8e914"}, 1, false},
		{"Search by existed created_at", SearchArgs{"created_at", "2016-01-20T01:23:55 -11:00"}, 1, false},
		{"Search by existed type", SearchArgs{"type", "incident"}, 2, false},
		{"Search by existed subject", SearchArgs{"subject", "A Problem in Ethiopia"}, 1, false},
		{"Search by existed description", SearchArgs{"description", "Ex sit ea sit exercitation tempor pariatur et do deserunt irure eiusmod. Exercitation anim consectetur amet anim id."}, 1, false},
		{"Search by existed priority", SearchArgs{"priority", "low"}, 2, false},
		{"Search by existed status", SearchArgs{"status", "hold"}, 1, false},
		{"Search by existed submitter_id", SearchArgs{"submitter_id", "67"}, 1, false},
		{"Search by existed assignee_id", SearchArgs{"assignee_id", "55"}, 1, false},
		{"Search by existed organization_id", SearchArgs{"organization_id", "101"}, 3, false},
		{"Search by existed tags", SearchArgs{"tags", "Maine"}, 1, false},
		{"Search by existed has_incidents", SearchArgs{"has_incidents", "false"}, 2, false},
		{"Search by existed due_at", SearchArgs{"due_at", "2016-08-08T07:24:14 -10:00"}, 1, false},
		{"Search by existed via", SearchArgs{"via", "web"}, 3, false},

		//// search not existed record..
		//{"Search by not existed _id", SearchArgs{"_id", "111"}, 0, false},
		//{"Search by not existed url", SearchArgs{"url", "http://initech.tokoin.io.com/api/v2/organizations/111.json"}, 0, false},
		//{"Search by not existed external_id", SearchArgs{"external_id", "9270ed79-35eb-4a38-a46f-35725197ea11"}, 0, false},
		//{"Search by not existed name", SearchArgs{"name", "Enthaze11"}, 0, false},
		//{"Search by not existed domain_names", SearchArgs{"domain_names", "zentix11.com"}, 0, false},
		//{"Search by not existed created_at", SearchArgs{"created_at", "2016-05-21T11:11:11 -10:00"}, 0, false},
		//{"Search by not existed details", SearchArgs{"details", "Non profit11"}, 0, false},
		//{"Search by not existed shared_tickets", SearchArgs{"shared_tickets", "true"}, 0, false},
		//{"Search by not existed tags", SearchArgs{"tags", "Collier11"}, 0, false},
		//
		//// search by invalid input.
		//{"Search by invalid _id", SearchArgs{"_id", "id"}, 0, true},
		//{"Search by invalid shared_tickets", SearchArgs{"shared_tickets", "fasdf"}, 0, true},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			result, err := orgRepo.List(testcase.Args.Key, testcase.Args.Value)
			assert.NotNil(t, result, err)
			assert.Equal(t, testcase.ExpectedResult, len(*result), err)
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}