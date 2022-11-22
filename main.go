package main

import (
	"fmt"
	"sync"
	"time"
)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

const conferenceTickets uint = 50
const conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)
var wg = sync.WaitGroup {}

func main() {
	greetUser()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketCount := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketCount {
		bookTicket(userTickets, firstName, lastName, email)
		
		// spin off new thread: GoRoutine
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The first names of the bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out!")
		}
	} else {
		if !isValidName {
			fmt.Println("Invalid name, please try again")
		}
		if !isValidEmail {
			fmt.Println("Invalid email address")
		}
		if !isValidTicketCount {
			fmt.Println("Number of tickets you entered is invalid")
		}
	}
	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("list of bookings %v", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) string {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##############################")
	fmt.Printf("Sending ticket:\n %v to email address:\n %v\n", ticket, email)
	fmt.Println("##############################")
	wg.Done()
	return ticket
}
