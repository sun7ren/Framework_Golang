package main

import "fmt"

func main() {
	boothName := "Thematics Booth"
	const maxQueue = 50
	var remainingQueue = 50

	fmt.Printf("Welcome to %v photobooth", boothName)
	fmt.Printf("We have a total of %v tickets, and a remaining of %v tickets", maxQueue, remainingQueue)
	fmt.Println("Get your ticket queue here to start")

	var firstName string
	var email string
	var amountOfTickets int
	var bookings [50]string

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter the amount of Tickets: ")
	fmt.Scan(&amountOfTickets)

	bookings[0] = firstName + " " + email

	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Array type: %T\n", bookings)
	fmt.Printf("Array length: %v\n", len(bookings))

	fmt.Printf("User %v booked %v queue tickets.\n", firstName, amountOfTickets)

	if remainingQueue >= amountOfTickets {
		remainingQueue = remainingQueue - amountOfTickets
		fmt.Printf("There are %v remaining queue tickets", remainingQueue)
	} else {
		fmt.Printf("Failed transaction.")
	}
}
