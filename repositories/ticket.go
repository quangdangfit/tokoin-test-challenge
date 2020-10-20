package repositories

import "tokoin/models"

type ITicketRepository interface {
	List(key, value string) (*models.Tickets, error)
}
