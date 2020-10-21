package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"tokoin/repositories/files"
)

func TestUserLoadData(t *testing.T) {
	testcases := []TestCase{
		{"Load from existed file", TestDataUserFilePath, nil, false},
		{"Load from not existed file", "", nil, true},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			repo := &files.UserRepo{}
			err := repo.LoadData(testcase.Args.(string))
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}

func TestUserSearchExistedRecord(t *testing.T) {
	orgRepo := &files.UserRepo{}
	assert.Nil(t, orgRepo.LoadData(TestDataUserFilePath))

	testcases := []SearchTestCase{
		// search existed record.
		{"Search by existed _id", SearchArgs{"_id", "39"}, 1, false},
		{"Search by existed url", SearchArgs{"url", "http://initech.tokoin.io.com/api/v2/users/55.json"}, 1, false},
		{"Search by existed external_id", SearchArgs{"external_id", "95387bef-5870-4453-9431-be6f9864bad8"}, 1, false},
		{"Search by existed name", SearchArgs{"name", "Benjamin Stephenson"}, 1, false},
		{"Search by existed alias", SearchArgs{"alias", "Miss Louisa"}, 1, false},
		{"Search by existed created_at", SearchArgs{"created_at", "2016-02-17T10:35:02 -11:00"}, 1, false},
		{"Search by existed active", SearchArgs{"active", "false"}, 2, false},
		{"Search by existed verified", SearchArgs{"verified", "false"}, 2, false},
		{"Search by existed shared", SearchArgs{"shared", "true"}, 1, false},
		{"Search by existed locale", SearchArgs{"locale", "en-AU"}, 3, false},
		{"Search by existed timezone", SearchArgs{"timezone", "Oman"}, 1, false},
		{"Search by existed last_login_at", SearchArgs{"last_login_at", "2016-02-19T03:04:47 -11:00"}, 1, false},
		{"Search by existed email", SearchArgs{"email", "hammondgaines@flotonic.com"}, 1, false},
		{"Search by existed phone", SearchArgs{"phone", "8804-843-526"}, 1, false},
		{"Search by existed signature", SearchArgs{"signature", "Don't Worry Be Happy!"}, 4, false},
		{"Search by existed organization_id", SearchArgs{"organization_id", "101"}, 2, false},
		{"Search by existed tags", SearchArgs{"tags", "Frank"}, 1, false},
		{"Search by existed suspended", SearchArgs{"suspended", "true"}, 1, false},
		{"Search by existed role", SearchArgs{"role", "agent"}, 2, false},
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

func TestUserSearchNotExistedRecord(t *testing.T) {
	orgRepo := &files.UserRepo{}
	assert.Nil(t, orgRepo.LoadData(TestDataUserFilePath))

	testcases := []SearchTestCase{
		// search not existed record..
		{"Search by not existed _id", SearchArgs{"_id", "01"}, 0, false},
		{"Search by not existed url", SearchArgs{"url", "http://initech.tokoin.io.com/api/v2/users/01.json"}, 0, false},
		{"Search by not existed external_id", SearchArgs{"external_id", "95387bef-5870-4453-9431"}, 0, false},
		{"Search by not existed name", SearchArgs{"name", "Quang Dang"}, 0, false},
		{"Search by not existed alias", SearchArgs{"alias", "Miss Hoa"}, 0, false},
		{"Search by not existed created_at", SearchArgs{"created_at", "2016-02-17T10:35:02 -11:11"}, 0, false},
		{"Search by not existed locale", SearchArgs{"locale", "vi-VN"}, 0, false},
		{"Search by not existed timezone", SearchArgs{"timezone", "Vietnam"}, 0, false},
		{"Search by not existed last_login_at", SearchArgs{"last_login_at", "2016-02-19T03:04:47 -11:11"}, 0, false},
		{"Search by not existed email", SearchArgs{"email", "email@email.com"}, 0, false},
		{"Search by not existed phone", SearchArgs{"phone", "1111-222-333"}, 0, false},
		{"Search by not existed signature", SearchArgs{"signature", "The signature"}, 0, false},
		{"Search by not existed organization_id", SearchArgs{"organization_id", "100"}, 0, false},
		{"Search by not existed tags", SearchArgs{"tags", "Dang"}, 0, false},
		{"Search by not existed role", SearchArgs{"role", "role"}, 0, false},
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

func TestUserSearchInvalidInput(t *testing.T) {
	orgRepo := &files.UserRepo{}
	assert.Nil(t, orgRepo.LoadData(TestDataUserFilePath))

	testcases := []SearchTestCase{
		// search by invalid input.
		{"Search by invalid _id", SearchArgs{"_id", "id"}, 0, true},
		{"Search by invalid active", SearchArgs{"active", "active"}, 0, true},
		{"Search by invalid verified", SearchArgs{"verified", "verified"}, 0, true},
		{"Search by invalid shared", SearchArgs{"shared", "shared"}, 0, true},
		{"Search by invalid organization_id", SearchArgs{"organization_id", "organization_id"}, 0, true},
		{"Search by invalid suspended", SearchArgs{"suspended", "suspended"}, 0, true},
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
