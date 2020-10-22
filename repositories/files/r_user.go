package files

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/pkg/errors"

	"tokoin/config"
	"tokoin/models"
	"tokoin/repositories"
	"tokoin/utils"
)

type UserRepo struct {
	users     models.Users
	userIDMap map[int]*models.User
}

func NewUserRepository() repositories.IUserRepository {
	userRepo := UserRepo{}
	userRepo.LoadDataFromFile(config.Config.Data.User)
	return &userRepo
}

func (r *UserRepo) LoadDataFromFile(path string) error {
	data, err := utils.ReadJsonFile(path)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("cannot load data from json file %s", path))
	}
	return r.LoadDataFromBytes(data)
}

func (r *UserRepo) LoadDataFromBytes(data []byte) error {
	var users models.Users
	err := json.Unmarshal(data, &users)
	if err != nil {
		return err
	}
	r.users = users
	r.userIDMap = map[int]*models.User{}
	for _, u := range users {
		r.userIDMap[u.ID] = u
	}

	return nil
}

func (r *UserRepo) Retrieve(id int) (*models.User, error) {
	return r.userIDMap[id], nil
}

func (r *UserRepo) List(key, value string) (*models.Users, error) {
	results := models.Users{}
	switch key {
	case "_id":
		id, err := strconv.Atoi(value)
		if err != nil {
			return &results, errors.New("input _id is invalid")
		}
		user, _ := r.Retrieve(id)
		if user != nil {
			results = append(results, user)
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
		v, err := utils.StringToBoolean(value)
		if err != nil {
			return &results, err
		}

		for _, user := range r.users {
			if user.Active == v {
				results = append(results, user)
			}
		}
	case "verified":
		v, err := utils.StringToBoolean(value)
		if err != nil {
			return &results, err
		}

		for _, user := range r.users {
			if user.Verified == v {
				results = append(results, user)
			}
		}
	case "shared":
		v, err := utils.StringToBoolean(value)
		if err != nil {
			return &results, err
		}

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
			return &results, errors.New("input organization_id is invalid")
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
		v, err := utils.StringToBoolean(value)
		if err != nil {
			return &results, err
		}

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
	default:
		return &results, errors.New("key is invalid")
	}

	return &results, nil
}

func (r *UserRepo) ListNames(key, value string) ([]string, error) {
	rs := []string{}
	users, err := r.List(key, value)
	if err != nil {
		return nil, err
	}

	for _, u := range *users {
		rs = append(rs, u.Name)
	}

	return rs, nil
}
