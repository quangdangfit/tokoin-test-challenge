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

var orgRepo *files.OrganizationRepo
var ticketRepo *files.TicketRepo
var userRepo *files.UserRepo

var mockOrgService *services.OrgService
var ticketService *services.TicketService
var userService *services.UserService

func init() {
	orgRepo = &files.OrganizationRepo{}
	orgRepo.LoadData(TestDataOrgFilePath)

	ticketRepo = &files.TicketRepo{}
	ticketRepo.LoadData(TestDataTicketFilePath)

	userRepo = &files.UserRepo{}
	userRepo.LoadData(TestDataUserFilePath)

	mockOrgService = services.NewOrgService(orgRepo, ticketRepo, userRepo)
	ticketService = services.NewTicketService(orgRepo, ticketRepo, userRepo)
	userService = services.NewUserService(orgRepo, ticketRepo, userRepo)
}
