package main

import (
	"fmt"
	"sync"
	"time"
)

// package level variable
const concertTickets uint = 50

var concertName = "MONSTA X WORLD TOUR ‘WE ARE HERE’ IN BANGKOK"
var remainingTickets uint = concertTickets //should never be negative
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	fmt.Printf("concertTickets is %T, remainingTickets is %T, concertName is %T\n", concertTickets, remainingTickets, concertName)

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketsNumber := validateUserInput(firstName, lastName, email, userTickets)

	if isValidName && isValidEmail && isValidTicketsNumber {

		bookTicket(userTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("-----------------------------------------------")
			fmt.Println("Our concert is booked out. Come back next year.")
			fmt.Println("-----------------------------------------------")
			// break
		}

	} else {
		if !isValidName {
			fmt.Println("The firstname or lastname is too short.")
		}
		if !isValidEmail {
			fmt.Println("The email doesn't contain '@' sign.")
		}
		if !isValidTicketsNumber {
			fmt.Println("The number of tickets you entered is invalid.")
		}
		// continue
	}

	wg.Wait()
}

func greetUsers() {
	fmt.Printf("welcome to %v booking application\n", concertName)
	fmt.Println("======================================================")
	fmt.Printf("We have total of %v tickets and %v are still available.\n", concertTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets

	// create a map for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("the booking is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, concertName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("######################")
	fmt.Printf("Sending ticket:\n %v\n to email address %v\n", ticket, email)
	fmt.Println("######################")
	wg.Done()
}
