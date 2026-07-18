package main

import (
	"fmt"
	"strings"
)

func main() {
	// //fmt.Println("Welcome to booking app!")
	// //fmt.Println("here's the place to collect tickets")

	// Confname := "GO Conf"
	// const noofTickets = 50
	// var reminaingtickets = 50

	// fmt.Printf("welcome to %v booking app!\n", Confname)
	// fmt.Printf("We have total of %v tickets and %v are still remaining.\n", noofTickets, reminaingtickets)
	// //fmt.Println("Name of the Conference: ", Confname)
	// //fmt.Println("Get your tickets here to attend")

	// //fmt.Println("Total no. of tickers:", noofTickets, "remaining tickets:", reminaingtickets)
	// fmt.Printf("Get your tickets from here to attend\n")

	// var userName string
	// var ticket int

	// userName = "Siva"
	// ticket = 2
	// fmt.Printf("user %v booked %v tickets\n", userName, ticket)
	// fmt.Printf("check")
	// fmt.Scan(userName)

	ConfName := "GO live conf"
	const MaxTickets int = 50
	var remainingtickets int = 50

	//var bookings [50]string Array declaration
	var bookings []string //slice declaration

	fmt.Printf("Welcomet to %v booking app\n", ConfName)
	fmt.Printf("We have total number of tickets %v and still %v remaining tickets.\n", MaxTickets, remainingtickets)

	//introducing loops
	for {
		var firstName string
		var lastName string
		var email string
		var tickets int
		fmt.Printf("Enter the FirstName: ")
		fmt.Scan(&firstName)
		fmt.Printf("Enter the LasttName: ")
		fmt.Scan(&lastName)
		fmt.Printf("Enter the email: ")
		fmt.Scan(&email)
		fmt.Printf("Enter the number of Tickets: ")
		fmt.Scan(&tickets)

		// if tickets > remainingtickets {
		// 	fmt.Printf("Sorry, we only have %v tickets remaining, you can't book %v tickets\n", remainingtickets, tickets)
		// 	break
		// }

		//validating inputs

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketnumber := tickets > 0 && tickets <= remainingtickets

		if isValidName && isValidEmail && isValidTicketnumber {
			remainingtickets = remainingtickets - tickets
			//bookings[0] = Name
			bookings = append(bookings, firstName+" "+lastName)
			fmt.Printf("Thank you %v for booking %v tickets. You will receive a confitmation email at %v\n", firstName+" "+lastName, tickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingtickets, ConfName)

			// fmt.Printf("The Array type is %T\n", bookings)      Retrieving the array or slice is same
			// fmt.Printf("The array length is %v\n", len(bookings))
			// fmt.Printf("The Array elements are %v\n", bookings)

			//get only first names
			firstnames := []string{}
			for _, booking := range bookings {
				strings := strings.Fields(booking)
				firstnames = append(firstnames, strings[0])
			}

			// fmt.Printf("The Slice type is %T\n", bookings)
			// fmt.Printf("The slice length is %v\n", len(bookings))
			// fmt.Printf("The slice elements are %v\n", bookings)

			fmt.Printf("The first names of bookings are :%v\n", firstnames)

			nomoretickets := remainingtickets == 0

			if nomoretickets {
				fmt.Println("No more tickets available.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your first name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Your email address is missing @sign")
			}
			if !isValidTicketnumber {
				fmt.Println("Your ticket number is invalid")
			}
			//fmt.Printf("Your input is invalid, please try again\n")
			//fmt.Printf("Sorry, we only have %v tickets remaining, you can't book %v tickets\n", remainingtickets, tickets)
		}

	}

}
