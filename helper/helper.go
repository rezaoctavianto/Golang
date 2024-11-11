package helper

import "strings"

func ValidateInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool){
    isNameValid := len(firstName) >= 2 && len(lastName) >= 2  
    isEmailValid := strings.Contains(email, "@")
    isTicketsValid := userTickets > 0 && userTickets <= remainingTickets

    return isNameValid, isEmailValid, isTicketsValid
}
