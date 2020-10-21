package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTicketServiceSearchExistedRecordAssigneeName(t *testing.T) {
	testcases := []SearchTestCase{
		// search existed record and check assignee name.
		{"Search by existed _id assignee name", SearchArgs{"_id", "27c447d9-cfda-4415-9a72-d5aa12942cf1"}, "Melissa Bishop", false},
		{"Search by existed url assignee name", SearchArgs{"url", "http://initech.tokoin.io.com/api/v2/tickets/27c447d9-cfda-4415-9a72-d5aa12942cf1.json"}, "Melissa Bishop", false},
		{"Search by existed external_id assignee name", SearchArgs{"external_id", "b17a9d1b-bc80-4262-a387-bb4f4209d7e2"}, "Thelma Wong", false},
		{"Search by existed subject assignee name", SearchArgs{"subject", "A Problem in Ethiopia"}, "Thelma Wong", false},
		{"Search by existed description assignee name", SearchArgs{"description", "Ex sit ea sit exercitation tempor pariatur et do deserunt irure eiusmod. Exercitation anim consectetur amet anim id."}, "Melissa Bishop", false},
		{"Search by existed status assignee name", SearchArgs{"status", "hold"}, "Thelma Wong", false},
		{"Search by existed submitter_id assignee name", SearchArgs{"submitter_id", "67"}, "Melissa Bishop", false},
		{"Search by existed assignee_id assignee name", SearchArgs{"assignee_id", "55"}, "Thelma Wong", false},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			results, err := mockTicketService.List(testcase.Args.Key, testcase.Args.Value)
			assert.NotNil(t, results, err)
			assert.Greater(t, len(*results), 0, err)
			assert.Equal(t, testcase.ExpectedResult, (*results)[0].AssigneeName, err)
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}

func TestTicketServiceSearchExistedRecordSubmitterName(t *testing.T) {
	testcases := []SearchTestCase{
		// search existed record and check submitted name.
		{"Search by existed _id assignee name", SearchArgs{"_id", "27c447d9-cfda-4415-9a72-d5aa12942cf1"}, "Benjamin Stephenson", false},
		{"Search by existed url assignee name", SearchArgs{"url", "http://initech.tokoin.io.com/api/v2/tickets/27c447d9-cfda-4415-9a72-d5aa12942cf1.json"}, "Benjamin Stephenson", false},
		{"Search by existed external_id assignee name", SearchArgs{"external_id", "b17a9d1b-bc80-4262-a387-bb4f4209d7e2"}, "Thelma Wong", false},
		{"Search by existed subject assignee name", SearchArgs{"subject", "A Problem in Ethiopia"}, "Thelma Wong", false},
		{"Search by existed description assignee name", SearchArgs{"description", "Ex sit ea sit exercitation tempor pariatur et do deserunt irure eiusmod. Exercitation anim consectetur amet anim id."}, "Benjamin Stephenson", false},
		{"Search by existed status assignee name", SearchArgs{"status", "hold"}, "Thelma Wong", false},
		{"Search by existed submitter_id assignee name", SearchArgs{"submitter_id", "67"}, "Benjamin Stephenson", false},
		{"Search by existed assignee_id assignee name", SearchArgs{"assignee_id", "55"}, "Thelma Wong", false},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			results, err := mockTicketService.List(testcase.Args.Key, testcase.Args.Value)
			assert.NotNil(t, results, err)
			assert.Greater(t, len(*results), 0, err)
			assert.Equal(t, testcase.ExpectedResult, (*results)[0].SubmitterName, err)
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}

func TestTicketServiceSearchExistedRecordOrganizationName(t *testing.T) {
	testcases := []SearchTestCase{
		// search existed record and check organization name.
		{"Search by existed _id organization name", SearchArgs{"_id", "27c447d9-cfda-4415-9a72-d5aa12942cf1"}, "Enthaze", false},
		{"Search by existed url organization name", SearchArgs{"url", "http://initech.tokoin.io.com/api/v2/tickets/27c447d9-cfda-4415-9a72-d5aa12942cf1.json"}, "Enthaze", false},
		{"Search by existed external_id organization name", SearchArgs{"external_id", "b17a9d1b-bc80-4262-a387-bb4f4209d7e2"}, "Enthaze", false},
		{"Search by existed subject organization name", SearchArgs{"subject", "A Problem in Ethiopia"}, "Enthaze", false},
		{"Search by existed description organization name", SearchArgs{"description", "Ex sit ea sit exercitation tempor pariatur et do deserunt irure eiusmod. Exercitation anim consectetur amet anim id."}, "Enthaze", false},
		{"Search by existed status organization name", SearchArgs{"status", "hold"}, "Enthaze", false},
		{"Search by existed submitter_id organization name", SearchArgs{"submitter_id", "67"}, "Enthaze", false},
		{"Search by existed assignee_id organization name", SearchArgs{"assignee_id", "55"}, "Enthaze", false},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			results, err := mockTicketService.List(testcase.Args.Key, testcase.Args.Value)
			assert.NotNil(t, results, err)
			assert.Greater(t, len(*results), 0, err)
			assert.Equal(t, testcase.ExpectedResult, (*results)[0].OrganizationName, err)
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}
