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
	) {
		results, err := ticketService.List("url", "http://initech.tokoin.io.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json")
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
