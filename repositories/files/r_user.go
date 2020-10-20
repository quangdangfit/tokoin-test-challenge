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

type UserRepo struct {
	users models.Users
}

func NewUserRepository() repositories.IUserRepository {
	path := config.Config.Data.User
	data, err := utils.ReadJsonFile(path)
	if err != nil {
		fmt.Printf("Cannot load data from file %s. Error: %s\n", path, err.Error())
		return nil
	}

	var users models.Users
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Data cannot marshal to json. Error: %s\n", err.Error())
		return nil
	}

	json.Unmarshal(bytes, &users)
	return &UserRepo{
		users: users,
	}
}

func (r *UserRepo) Retrieve(id int) (*models.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, nil
}

func (r *UserRepo) List(key, value string) (*models.Users, error) {
	results := models.Users{}
	switch key {
	case "_id":
		id, err := strconv.Atoi(value)
		if err != nil {
			return nil, errors.New("input _id is invalid")
		}
		for _, user := range r.users {
			if user.ID == id {
				results = append(results, user)
			}
		}
	case "url":
		for _, user := range r.users {
			if user.URL == value {
				results = append(results, user)
			}
		}
	case "external_id":
		for _, user := range r.users {
			if user.ExternalID == value {
				results = append(results, user)
			}
		}
	case "name":
		for _, user := range r.users {
			if user.Name == value {
				results = append(results, user)
			}
		}
	case "alias":
		for _, user := range r.users {
			if user.Alias == value {
				results = append(results, user)
			}
		}
	case "created_at":
		for _, user := range r.users {
			if user.CreatedAt == value {
				results = append(results, user)
			}
		}
	case "active":
		v := strings.ToLower(value) == "true"
		for _, user := range r.users {
			if user.Active == v {
				results = append(results, user)
			}
		}
	case "verified":
		v := strings.ToLower(value) == "true"
		for _, user := range r.users {
			if user.Verified == v {
				results = append(results, user)
			}
		}
	case "shared":
		v := strings.ToLower(value) == "true"
		for _, user := range r.users {
			if user.Shared == v {
				results = append(results, user)
			}
		}
	case "locale":
		for _, user := range r.users {
			if user.Locale == value {
				results = append(results, user)
			}
		}
	case "timezone":
		for _, user := range r.users {
			if user.Timezone == value {
				results = append(results, user)
			}
		}
	case "last_login_at":
		for _, user := range r.users {
			if user.LastLoginAt == value {
				results = append(results, user)
			}
		}
	case "email":
		for _, user := range r.users {
			if user.Email == value {
				results = append(results, user)
			}
		}
	case "phone":
		for _, user := range r.users {
			if user.Phone == value {
				results = append(results, user)
			}
		}
	case "signature":
		for _, user := range r.users {
			if user.Signature == value {
				results = append(results, user)
			}
		}
	case "organization_id":
		id, err := strconv.Atoi(value)
		if err != nil {
			return nil, errors.New("input organization_id is invalid")
		}
		for _, user := range r.users {
			if user.OrganizationID == id {
				results = append(results, user)
			}
		}
	case "tags":
		for _, user := range r.users {
			for _, tag := range user.Tags {
				if tag == value {
					results = append(results, user)
					break
				}
			}
		}
	case "suspended":
		v := strings.ToLower(value) == "true"
		for _, user := range r.users {
			if user.Suspended == v {
				results = append(results, user)
			}
		}
	case "role":
		for _, user := range r.users {
			if user.Role == value {
				results = append(results, user)
			}
		}
	}

	return &results, nil
}
