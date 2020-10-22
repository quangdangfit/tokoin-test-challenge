package main

import (
	"bufio"
	"fmt"
	"os"

	"go.uber.org/dig"

	"tokoin/repositories/files"
	"tokoin/schema"
	"tokoin/services"
)

const welcomeMessage = `Type 'quit' to exit at any time. Select search options:
 	* Enter '1' to search.	
 	* Enter '2' to view a list of searchable fields.
`

const chooseModelMessage = `Please choose model you want to search. Select options:
	 * Enter '1' for Organization.
	 * Enter '2' for Ticket.
	 * Enter '3' for User.
`

const inputFieldMessage = `Enter search term: `
const inputValueMessage = `Enter search value: `

var scanner = bufio.NewScanner(os.Stdin)

func Run(container *dig.Container) {
	for {
		fmt.Print(welcomeMessage)
		cmd, quit := readInput()
		if quit {
			break
		}

		if quit = commandHandler(container, cmd); quit {
			break
		}
	}
}

func commandHandler(container *dig.Container, cmd string) bool {
	switch cmd {
	case "1":
		return searchHandler(container)
	case "2":
		showSearchableFields()
		return false
	}

	return false
}

func searchHandler(container *dig.Container) bool {
	var (
		searchModel string
		searchField string
		searchValue string
		quit        bool
	)

	fmt.Print(chooseModelMessage)
	for {
		line, quit := readInput()
		if quit {
			return true
		}

		if line == "1" || line == "2" || line == "3" {
			searchModel = line
			break
		}
		fmt.Print("Please enter '1', '2' or '3':")
	}

	fmt.Print(inputFieldMessage)
	searchField, quit = readInput()
	if quit {
		return true
	}

	fmt.Print(inputValueMessage)
	searchValue, quit = readInput()
	if quit {
		return true
	}

	var results schema.ListResult
	var err error
	container.Invoke(func(
		orgService *services.OrgService,
		ticketService *services.TicketService,
		userService *services.UserService,
	) {
		switch searchModel {
		case "1":
			results, err = orgService.List(searchField, searchValue)
		case "2":
			results, err = ticketService.List(searchField, searchValue)
		case "3":
			results, err = userService.List(searchField, searchValue)
		}

	})
	if err != nil {
		fmt.Println("Error occurred: ", err)
		return false
	}

	fmt.Print("-------------------------------------------")
	fmt.Println(results.ToString())
	fmt.Printf("Found %d record(s).\n", results.Length())

	return false
}

func showSearchableFields() {
	userSearchableFields := []string{"_id", "url", "external_id", "name", "alias", "created_at",
		"active", "verified", "shared", "locale", "timezone", "last_login_at", "email",
		"phone", "signature", "organization_id", "tags", "suspended", "role",
	}
	ticketSearchableFields := []string{"_id", "url", "external_id", "created_at",
		"type", "subject", "description", "priority", "status", "submitter_id",
		"assignee_id", "organization_id", "tags", "has_incidents", "due_at", "via",
	}
	orgSearchableFields := []string{
		"_id", "url", "external_id", "name", "domain_names", "created_at", "details", "shared_tickets", "tags",
	}

	fmt.Println("------------------------------------------")
	fmt.Println("Search user with")
	for _, e := range userSearchableFields {
		fmt.Println(e)
	}

	fmt.Println("------------------------------------------")
	fmt.Println("Search ticket with")
	for _, e := range ticketSearchableFields {
		fmt.Println(e)
	}

	fmt.Println("------------------------------------------")
	fmt.Println("Search organization with")
	for _, e := range orgSearchableFields {
		fmt.Println(e)
	}

	fmt.Println("------------------------------------------")
}

func readInput() (string, bool) {
	scanner.Scan()
	line := scanner.Text()
	if line == "quit" {
		return "", true
	}

	return line, false
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
