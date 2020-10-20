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

		tickets, err := s.ticketRepo.List("organization_id", strOrgID)
		if err != nil {
			fmt.Printf("Cannot get tickets for organization %s. Error: %s\n", org.ID, err)
		}
		rs.Tickets = *tickets

		users, err := s.userRepo.List("organization_id", strOrgID)
		if err != nil {
			fmt.Printf("Cannot get tickets for organization %s. Error: %s\n", org.ID, err)
		}

		rs.UserNames = []string{}
		for _, u := range *users {
			rs.UserNames = append(rs.UserNames, u.Name)
		}

		results = append(results, &rs)
	}

	return &results, nil
}
