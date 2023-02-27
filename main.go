package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
	"sync"
	"time"
)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	noOfTickets int
}

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets = 50
var bookings = make([]map[string]string, 0)
var customBookings = make([]UserData, 0)

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := requestUserInput()

		bookTicketsUsingStruct(firstName, lastName, email, userTickets)

		wg.Add(1)
		go sendEmail(firstName, lastName, email, userTickets)

		if remainingTickets <= 0 {
			fmt.Print("Our conference is booked out. Come back next year.")
			wg.Wait()
			break
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v are remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here.")
}

func requestUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Print("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Please enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("How many tickets do you want? ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(firstName string, lastName string, email string, userTickets int) {
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["noOfTickets"] = strconv.FormatInt(int64(userTickets), 10)

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings user data: %v\n", bookings)

	fmt.Printf("The first names of the bookings are: %v\n", helper.GetFirstNames(bookings))

	fmt.Printf(
		"Thank you %v %v for booking %v tickets. "+
			"You will receive a confirmation email at %v\n",
		firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func bookTicketsUsingStruct(firstName string, lastName string, email string, userTickets int) {
	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		noOfTickets: userTickets,
	}

	remainingTickets = remainingTickets - userTickets
	customBookings = append(customBookings, userData)
	fmt.Printf("List of bookings user data: %v\n", customBookings)

	fmt.Printf(
		"Thank you %v %v for booking %v tickets. "+
			"You will receive a confirmation email at %v\n",
		firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendEmail(firstName string, lastName string, email string, userTickets int) {
	time.Sleep(50 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#####################################################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#####################################################")
	wg.Done()
}
