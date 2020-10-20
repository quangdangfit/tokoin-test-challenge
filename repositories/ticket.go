package repositories

import "tokoin/models"

type ITicketRepository interface {
	List(key string, value string) (*models.Tickets, error)
}
