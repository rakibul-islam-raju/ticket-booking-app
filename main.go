package main

import (
	"fmt"
	"strings"
)

func main()  {
	conferanceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application.\n", conferanceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// ask the user
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		// ticket booking logic
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName + " " + lastName)

		// confirmation message
		fmt.Printf("Thank you %v %v for booking %v tickets. You will get a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferanceName)

		// get users firstNames
		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		// first names of bookings
		fmt.Printf("The first names of bookings are: %v\n", firstNames)
	}
}