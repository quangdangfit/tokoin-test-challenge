package repositories

import "tokoin/models"

type IOrgRepository interface {
	Retrieve(key string, value string) (*models.Organization, error)
	List(key string, value string) (*[]models.Organization, error)
}
