package main

import (
	"fmt"

	"go.uber.org/dig"

	"tokoin/repositories"
	"tokoin/repositories/files"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	// Inject repositories
	err := files.Inject(container)
	if err != nil {
		fmt.Println("Failed to inject repositories ", err)
	}

	return container
}

func main() {
	fmt.Println("Welcome to searching program ...")

	container := BuildContainer()

	container.Invoke(func(
		orgRepo repositories.IOrgRepository,
		ticketRepo repositories.ITicketRepository,
	) error {
		orgs, err := orgRepo.List("_id", "101")
		if err != nil {
			fmt.Println(err)
		}

		if orgs != nil {
			for _, org := range *orgs {
				fmt.Println(org.ToString())
			}
		}

		tickets, err := ticketRepo.List("organization_id", "116")
		if tickets != nil {
			for _, ticket := range *tickets {
				fmt.Println(ticket.ToString())
			}
		}

		return nil
	})
}
