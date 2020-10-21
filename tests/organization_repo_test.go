package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"tokoin/repositories/files"
)

func TestOrganizationRepoLoadData(t *testing.T) {
	testcases := []TestCase{
		{"Load from existed file", TestDataOrgFilePath, nil, false},
		{"Load from not existed file", "", nil, true},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			repo := &files.OrganizationRepo{}
			err := repo.LoadData(testcase.Args.(string))
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}

func TestOrganizationRepoListExistedRecord(t *testing.T) {
	testcases := []SearchTestCase{
		// search existed record.
		{"Search by existed _id", SearchArgs{"_id", "101"}, 1, false},
		{"Search by existed url", SearchArgs{"url", "http://initech.tokoin.io.com/api/v2/organizations/101.json"}, 1, false},
		{"Search by existed external_id", SearchArgs{"external_id", "9270ed79-35eb-4a38-a46f-35725197ea8d"}, 1, false},
		{"Search by existed name", SearchArgs{"name", "Enthaze"}, 1, false},
		{"Search by existed domain_names", SearchArgs{"domain_names", "zentix.com"}, 1, false},
		{"Search by existed created_at", SearchArgs{"created_at", "2016-05-21T11:10:28 -10:00"}, 1, false},
		{"Search by existed details", SearchArgs{"details", "Non profit"}, 1, false},
		{"Search by existed shared_tickets", SearchArgs{"shared_tickets", "false"}, 2, false},
		{"Search by existed tags", SearchArgs{"tags", "Collier"}, 1, false},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			results, err := mockOrgRepo.List(testcase.Args.Key, testcase.Args.Value)
			assert.NotNil(t, results, err)
			assert.Equal(t, testcase.ExpectedResult, len(*results), err)
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}

func TestOrganizationRepoListNotExistedRecord(t *testing.T) {
	testcases := []SearchTestCase{
		// search not existed record..
		{"Search by not existed _id", SearchArgs{"_id", "111"}, 0, false},
		{"Search by not existed url", SearchArgs{"url", "http://initech.tokoin.io.com/api/v2/organizations/111.json"}, 0, false},
		{"Search by not existed external_id", SearchArgs{"external_id", "9270ed79-35eb-4a38-a46f-35725197ea11"}, 0, false},
		{"Search by not existed name", SearchArgs{"name", "Enthaze11"}, 0, false},
		{"Search by not existed domain_names", SearchArgs{"domain_names", "zentix11.com"}, 0, false},
		{"Search by not existed created_at", SearchArgs{"created_at", "2016-05-21T11:11:11 -10:00"}, 0, false},
		{"Search by not existed details", SearchArgs{"details", "Non profit11"}, 0, false},
		{"Search by not existed shared_tickets", SearchArgs{"shared_tickets", "true"}, 0, false},
		{"Search by not existed tags", SearchArgs{"tags", "Collier11"}, 0, false},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			results, err := mockOrgRepo.List(testcase.Args.Key, testcase.Args.Value)
			assert.NotNil(t, results, err)
			assert.Equal(t, testcase.ExpectedResult, len(*results), err)
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}

func TestOrganizationRepoListInvalidInput(t *testing.T) {
	testcases := []SearchTestCase{
		// search by invalid input.
		{"Search by invalid _id", SearchArgs{"_id", "id"}, 0, true},
		{"Search by invalid shared_tickets", SearchArgs{"shared_tickets", "fasdf"}, 0, true},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			results, err := mockOrgRepo.List(testcase.Args.Key, testcase.Args.Value)
			assert.NotNil(t, results, err)
			assert.Equal(t, testcase.ExpectedResult, len(*results), err)
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}
