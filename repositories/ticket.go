package repositories

import "tokoin/models"

type ITicketRepository interface {
	Retrieve(key string, value string) (*models.Ticket, error)
	List(key string, value string) (*[]models.Ticket, error)
}
