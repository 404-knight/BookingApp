package main

import (
	"fmt"
)

const NumberOfTickets = 50
var remainingTickets = 50
var conferenceName = "Go Conference"
var booking [] string

func main () {
	greetUser()

	firstName, lastName, email, userTickets := userInput()

	// Add both the first name and last names to booking slice
	userName := firstName + " " + lastName
  booking = append(booking, userName)

	fmt.Printf("You will receive email confirmation at %v for %v tickets.\n", email, userTickets)
	fmt.Printf("Here's our booking list %v:\n", booking)	

}

func greetUser() {
	fmt.Printf("Welcome to our ticket booking application for %v.\n", conferenceName)
	fmt.Printf("We have %v tickets remaining of %v\n", remainingTickets, NumberOfTickets)
	fmt.Println("Enter your details to continue:")
}

func userInput() (string, string, string, uint) {

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

	fmt.Println("Enter the number of tickets you wish to purchase:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
