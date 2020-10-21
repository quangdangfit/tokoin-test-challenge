package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrganizationServiceSearchExistedRecordTicketSubjects(t *testing.T) {
	testcases := []SearchTestCase{
		// search existed record and check ticket subjects.
		{"Search by existed _id and check ticket subjects", SearchArgs{"_id", "101"}, 3, false},
		{"Search by existed url and check ticket subjects", SearchArgs{"url", "http://initech.tokoin.io.com/api/v2/organizations/101.json"}, 3, false},
		{"Search by existed external_id and check ticket subjects", SearchArgs{"external_id", "9270ed79-35eb-4a38-a46f-35725197ea8d"}, 3, false},
		{"Search by existed name and check ticket subjects", SearchArgs{"name", "Enthaze"}, 3, false},
		{"Search by existed domain_names and check ticket subjects", SearchArgs{"domain_names", "zentix.com"}, 3, false},
		{"Search by existed created_at and check ticket subjects", SearchArgs{"created_at", "2016-05-21T11:10:28 -10:00"}, 3, false},
		{"Search by existed details and check ticket subjects", SearchArgs{"details", "Non profit"}, 0, false},
		{"Search by existed tags and check ticket subjects", SearchArgs{"tags", "Collier"}, 0, false},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			results, err := mockOrgService.List(testcase.Args.Key, testcase.Args.Value)
			assert.NotNil(t, results, err)
			assert.Greater(t, len(*results), 0, err)
			assert.Equal(t, testcase.ExpectedResult, len((*results)[0].TicketSubjects), err)
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}

func TestOrganizationServiceSearchExistedRecordUserNames(t *testing.T) {
	testcases := []SearchTestCase{
		// search existed record and check ticket subjects.
		{"Search by existed _id and check user names", SearchArgs{"_id", "101"}, 2, false},
		{"Search by existed url and check user names", SearchArgs{"url", "http://initech.tokoin.io.com/api/v2/organizations/101.json"}, 2, false},
		{"Search by existed external_id and check user names", SearchArgs{"external_id", "9270ed79-35eb-4a38-a46f-35725197ea8d"}, 2, false},
		{"Search by existed name and check user names", SearchArgs{"name", "Enthaze"}, 2, false},
		{"Search by existed domain_names and check user names", SearchArgs{"domain_names", "zentix.com"}, 2, false},
		{"Search by existed created_at and check user names", SearchArgs{"created_at", "2016-05-21T11:10:28 -10:00"}, 2, false},
		{"Search by existed details and check user names", SearchArgs{"details", "Non profit"}, 2, false},
		{"Search by existed tags and check user names", SearchArgs{"tags", "Collier"}, 2, false},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			results, err := mockOrgService.List(testcase.Args.Key, testcase.Args.Value)
			assert.NotNil(t, results, err)
			assert.Greater(t, len(*results), 0, err)
			assert.Equal(t, testcase.ExpectedResult, len((*results)[0].UserNames), err)
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}
