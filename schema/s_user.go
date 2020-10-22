package schema

import "tokoin/utils"

type User struct {
	ID                      int      `json:"_id"`
	URL                     string   `json:"url"`
	ExternalID              string   `json:"external_id"`
	Name                    string   `json:"name"`
	Alias                   string   `json:"alias"`
	CreatedAt               string   `json:"created_at"`
	Active                  bool     `json:"active"`
	Verified                bool     `json:"verified"`
	Shared                  bool     `json:"shared"`
	Locale                  string   `json:"locale"`
	Timezone                string   `json:"timezone"`
	LastLoginAt             string   `json:"last_login_at"`
	Email                   string   `json:"email"`
	Phone                   string   `json:"phone"`
	Signature               string   `json:"signature"`
	OrganizationID          int      `json:"organization_id"`
	Tags                    []string `json:"tags"`
	Suspended               bool     `json:"suspended"`
	Role                    string   `json:"role"`
	AssigneeTicketSubjects  []string `json:"assignee_ticket_subjects"`
	SubmittedTicketSubjects []string `json:"submitted_ticket_subjects"`
	OrganizationName        string   `json:"organization_name"`
}

type Users []*User

func (u User) ToString() string {
	return utils.Jsonify(u)
}

func (u Users) ToString() string {
	rs := ""
	for _, e := range u {
		if e != nil {
			rs += e.ToString() + "\n" + "\n-------------------------------------------\n"
		}
	}
	return rs
}

func (u Users) Length() int {
	return len(u)
}
