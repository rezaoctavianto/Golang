package main

import "fmt"

func main() {
    var conferenceName = "Go Conference"
    const conferenceTickets = 50
    var remainingTickets uint = 50
    var bookings [50]string


    fmt.Printf("Welcome to %v booking application!\n", conferenceName)
    fmt.Printf("Only %v tickets are available. %v tickets left.\n", conferenceTickets, remainingTickets)
    fmt.Printf("Get your %v tickets here!\n", conferenceName)


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

    remainingTickets = remainingTickets - userTickets
    bookings[0] = firstName + " " + lastName

    fmt.Printf("Thank you %v for booking %v tickets. We send you a confirmation email to %v\n", firstName, userTickets, email)
    fmt.Printf("%v remaining tickets available", remainingTickets)
}