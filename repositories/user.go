package repositories

import "tokoin/models"

type IUserRepository interface {
	Retrieve(key string, value string) (*models.User, error)
	List(key string, value string) (*[]models.User, error)
}
