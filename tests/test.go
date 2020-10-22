package tests

import (
	"tokoin/repositories/files"
	"tokoin/services"
)

const (
	TestDataOrgFilePath    = "data/organizations.json"
	TestDataTicketFilePath = "data/tickets.json"
	TestDataUserFilePath   = "data/users.json"
)

type SearchArgs struct {
	Key   string
	Value string
}

type TestCase struct {
	Name           string
	Args           interface{}
	ExpectedResult interface{}
	ExpectedError  bool
}

type SearchTestCase struct {
	Name           string
	Args           SearchArgs
	ExpectedResult interface{}
	ExpectedError  bool
}

var mockOrgRepo *files.OrganizationRepo
var mockTicketRepo *files.TicketRepo
var mockUserRepo *files.UserRepo

var mockOrgService *services.OrgService
var mockTicketService *services.TicketService
var mockUserService *services.UserService

func init() {
	mockOrgRepo = &files.OrganizationRepo{}
	mockOrgRepo.LoadDataFromFile(TestDataOrgFilePath)

	mockTicketRepo = &files.TicketRepo{}
	mockTicketRepo.LoadDataFromFile(TestDataTicketFilePath)

	mockUserRepo = &files.UserRepo{}
	mockUserRepo.LoadDataFromFile(TestDataUserFilePath)

	mockOrgService = services.NewOrgService(mockOrgRepo, mockTicketRepo, mockUserRepo)
	mockTicketService = services.NewTicketService(mockOrgRepo, mockTicketRepo, mockUserRepo)
	mockUserService = services.NewUserService(mockOrgRepo, mockTicketRepo, mockUserRepo)
}
