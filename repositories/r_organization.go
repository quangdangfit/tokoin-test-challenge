package repositories

import "tokoin/models"

type IOrgRepository interface {
	Retrieve(id int) (*models.Organization, error)
	List(key, value string) (*models.Organizations, error)
}
