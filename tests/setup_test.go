package tests

import (
	"tokoin/repositories/files"
	"tokoin/services"
)

const (
	sampleFilePath = "data/sample.json"
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

func NewOrgRepo() *files.OrganizationRepo {
	mockOrgRepo := &files.OrganizationRepo{}
	mockOrgRepo.LoadDataFromBytes([]byte(mockDataOrg))

	return mockOrgRepo
}

func NewTicketRepo() *files.TicketRepo {
	mockTicketRepo := &files.TicketRepo{}
	mockTicketRepo.LoadDataFromBytes([]byte(mockDataTickets))

	return mockTicketRepo
}

func NewUserRepo() *files.UserRepo {
	mockUserRepo := &files.UserRepo{}
	mockUserRepo.LoadDataFromBytes([]byte(mockDataUsers))

	return mockUserRepo
}

func init() {
	mockOrgRepo = NewOrgRepo()
	mockTicketRepo = NewTicketRepo()
	mockUserRepo = NewUserRepo()

	mockOrgService = services.NewOrgService(mockOrgRepo, mockTicketRepo, mockUserRepo)
	mockTicketService = services.NewTicketService(mockOrgRepo, mockTicketRepo, mockUserRepo)
	mockUserService = services.NewUserService(mockOrgRepo, mockTicketRepo, mockUserRepo)
}
