package main

import (
	"fmt"
    "time"
    "sync"
)
var conferenceName = "Go Conference"
const conferenceTickets int = 50
var remainingTickets uint = 50
var bookings = make([]Datas, 0)

type Datas struct {
    firstName string
    lastName string
    email string
    userTickets uint
} 

var wg = sync.WaitGroup{}

func main() {

    greetUser()

	// fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	// fmt.Printf("Only %v tickets are available. %v tickets left.\n", conferenceTickets, remainingTickets)
	// fmt.Printf("Get your %v tickets here!\n", conferenceName)

    firstName, lastName, email, userTickets := getUserInput()

    isNameValid, isEmailValid, isTicketsValid := validateInput(firstName, lastName, email, userTickets)

    if isNameValid && isEmailValid && isTicketsValid {
        bookTicket(userTickets, firstName, lastName, email)

        wg.Add(1)
        go sendTicket(userTickets, firstName, lastName, email)
        
        // print names
        firstNames := getFirstName() 
        fmt.Printf("All of our bookings: %v\n", firstNames)

        if remainingTickets == 0 {
            fmt.Printf("Unfortunately our tickets is sold out, Thank you for your enthusiasm. see you on the next conference")
            // break
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
        // continue
    } 

    // city := "London"

    // switch city {
    //     case "New York", "Hong Kong":

    //     case "Berlin", "Mexico":
        
    //     case "Mexico City":
        
    // }
    wg.Wait() 

}

func greetUser() {
    fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("Only %v tickets are available. %v tickets left.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your %v tickets here!\n", conferenceName)
}

func getFirstName() []string {
    firstNames := []string{}
    for _, booking := range bookings {
        firstNames = append(firstNames, booking.firstName)
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
    var userData = Datas {
        firstName: firstName,
        lastName: lastName,
        email: email,
        userTickets: userTickets,
    }
    // userData["firstName"] = firstName
    // userData["lastName"] = lastName
    // userData["email"] = email
    // userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

    bookings = append(bookings, userData)
    fmt.Printf("List of Map %v\n", bookings)
    // fmt.Printf("The whole slice: %v\n", bookings)
    // fmt.Printf("The first value: %v\n", bookings[0])
    // fmt.Printf("slice type: %T\n", bookings)
    // fmt.Printf("slice length: %v\n", len(bookings))

    fmt.Printf("Thank you %v for booking %v tickets. We send you a booking confirmation to %v\n", firstName, userTickets, email)
    fmt.Printf("%v remaining tickets available for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
    time.Sleep(10 * time.Second)
    var ticket = fmt.Sprintf("%v tickets are sent for %v %v", userTickets, firstName, lastName)
    fmt.Println("##################")
    fmt.Printf("Sending ticket:\n%v\nto email address %v\n", ticket, email)
    fmt.Println("######################")
    wg.Done()
}