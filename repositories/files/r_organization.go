package files

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"tokoin/config"
	"tokoin/models"
	"tokoin/repositories"
	"tokoin/utils"
)

type OrganizationRepo struct {
	organizations models.Organizations
}

func NewOrgRepository() repositories.IOrgRepository {
	path := config.Config.Data.Organization
	data, err := utils.ReadJsonFile(path)
	if err != nil {
		fmt.Printf("Cannot load data from file %s. Error: %s\n", path, err.Error())
		return nil
	}

	var orgs models.Organizations
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Data cannot marshal to json. Error: %s\n", err.Error())
		return nil
	}

	json.Unmarshal(bytes, &orgs)
	return &OrganizationRepo{
		organizations: orgs,
	}
}

func (r *OrganizationRepo) Retrieve(id int) (*models.Organization, error) {
	for _, org := range r.organizations {
		if org.ID == id {
			return org, nil
		}
	}

	return nil, nil
}

func (r *OrganizationRepo) List(key, value string) (*models.Organizations, error) {
	results := models.Organizations{}
	switch key {
	case "_id":
		id, err := strconv.Atoi(value)
		if err != nil {
			return nil, errors.New("input _id is invalid")
		}
		for _, org := range r.organizations {
			if org.ID == id {
				results = append(results, org)
			}
		}
	case "url":
		for _, org := range r.organizations {
			if org.URL == value {
				results = append(results, org)
			}
		}
	case "external_id":
		for _, org := range r.organizations {
			if org.ExternalID == value {
				results = append(results, org)
			}
		}
	case "name":
		for _, org := range r.organizations {
			if org.Name == value {
				results = append(results, org)
			}
		}
	case "domain_names":
		for _, org := range r.organizations {
			for _, d := range org.DomainNames {
				if d == value {
					results = append(results, org)
					break
				}
			}
		}
	case "created_at":
		for _, org := range r.organizations {
			if org.CreatedAt == value {
				results = append(results, org)
			}
		}
	case "details":
		for _, org := range r.organizations {
			if org.Details == value {
				results = append(results, org)
			}
		}
	case "shared_tickets":
		v := value == "true"
		for _, org := range r.organizations {
			if org.SharedTickets == v {
				results = append(results, org)
			}
		}
	case "tags":
		for _, org := range r.organizations {
			for _, tag := range org.Tags {
				if tag == value {
					results = append(results, org)
					break
				}
			}
		}
	}

	return &results, nil
}
