package main

import (
	"fmt"
	"strings"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		//isValidCity := city =="Singapore" || city =="London"
		//isInvalidCity := city != "Singapore" && city !="London"- if invalidcity not equal to singapore and london.
		//!isValidCity - if value is true, make it false.

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)
			sendTicket(userTickets, firstName, lastName, email)

			//call function print first names
			firstNames := getFirstNames()
			fmt.Printf("The first names of booking are: %v\n", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("email adrress you entered does not contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets is invalid")
			}
		}
	}

	// city:= "London"
	// switch city{
	// case "New York", "Nairobi":
	// 	//execute code for booking New York conference tickets
	// case "Singapore ":
	// 	//execute code for booking New York conference tickets
	// case "London":
	// 	//execute code for booking New York conference tickets
	// case "Berlin":
	// 	//execute code for booking New York conference tickets
	// case "Mexico City", "Hong Kong":
	// 	//execute code for booking New York conference tickets
	// default:
	// 	fmt.Print("NO VALID CITY SELECTED")
	// }
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
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

	//ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter your userTickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	//send confirmation of bookinggo run
	fmt.Printf("Dear %v %v, You have successfully booked %v tickets for the %v. Please check %v for confirmation\n", firstName, lastName, userTickets, conferenceName, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("########")
	fmt.Printf("Sending ticket %v \nto email address %v\n", ticket, email)
	fmt.Println("########")
}
