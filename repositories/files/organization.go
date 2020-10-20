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

type Organization struct {
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
	return &Organization{
		organizations: orgs,
	}
}

func (o *Organization) Retrieve(id int) (*models.Organization, error) {
	for _, org := range o.organizations {
		if org.ID == id {
			return org, nil
		}
	}

	return nil, nil
}

func (o *Organization) List(key, value string) (*models.Organizations, error) {
	results := models.Organizations{}
	switch key {
	case "_id":
		id, err := strconv.Atoi(value)
		if err != nil {
			return nil, errors.New("input _id is invalid")
		}
		for _, org := range o.organizations {
			if org.ID == id {
				results = append(results, org)
			}
		}
	}

	return &results, nil
}
