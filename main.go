package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)
var conferenceName = "Go Conference"
const conferenceTickets int = 50
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)


func main() {

    greetUser()

	// fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	// fmt.Printf("Only %v tickets are available. %v tickets left.\n", conferenceTickets, remainingTickets)
	// fmt.Printf("Get your %v tickets here!\n", conferenceName)

    for  remainingTickets > 0 {

        firstName, lastName, email, userTickets := getUserInput()

        isNameValid, isEmailValid, isTicketsValid := helper.ValidateInput(firstName, lastName, email, userTickets, remainingTickets)

        if isNameValid && isEmailValid && isTicketsValid {
            bookTicket(userTickets, firstName, lastName, email)
            // print names
            firstNames := getFirstName() 
            fmt.Printf("All of our bookings: %v\n", firstNames)

            if remainingTickets == 0 {
                fmt.Printf("Unfortunately our tickets is sold out, Thank you for your enthusiasm. see you on the next conference")
                break
            }
        } else {
            if !isNameValid {
                fmt.Printf("Your name is to short\n")
            } 
            if !isEmailValid {
                fmt.Printf("email %v are not available\n", email)
            } 
            if !isTicketsValid {
                fmt.Printf("Number of ticket are not available\n")    
            }

            fmt.Printf("Your input data are invalid. Try again!\n")
            continue
        }
    }

    // city := "London"

    // switch city {
    //     case "New York", "Hong Kong":

    //     case "Berlin", "Mexico":
        
    //     case "Mexico City":
        
    // }

}

func greetUser() {
    fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("Only %v tickets are available. %v tickets left.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your %v tickets here!\n", conferenceName)
}

func getFirstName() []string {
    firstNames := []string{}
    for _, booking := range bookings {
        firstNames = append(firstNames, booking["firstName"])
    }
    return firstNames
}

// func validateInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool){
//     isNameValid := len(firstName) >= 2 && len(lastName) >= 2  
//     isEmailValid := strings.Contains(email, "@")
//     isTicketsValid := userTickets > 0 && userTickets <= remainingTickets

//     return isNameValid, isEmailValid, isTicketsValid
// }

func getUserInput() (string, string, string, uint){
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

    return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
    remainingTickets = remainingTickets - userTickets
    //map
    var userData = make(map[string]string)
    userData["firstName"] = firstName
    userData["lastName"] = lastName
    userData["email"] = email
    userData["userTicket"] = strconv.FormatUint(uint64(userTickets), 10)

    bookings = append(bookings, userData)
    fmt.Printf("List of Map %v\n", bookings)
    // fmt.Printf("The whole slice: %v\n", bookings)
    // fmt.Printf("The first value: %v\n", bookings[0])
    // fmt.Printf("slice type: %T\n", bookings)
    // fmt.Printf("slice length: %v\n", len(bookings))

    fmt.Printf("Thank you %v for booking %v tickets. We send you a booking confirmation to %v\n", firstName, userTickets, email)
    fmt.Printf("%v remaining tickets available for %v\n", remainingTickets, conferenceName)
}