package services

import (
	"fmt"

	"github.com/jinzhu/copier"

	"tokoin/repositories"
	"tokoin/schema"
)

type TicketService struct {
	orgRepo    repositories.IOrgRepository
	ticketRepo repositories.ITicketRepository
	userRepo   repositories.IUserRepository
}

func NewTicketService(orgRepo repositories.IOrgRepository, ticketRepo repositories.ITicketRepository,
	userRepo repositories.IUserRepository) *TicketService {
	return &TicketService{
		orgRepo:    orgRepo,
		ticketRepo: ticketRepo,
		userRepo:   userRepo,
	}
}

func (s *TicketService) List(key, value string) (*schema.Tickets, error) {
	tickets, err := s.ticketRepo.List(key, value)
	if err != nil {
		return nil, err
	}
	results := schema.Tickets{}
	for _, ticket := range *tickets {
		var rs schema.Ticket
		copier.Copy(&rs, &ticket)

		// Get organization of ticket
		org, err := s.orgRepo.Retrieve(ticket.OrganizationID)
		if err != nil {
			fmt.Printf("Cannot get organization of ticket %s. Error: %s\n", ticket.ID, err)
		}
		rs.OrganizationName = org.Name

		// Get assignee for ticket
		assignee, err := s.userRepo.Retrieve(ticket.AssigneeID)
		if err != nil {
			fmt.Printf("Cannot get assignee of ticket %s. Error: %s\n", ticket.ID, err)
		}
		rs.AssigneeName = assignee.Name

		// Get submitter for ticket
		submitter, err := s.userRepo.Retrieve(ticket.SubmitterID)
		if err != nil {
			fmt.Printf("Cannot get submitter of ticket %s. Error: %s\n", ticket.ID, err)
		}
		rs.SubmitterName = submitter.Name

		results = append(results, &rs)
	}

	return &results, nil
}
