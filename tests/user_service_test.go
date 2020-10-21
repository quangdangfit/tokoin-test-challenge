package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserServiceListExistedRecordAssigneeTicketSubjects(t *testing.T) {
	testcases := []SearchTestCase{
		// search existed record.
		{"Search by existed _id and check assignee ticket subjects", SearchArgs{"_id", "55"}, 1, false},
		{"Search by existed url and check assignee ticket subjects", SearchArgs{"url", "http://initech.tokoin.io.com/api/v2/users/55.json"}, 1, false},
		{"Search by existed external_id and check assignee ticket subjects", SearchArgs{"external_id", "95387bef-5870-4453-9431-be6f9864bad8"}, 1, false},
		{"Search by existed name and check assignee ticket subjects", SearchArgs{"name", "Benjamin Stephenson"}, 0, false},
		{"Search by existed alias and check assignee ticket subjects", SearchArgs{"alias", "Miss Louisa"}, 0, false},
		{"Search by existed created_at and check assignee ticket subjects", SearchArgs{"created_at", "2016-02-17T10:35:02 -11:00"}, 1, false},
		{"Search by existed shared and check assignee ticket subjects", SearchArgs{"shared", "true"}, 0, false},
		{"Search by existed timezone and check assignee ticket subjects", SearchArgs{"timezone", "Oman"}, 0, false},
		{"Search by existed last_login_at and check assignee ticket subjects", SearchArgs{"last_login_at", "2016-02-19T03:04:47 -11:00"}, 0, false},
		{"Search by existed email and check assignee ticket subjects", SearchArgs{"email", "hammondgaines@flotonic.com"}, 0, false},
		{"Search by existed phone and check assignee ticket subjects", SearchArgs{"phone", "8804-843-526"}, 0, false},
		{"Search by existed tags and check assignee ticket subjects", SearchArgs{"tags", "Frank"}, 0, false},
		{"Search by existed suspended and check assignee ticket subjects", SearchArgs{"suspended", "true"}, 0, false},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			results, err := mockUserService.List(testcase.Args.Key, testcase.Args.Value)
			assert.NotNil(t, results, err)
			assert.Greater(t, len(*results), 0, err)
			assert.Equal(t, testcase.ExpectedResult, len((*results)[0].AssigneeTicketSubjects), err)
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}

func TestUserServiceListExistedRecordSubmittedTicketSubjects(t *testing.T) {
	testcases := []SearchTestCase{
		// search existed record.
		{"Search by existed _id and check submitted ticket subjects", SearchArgs{"_id", "55"}, 1, false},
		{"Search by existed url and check submitted ticket subjects", SearchArgs{"url", "http://initech.tokoin.io.com/api/v2/users/55.json"}, 1, false},
		{"Search by existed external_id and check submitted ticket subjects", SearchArgs{"external_id", "95387bef-5870-4453-9431-be6f9864bad8"}, 1, false},
		{"Search by existed name and check submitted ticket subjects", SearchArgs{"name", "Benjamin Stephenson"}, 1, false},
		{"Search by existed alias and check submitted ticket subjects", SearchArgs{"alias", "Miss Louisa"}, 1, false},
		{"Search by existed created_at and check submitted ticket subjects", SearchArgs{"created_at", "2016-02-17T10:35:02 -11:00"}, 0, false},
		{"Search by existed shared and check submitted ticket subjects", SearchArgs{"shared", "true"}, 1, false},
		{"Search by existed timezone and check submitted ticket subjects", SearchArgs{"timezone", "Oman"}, 1, false},
		{"Search by existed last_login_at and check submitted ticket subjects", SearchArgs{"last_login_at", "2016-02-19T03:04:47 -11:00"}, 1, false},
		{"Search by existed email and check submitted ticket subjects", SearchArgs{"email", "hammondgaines@flotonic.com"}, 1, false},
		{"Search by existed phone and check submitted ticket subjects", SearchArgs{"phone", "8804-843-526"}, 1, false},
		{"Search by existed tags and check submitted ticket subjects", SearchArgs{"tags", "Frank"}, 1, false},
		{"Search by existed suspended and check submitted ticket subjects", SearchArgs{"suspended", "true"}, 1, false},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			results, err := mockUserService.List(testcase.Args.Key, testcase.Args.Value)
			assert.NotNil(t, results, err)
			assert.Greater(t, len(*results), 0, err)
			assert.Equal(t, testcase.ExpectedResult, len((*results)[0].SubmittedTicketSubjects), err)
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}

func TestUserServiceListExistedRecordOrganizationName(t *testing.T) {
	testcases := []SearchTestCase{
		// search existed record.
		{"Search by existed _id and check organization name", SearchArgs{"_id", "55"}, "Nutralab", false},
		{"Search by existed url and check organization name", SearchArgs{"url", "http://initech.tokoin.io.com/api/v2/users/55.json"}, "Nutralab", false},
		{"Search by existed external_id and check organization name", SearchArgs{"external_id", "95387bef-5870-4453-9431-be6f9864bad8"}, "Nutralab", false},
		{"Search by existed name and check organization name", SearchArgs{"name", "Benjamin Stephenson"}, "Enthaze", false},
		{"Search by existed alias and check organization name", SearchArgs{"alias", "Miss Louisa"}, "Enthaze", false},
		{"Search by existed created_at and check organization name", SearchArgs{"created_at", "2016-02-17T10:35:02 -11:00"}, "Nutralab", false},
		{"Search by existed shared and check organization name", SearchArgs{"shared", "true"}, "Enthaze", false},
		{"Search by existed timezone and check organization name", SearchArgs{"timezone", "Oman"}, "Enthaze", false},
		{"Search by existed last_login_at and check organization name", SearchArgs{"last_login_at", "2016-02-19T03:04:47 -11:00"}, "Enthaze", false},
		{"Search by existed email and check organization name", SearchArgs{"email", "hammondgaines@flotonic.com"}, "Enthaze", false},
		{"Search by existed phone and check organization name", SearchArgs{"phone", "8804-843-526"}, "Enthaze", false},
		{"Search by existed tags and check organization name", SearchArgs{"tags", "Frank"}, "Enthaze", false},
		{"Search by existed suspended and check organization name", SearchArgs{"suspended", "true"}, "Enthaze", false},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			results, err := mockUserService.List(testcase.Args.Key, testcase.Args.Value)
			assert.NotNil(t, results, err)
			assert.Greater(t, len(*results), 0, err)
			assert.Equal(t, testcase.ExpectedResult, (*results)[0].OrganizationName, err)
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}
