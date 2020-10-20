package main

import (
	"fmt"

	"go.uber.org/dig"

	"tokoin/repositories/files"
	"tokoin/services"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	// Inject repositories
	err := files.Inject(container)
	if err != nil {
		fmt.Println("Failed to inject repositories ", err)
	}

	// Inject repositories
	err = services.Inject(container)
	if err != nil {
		fmt.Println("Failed to inject services ", err)
	}

	return container
}

func main() {
	fmt.Println("Welcome to searching program ...")

	container := BuildContainer()

	container.Invoke(func(
		orgService *services.OrgService,
		ticketService *services.TicketService,
		userService *services.UserService,
	) {
		results, err := userService.List("url", "http://initech.tokoin.io.com/api/v2/users/38.json")
		if err != nil {
			fmt.Println(err)
		}

		if results != nil {
			for _, e := range *results {
				fmt.Println(e.ToString())
			}
		}

	})
}
