package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"tokoin/repositories/files"
)

func TestTicketRepoLoadData(t *testing.T) {
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

func TestTicketRepoListExistedRecord(t *testing.T) {
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
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			results, err := mockTicketRepo.List(testcase.Args.Key, testcase.Args.Value)
			assert.NotNil(t, results, err)
			assert.Equal(t, testcase.ExpectedResult, len(*results), err)
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}

func TestTicketRepoListNotExistedRecord(t *testing.T) {
	testcases := []SearchTestCase{
		// search not existed record..
		{"Search by not existed _id", SearchArgs{"_id", "27c447d9-cfda-4415-9a72"}, 0, false},
		{"Search by not existed url", SearchArgs{"url", "http://initech.tokoin.io.com/api/v2/tickets/27c447d9-cfda-4415-9a72.json"}, 0, false},
		{"Search by not existed external_id", SearchArgs{"external_id", "385ac1f0-e1e9-4bed-ba06"}, 0, false},
		{"Search by not existed created_at", SearchArgs{"created_at", "2016-01-20T01:23:55 -11:11"}, 0, false},
		{"Search by not existed type", SearchArgs{"type", "type"}, 0, false},
		{"Search by not existed subject", SearchArgs{"subject", "A Problem in Vietnam"}, 0, false},
		{"Search by not existed description", SearchArgs{"description", "The description."}, 0, false},
		{"Search by not existed priority", SearchArgs{"priority", "unknown"}, 0, false},
		{"Search by not existed status", SearchArgs{"status", "sold"}, 0, false},
		{"Search by not existed submitter_id", SearchArgs{"submitter_id", "100"}, 0, false},
		{"Search by not existed assignee_id", SearchArgs{"assignee_id", "100"}, 0, false},
		{"Search by not existed organization_id", SearchArgs{"organization_id", "100"}, 0, false},
		{"Search by not existed tags", SearchArgs{"tags", "Quang"}, 0, false},
		{"Search by not existed due_at", SearchArgs{"due_at", "2016-08-08T07:24:14 -10:11"}, 0, false},
		{"Search by not existed via", SearchArgs{"via", "tv"}, 0, false},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			results, err := mockTicketRepo.List(testcase.Args.Key, testcase.Args.Value)
			assert.NotNil(t, results, err)
			assert.Equal(t, testcase.ExpectedResult, len(*results), err)
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}

func TestTicketSearchInvalidInput(t *testing.T) {
	testcases := []SearchTestCase{
		// search by invalid input.
		{"Search by invalid submitter_id", SearchArgs{"submitter_id", "submitter_id"}, 0, true},
		{"Search by invalid assignee_id", SearchArgs{"assignee_id", "assignee_id"}, 0, true},
		{"Search by invalid organization_id", SearchArgs{"organization_id", "organization_id"}, 0, true},
		{"Search by invalid has_incidents", SearchArgs{"has_incidents", "has_incidents"}, 0, true},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			results, err := mockTicketRepo.List(testcase.Args.Key, testcase.Args.Value)
			assert.NotNil(t, results, err)
			assert.Equal(t, testcase.ExpectedResult, len(*results), err)
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}
