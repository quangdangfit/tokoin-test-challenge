package services

import (
	"go.uber.org/dig"
)

func Inject(container *dig.Container) error {
	_ = container.Provide(NewOrgService)
	_ = container.Provide(NewTicketService)
	return nil
}
