package main

import "strings";

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	// in go you can return any number of values
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketCount := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketCount
}

