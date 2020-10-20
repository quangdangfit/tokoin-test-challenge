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
	) {
		orgs, err := orgService.List("_id", "119")
		if err != nil {
			fmt.Println(err)
		}

		if orgs != nil {
			for _, org := range *orgs {
				fmt.Println(org.ToString())
			}
		}

	})
}
