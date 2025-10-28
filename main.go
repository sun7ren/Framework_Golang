package main

import (
	"fmt"
	"strings"
)

func main() {
	boothName := "Thematics Booth"
	const maxQueue = 50
	var remainingQueue = 50

	var bookings = []string{}
	var firstNames = []string{}

	fmt.Printf("Welcome to %v photobooth\n", boothName)
	fmt.Printf("We have a total of %v tickets, and a remaining of %v tickets\n", maxQueue, remainingQueue)
	fmt.Println("Get your ticket queue here to start")

	for {
		var firstName string
		var email string
		var amountOfTickets int

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter the amount of Tickets: ")
		fmt.Scan(&amountOfTickets)

		fmt.Printf("User %v booked %v queue tickets.\n", firstName, amountOfTickets)

		if remainingQueue >= amountOfTickets && amountOfTickets > 0 {
			remainingQueue = remainingQueue - amountOfTickets
			fmt.Printf("SUCCESS! There are %v remaining queue tickets.\n", remainingQueue)

			bookings = append(bookings, firstName+" "+email)
			firstNames = append(firstNames, firstName)

		} else {
			if amountOfTickets <= 0 {
				fmt.Println("Failed transaction. Please enter a valid ticket amount.")
			} else {
				fmt.Printf("Failed transaction. We only have %v tickets remaining.\n", remainingQueue)
			}
		}

		fmt.Printf("The first names of current bookings are: %v\n", strings.Join(firstNames, ", "))
		fmt.Printf("The current bookings are: %v\n", bookings)

		if remainingQueue == 0 {
			fmt.Println("All tickets are sold out! Closing the booth.")
			break
		}
	}
}
