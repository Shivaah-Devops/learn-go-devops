package main

import "fmt"

func main() {
	//fmt.Println("Welcome to booking app!")
	//fmt.Println("here's the place to collect tickets")

	Confname := "GO Conf"
	const noofTickets = 50
	var reminaingtickets = 50

	fmt.Printf("welcome to %v booking app!\n", Confname)
	fmt.Printf("We have total of %v tickets and %v are still remaining.\n", noofTickets, reminaingtickets)
	//fmt.Println("Name of the Conference: ", Confname)
	//fmt.Println("Get your tickets here to attend")

	//fmt.Println("Total no. of tickers:", noofTickets, "remaining tickets:", reminaingtickets)
	fmt.Printf("Get your tickets from here to attend\n")

	var userName string
	var ticket int

	userName = "Siva"
	ticket = 2
	fmt.Printf("user %v booked %v tickets\n", userName, ticket)
	fmt.Printf("check")

}
