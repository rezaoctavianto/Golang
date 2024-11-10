package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("Only %v tickets are available. %v tickets left.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your %v tickets here!\n", conferenceName)

    for  remainingTickets > 0 {
        var firstName string
        var lastName string
        var email string
        var userTickets uint

        fmt.Print("Enter your first name: ")
        fmt.Scan(&firstName)

        fmt.Print("Enter your last name: ")
        fmt.Scan(&lastName)

        fmt.Print("Enter your email: ")
        fmt.Scan(&email)

        fmt.Print("Enter number of ticket: ")
        fmt.Scan(&userTickets)

        if userTickets < remainingTickets {
            remainingTickets = remainingTickets - userTickets
            bookings = append(bookings, firstName+" "+lastName)

            // fmt.Printf("The whole slice: %v\n", bookings)
            // fmt.Printf("The first value: %v\n", bookings[0])
            // fmt.Printf("slice type: %T\n", bookings)
            // fmt.Printf("slice length: %v\n", len(bookings))

            fmt.Printf("Thank you %v for booking %v tickets. We send you a booking confirmation to %v\n", bookings[0], userTickets, email)
            fmt.Printf("%v remaining tickets available\n", remainingTickets)

            firstNames := []string{}
            for _, booking := range bookings {
                var names = strings.Fields(booking)
                firstNames = append(firstNames, names[0])

            }
            fmt.Printf("All of our bookings: %v\n", firstNames)

            if remainingTickets == 0 {
                fmt.Printf("Unfortunately our tickets is sold out, Thank you for your enthusiasm. see you on the next conference")
                break
            }
        } else {
            fmt.Printf("We only have %v tickets remaining\n", remainingTickets)
            continue
        }
    }
}
