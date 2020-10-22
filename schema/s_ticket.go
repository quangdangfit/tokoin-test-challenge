package schema

import "tokoin/utils"

type Ticket struct {
	ID               string   `json:"_id"`
	URL              string   `json:"url"`
	ExternalID       string   `json:"external_id"`
	CreatedAt        string   `json:"created_at"`
	Type             string   `json:"type"`
	Subject          string   `json:"subject"`
	Description      string   `json:"description"`
	Priority         string   `json:"priority"`
	Status           string   `json:"status"`
	SubmitterID      int      `json:"submitter_id"`
	AssigneeID       int      `json:"assignee_id"`
	OrganizationID   int      `json:"organization_id"`
	Tags             []string `json:"tags"`
	HasIncidents     bool     `json:"has_incidents"`
	DueAt            string   `json:"due_at"`
	Via              string   `json:"via"`
	AssigneeName     string   `json:"assignee_name"`
	SubmitterName    string   `json:"submitter_name"`
	OrganizationName string   `json:"organization_name"`
}

type Tickets []*Ticket

func (t Ticket) ToString() string {
	return utils.Jsonify(t)
}

func (t Tickets) ToString() string {
	rs := ""
	for _, e := range t {
		if e != nil {
			rs += e.ToString() + "\n" + "\n-------------------------------------------\n"
		}
	}
	return rs
}

func (t Tickets) Length() int {
	return len(t)
}
