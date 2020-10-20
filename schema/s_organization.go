package schema

import (
	"tokoin/models"
	"tokoin/utils"
)

type Organization struct {
	ID            int            `json:"_id"`
	URL           string         `json:"url"`
	ExternalID    string         `json:"external_id"`
	Name          string         `json:"name"`
	DomainNames   []string       `json:"domain_names"`
	CreatedAt     string         `json:"created_at"`
	Details       string         `json:"details"`
	SharedTickets bool           `json:"shared_tickets"`
	Tags          []string       `json:"tags"`
	Tickets       models.Tickets `json:"tickets"`
	UserNames     []string       `json:"user_names"`
}

type Organizations []*Organization

func (o Organization) ToString() string {
	return utils.Jsonify(o)
}
