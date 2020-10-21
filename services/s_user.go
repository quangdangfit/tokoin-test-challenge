package services

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/copier"

	"tokoin/repositories"
	"tokoin/schema"
)

type UserService struct {
	orgRepo    repositories.IOrgRepository
	ticketRepo repositories.ITicketRepository
	userRepo   repositories.IUserRepository
}

func NewUserService(orgRepo repositories.IOrgRepository, ticketRepo repositories.ITicketRepository,
	userRepo repositories.IUserRepository) *UserService {
	return &UserService{
		orgRepo:    orgRepo,
		ticketRepo: ticketRepo,
		userRepo:   userRepo,
	}
}

func (s *UserService) List(key, value string) (*schema.Users, error) {
	users, err := s.userRepo.List(key, value)
	if err != nil {
		return nil, err
	}
	results := schema.Users{}
	for _, u := range *users {
		var rs schema.User
		copier.Copy(&rs, &u)
		strUID := strconv.Itoa(u.ID)

		// Get assignee tickets tickets for user
		assigneeTickets, err := s.ticketRepo.List("assignee_id", strUID)
		if err != nil {
			fmt.Printf("Cannot get assignee tickets of user %d. Error: %s\n", u.ID, err)
		}
		for _, t := range *assigneeTickets {
			rs.AssigneeTicketSubjects = append(rs.AssigneeTicketSubjects, t.Subject)
		}

		// Get submitted tickets tickets for user
		submittedTickets, err := s.ticketRepo.List("submitter_id", strUID)
		if err != nil {
			fmt.Printf("Cannot get submitted tickets of user %d. Error: %s\n", u.ID, err)
		}
		for _, t := range *submittedTickets {
			rs.SubmittedTicketSubjects = append(rs.SubmittedTicketSubjects, t.Subject)
		}

		// Get organization of user
		org, err := s.orgRepo.Retrieve(u.OrganizationID)
		if err != nil {
			fmt.Printf("Cannot get organization of user %d. Error: %s\n", u.ID, err)
		}
		rs.OrganizationName = org.Name

		results = append(results, &rs)
	}

	return &results, nil
}
