package files

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"tokoin/config"
	"tokoin/models"
	"tokoin/repositories"
	"tokoin/utils"
)

type TicketRepo struct {
	tickets models.Tickets
}

func NewTicketRepository() repositories.ITicketRepository {
	path := config.Config.Data.Ticket
	data, err := utils.ReadJsonFile(path)
	if err != nil {
		fmt.Printf("Cannot load data from file %s. Error: %s\n", path, err.Error())
		return nil
	}

	var tickets models.Tickets
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Data cannot marshal to json. Error: %s\n", err.Error())
		return nil
	}

	json.Unmarshal(bytes, &tickets)
	return &TicketRepo{
		tickets: tickets,
	}
}

func (r *TicketRepo) List(key, value string) (*models.Tickets, error) {
	results := models.Tickets{}
	switch key {
	case "_id":
		for _, ticket := range r.tickets {
			if ticket.ID == value {
				results = append(results, ticket)
			}
		}
	case "url":
		for _, org := range r.tickets {
			if org.URL == value {
				results = append(results, org)
			}
		}
	case "external_id":
		for _, org := range r.tickets {
			if org.ExternalID == value {
				results = append(results, org)
			}
		}
	case "created_at":
		for _, org := range r.tickets {
			if org.CreatedAt == value {
				results = append(results, org)
			}
		}
	case "type":
		for _, org := range r.tickets {
			if org.Type == value {
				results = append(results, org)
			}
		}
	case "subject":
		for _, org := range r.tickets {
			if org.Subject == value {
				results = append(results, org)
			}
		}
	case "description":
		for _, org := range r.tickets {
			if org.Description == value {
				results = append(results, org)
			}
		}
	case "priority":
		for _, org := range r.tickets {
			if org.Priority == value {
				results = append(results, org)
			}
		}
	case "status":
		for _, org := range r.tickets {
			if org.Status == value {
				results = append(results, org)
			}
		}
	case "submitter_id":
		id, err := strconv.Atoi(value)
		if err != nil {
			return nil, errors.New("input submitter_id is invalid")
		}

		for _, org := range r.tickets {
			if org.SubmitterID == id {
				results = append(results, org)
			}
		}
	case "assignee_id":
		id, err := strconv.Atoi(value)
		if err != nil {
			return nil, errors.New("input assignee_id is invalid")
		}

		for _, org := range r.tickets {
			if org.AssigneeID == id {
				results = append(results, org)
			}
		}
	case "organization_id":
		id, err := strconv.Atoi(value)
		if err != nil {
			return nil, errors.New("input organization_id is invalid")
		}

		for _, org := range r.tickets {
			if org.OrganizationID == id {
				results = append(results, org)
			}
		}
	case "tags":
		for _, org := range r.tickets {
			for _, tag := range org.Tags {
				if tag == value {
					results = append(results, org)
					break
				}
			}
		}
	case "has_incidents":
		v := strings.ToLower(value) == "true"
		for _, org := range r.tickets {
			if org.HasIncidents == v {
				results = append(results, org)
			}
		}
	case "due_at":
		for _, org := range r.tickets {
			if org.DueAt == value {
				results = append(results, org)
			}
		}
	case "via":
		for _, org := range r.tickets {
			if org.Via == value {
				results = append(results, org)
			}
		}
	}

	return &results, nil
}
