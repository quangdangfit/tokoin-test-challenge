package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"go.uber.org/dig"

	"tokoin/repositories/files"
	"tokoin/services"
)

const welcomeMessage = `Welcome to Tokoin Code Challenge Program. Select search options:
 * Type '0' to exit.	
 * Type '1' to search.
`

const chooseModelMessage = `Please choose model you want to search. Select options:
 * Type '0' to exit.	
 * Type '1' to search organization.
 * Type '2' to search ticket.
 * Type '3' to search user.
`

const inputFieldMessage = `Input field: `
const inputValueMessage = `Input value: `

const loopMessage = `-----------------------------------------------------------
 * Type '0' to exit.	
 * Type '1' to search.
`

const retryMessage = `Your input is invalid!
 * Type '0' to exit.	
 * Type '1' to search.
`

func Run(container *dig.Container) {
	fmt.Print(welcomeMessage)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	opt := scanner.Text()
	for opt != "0" && opt != "1" {
		fmt.Print(retryMessage)

		scanner.Scan()
		opt = scanner.Text()
	}

	for opt != "0" {
		fmt.Print(chooseModelMessage)
		scanner.Scan()
		opt := scanner.Text()
		if opt == "0" {
			return
		}
		intOpt, err := strconv.Atoi(opt)
		if err != nil || intOpt < 0 || intOpt > 3 {
			fmt.Print("Invalid option, error: ", err)
		}

		fmt.Print(inputFieldMessage)
		scanner.Scan()
		field := scanner.Text()

		fmt.Print(inputValueMessage)
		scanner.Scan()
		value := scanner.Text()

		SearchHandler(container, intOpt, field, value)

		fmt.Print(loopMessage)
		scanner.Scan()
		opt = scanner.Text()
		print("Here" + opt)
	}

}

func SearchHandler(container *dig.Container, model int, key string, value string) {
	container.Invoke(func(
		orgService *services.OrgService,
		ticketService *services.TicketService,
		userService *services.UserService,
	) {
		switch model {
		case 1:
			results, err := orgService.List(key, value)
			if err != nil {
				fmt.Println(err)
			}

			if results != nil {
				fmt.Printf("Found %d records\n", len(*results))
				for _, e := range *results {
					fmt.Println("-----------------------------------------------------------")
					fmt.Println(e.ToString())
				}
			}
			break
		case 2:
			results, err := ticketService.List(key, value)
			if err != nil {
				fmt.Println(err)
			}

			if results != nil {
				fmt.Printf("Found %d records\n", len(*results))
				for _, e := range *results {
					fmt.Println("-----------------------------------------------------------")
					fmt.Println(e.ToString())
				}
			}
			break
		case 3:
			results, err := userService.List(key, value)
			if err != nil {
				fmt.Println(err)
			}

			if results != nil {
				fmt.Printf("Found %d records\n", len(*results))
				for _, e := range *results {
					fmt.Println("-----------------------------------------------------------")
					fmt.Println(e.ToString())
				}
			}
			break
		}

	})
}

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
	container := BuildContainer()
	Run(container)
}
