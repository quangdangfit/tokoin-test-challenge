package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"tokoin/repositories/files"
)

const invalidDataUsers = `users`
const mockDataUsers = `[
  {
    "_id": 39,
    "url": "http://initech.tokoin.io.com/api/v2/users/39.json",
    "external_id": "66a3589c-8aab-485b-9da3-e79adb85a4de",
    "name": "Jennifer Gaines",
    "alias": "Miss Hammond",
    "created_at": "2016-03-25T02:02:13 -11:00",
    "active": false,
    "verified": true,
    "shared": true,
    "locale": "en-AU",
    "timezone": "Mayotte",
    "last_login_at": "2016-02-19T03:04:47 -11:00",
    "email": "hammondgaines@flotonic.com",
    "phone": "8804-843-526",
    "signature": "Don't Worry Be Happy!",
    "organization_id": 101,
    "tags": [
      "Tilleda",
      "Frank",
      "Fairmount",
      "Marion"
    ],
    "suspended": true,
    "role": "agent"
  },
  {
    "_id": 55,
    "url": "http://initech.tokoin.io.com/api/v2/users/55.json",
    "external_id": "95387bef-5870-4453-9431-be6f9864bad8",
    "name": "Thelma Wong",
    "alias": "Miss Robertson",
    "created_at": "2016-04-24T03:09:27 -10:00",
    "active": true,
    "shared": false,
    "locale": "en-AU",
    "timezone": "Guyana",
    "last_login_at": "2015-11-16T09:40:58 -11:00",
    "email": "robertsonwong@flotonic.com",
    "phone": "8685-603-206",
    "signature": "Don't Worry Be Happy!",
    "organization_id": 102,
    "tags": [
      "Delshire",
      "Ronco",
      "Farmers",
      "Foscoe"
    ],
    "suspended": false,
    "role": "end-user"
  },
  {
    "_id": 67,
    "url": "http://initech.tokoin.io.com/api/v2/users/67.json",
    "external_id": "01dd136a-bf3b-4340-94ee-8ecbea6fd65f",
    "name": "Benjamin Stephenson",
    "alias": "Miss Louisa",
    "created_at": "2016-04-20T12:26:44 -10:00",
    "active": true,
    "verified": true,
    "shared": false,
    "locale": "zh-CN",
    "timezone": "Oman",
    "last_login_at": "2016-04-06T12:17:21 -10:00",
    "email": "louisastephenson@flotonic.com",
    "phone": "9914-712-033",
    "signature": "Don't Worry Be Happy!",
    "organization_id": 101,
    "tags": [
      "Fairhaven",
      "Kraemer",
      "Mayfair",
      "Saranap"
    ],
    "suspended": false,
    "role": "agent"
  },
  {
    "_id": 74,
    "url": "http://initech.tokoin.io.com/api/v2/users/74.json",
    "external_id": "8fa4f74b-e690-4478-bf09-40fed1ebc417",
    "name": "Melissa Bishop",
    "alias": "Mr Katharine",
    "created_at": "2016-02-17T10:35:02 -11:00",
    "active": false,
    "verified": false,
    "shared": false,
    "locale": "en-AU",
    "timezone": "Sao Tome and Principe",
    "last_login_at": "2012-04-20T02:26:59 -10:00",
    "email": "katharinebishop@flotonic.com",
    "phone": "9025-522-621",
    "signature": "Don't Worry Be Happy!",
    "organization_id": 102,
    "tags": [
      "Shrewsbury",
      "Ryderwood",
      "Edmund",
      "Kersey"
    ],
    "suspended": false,
    "role": "admin"
  }
]
`

func TestUserRepoLoadDataFromFile(t *testing.T) {
	testcases := []TestCase{
		{"Load from existed file", sampleFilePath, nil, false},
		{"Load from not existed file", "", nil, true},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			repo := &files.UserRepo{}
			err := repo.LoadDataFromFile(testcase.Args.(string))
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}

func TestUserRepoLoadData(t *testing.T) {
	testcases := []TestCase{
		{"Load from existed file", mockDataUsers, nil, false},
		{"Load from not existed file", invalidDataUsers, nil, true},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			repo := &files.UserRepo{}
			err := repo.LoadDataFromBytes([]byte(testcase.Args.(string)))
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}

func TestUserRepoListExistedRecord(t *testing.T) {
	mockUserRepo := &files.UserRepo{}
	mockUserRepo.LoadDataFromBytes([]byte(mockDataUsers))

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
			results, err := mockUserRepo.List(testcase.Args.Key, testcase.Args.Value)
			assert.NotNil(t, results, err)
			assert.Equal(t, testcase.ExpectedResult, len(*results), err)
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}

func TestUserRepoListNotExistedRecord(t *testing.T) {
	mockUserRepo := &files.UserRepo{}
	mockUserRepo.LoadDataFromBytes([]byte(mockDataUsers))

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
			results, err := mockUserRepo.List(testcase.Args.Key, testcase.Args.Value)
			assert.NotNil(t, results, err)
			assert.Equal(t, testcase.ExpectedResult, len(*results), err)
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}

func TestUserRepoListInvalidInput(t *testing.T) {
	mockUserRepo := &files.UserRepo{}
	mockUserRepo.LoadDataFromBytes([]byte(mockDataUsers))

	testcases := []SearchTestCase{
		// search by invalid input.
		{"Search by invalid _id", SearchArgs{"_id", "id"}, 0, true},
		{"Search by invalid active", SearchArgs{"active", "active"}, 0, true},
		{"Search by invalid verified", SearchArgs{"verified", "verified"}, 0, true},
		{"Search by invalid shared", SearchArgs{"shared", "shared"}, 0, true},
		{"Search by invalid organization_id", SearchArgs{"organization_id", "organization_id"}, 0, true},
		{"Search by invalid suspended", SearchArgs{"suspended", "suspended"}, 0, true},
		{"Search by invalid key", SearchArgs{"key", "value"}, 0, true},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			results, err := mockUserRepo.List(testcase.Args.Key, testcase.Args.Value)
			assert.NotNil(t, results, err)
			assert.Equal(t, testcase.ExpectedResult, len(*results), err)
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}
