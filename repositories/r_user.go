package repositories

import "tokoin/models"

type IUserRepository interface {
	Retrieve(id int) (*models.User, error)
	List(key, value string) (*models.Users, error)
	ListNames(key, value string) ([]string, error)
}
