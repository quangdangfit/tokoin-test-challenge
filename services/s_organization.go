package services

import (
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
		rs.TicketSubjects, _ = s.ticketRepo.ListSubjects("organization_id", strOrgID)

		// Get user names of organization
		rs.UserNames, _ = s.userRepo.ListNames("organization_id", strOrgID)

		results = append(results, &rs)
	}

	return &results, nil
}
