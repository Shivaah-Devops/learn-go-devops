package main

import "fmt"

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

	fmt.Printf("Welcomet to %v booking app\n", ConfName)
	fmt.Printf("We have total number of tickets %v and still %v remaining tickets.\n", MaxTickets, remainingtickets)

	var Name string
	var email string
	var tickets int
	fmt.Printf("Enter the Name: ")
	fmt.Scan(&Name)
	fmt.Printf("Enter the email: ")
	fmt.Scan(&email)
	fmt.Println("Enter the number of Tickets: ")
	fmt.Scan(&tickets)

	remainingtickets = remainingtickets - tickets

	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confitmation email at %v\n", Name, tickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingtickets, ConfName)

}
