package files

import (
	"go.uber.org/dig"
)

func Inject(container *dig.Container) error {
	_ = container.Provide(NewOrgRepository)
	_ = container.Provide(NewTicketRepository)
	_ = container.Provide(NewUserRepository)
	return nil
}
