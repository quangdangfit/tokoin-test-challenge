package services

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/copier"

	"tokoin/repositories"
	"tokoin/schema"
)

type OrgService struct {
	orgRepo    repositories.IOrgRepository
	ticketRepo repositories.ITicketRepository
	userRepo   repositories.IUserRepository
}

func NewOrgService(orgRepo repositories.IOrgRepository, ticketRepo repositories.ITicketRepository,
	userRepo repositories.IUserRepository) *OrgService {
	return &OrgService{
		orgRepo:    orgRepo,
		ticketRepo: ticketRepo,
		userRepo:   userRepo,
	}
}

func (s *OrgService) List(key, value string) (*schema.Organizations, error) {
	orgs, err := s.orgRepo.List(key, value)
	if err != nil {
		return nil, err
	}
	results := schema.Organizations{}
	for _, org := range *orgs {
		var rs schema.Organization
		copier.Copy(&rs, &org)
		strOrgID := strconv.Itoa(org.ID)

		// Get tickets of organization
		ticketSubjects, err := s.getTicketSubjects(strOrgID)
		if err != nil {
			fmt.Printf("Cannot get ticket subjects of organization %s. Error: %s\n", strOrgID, err)
		}
		rs.TicketSubjects = ticketSubjects

		// Get user names of organization
		userNames, err := s.getUserNames(strOrgID)
		if err != nil {
			fmt.Printf("Cannot get user names of organization %s. Error: %s\n", strOrgID, err)
		}
		rs.UserNames = userNames

		results = append(results, &rs)
	}

	return &results, nil
}

func (s *OrgService) getTicketSubjects(orgID string) ([]string, error) {
	tSubjects := []string{}
	tickets, err := s.ticketRepo.List("organization_id", orgID)
	if err != nil {
		return tSubjects, err
	}

	for _, t := range *tickets {
		tSubjects = append(tSubjects, t.Subject)
	}

	return tSubjects, nil
}

func (s *OrgService) getUserNames(orgID string) ([]string, error) {
	uNames := []string{}
	users, err := s.userRepo.List("organization_id", orgID)
	if err != nil {
		return uNames, err
	}

	for _, u := range *users {
		uNames = append(uNames, u.Name)
	}

	return uNames, nil
}
