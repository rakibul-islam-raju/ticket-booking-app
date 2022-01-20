package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const conferenceTickets uint = 50
var conferanceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main()  {
	// greet user
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		// vallidation
		isValidName, isValidEmail, isValidTicketNumber :=  validateUserInputs(firstName, lastName, email, userTickets, remainingTickets)
		
		// check the number of tickets user wants to book
		if isValidName && isValidEmail && isValidTicketNumber  {
			// book ticket
			bookTicket(userTickets, firstName, lastName, email)

			// send the ticket
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			// first names of bookings
			firstNames := printFirstName()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			

			// check if all tickets are booked
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out! Please come back next year.")
				break
			}
		}else{
			if !isValidName {
				fmt.Println("First name or last name is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email address doesn't contain @ sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets are not available.")
			}
		}
	}
	wg.Wait()
}

func getUserInput() (string, string, string, uint) {
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

		return firstName, lastName, email, userTickets
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", conferanceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
}

func printFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func validateUserInputs(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	// map for userInfo
	userData := UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	// ticket booking logic
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)

	// confirmation message
	fmt.Printf("Thank you %v %v for booking %v tickets. You will get a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferanceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Printf("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Printf("#################")
	wg.Done()
}