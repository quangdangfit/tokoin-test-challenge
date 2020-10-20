package services

import (
	"go.uber.org/dig"
)

func Inject(container *dig.Container) error {
	_ = container.Provide(NewOrgService)
	_ = container.Provide(NewTicketService)
	_ = container.Provide(NewUserService)
	return nil
}
