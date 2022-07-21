package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	var bookings []string //make empty arrary

	// var can change
	// const can not change
	fmt.Println("Welcome to our", conferenceName, "booking application")
	fmt.Println("Get your tickets here to attend")
	fmt.Println("Tickets is now", conferenceTickets, "Available")
	fmt.Printf("Tickets is now %v Available\n", remainingTickets)
	// printf('%v\n', variable) :

	// loop : go only has for loop
	for {
		var userName string
		var userTickets int // var 변수명 + 형태
		// ask user for their name
		userTickets = 10
		fmt.Printf("userName types %T \n", userName)
		//pointer variable porint to memeory address
		//special variable
		//specail varaible & assign data to memeory address
		fmt.Println("Enter Your userName")
		fmt.Scan(&userName) //fmt.Scan : io method

		fmt.Printf("Hello %v Now we have %v Tickets \n", userName, remainingTickets)
		//Printf -> %v : variable 값을 받아줌
		//       -> %T : type 값을 보내줌

		fmt.Println("how many ticket you wanna book")
		fmt.Scan(&userTickets)
		fmt.Printf("Tickets remaining is now %v \n", remainingTickets)

		// CREATE ARRARY
		// var firstName string
		// var lastName string
		// firstName = "hyunsuk"
		// lastName = "jeong"

		// var bookings [50]string //make empty arrary
		// bookings[0] = firstName + " " + lastName
		//Fixed size // datatypes
		// datatype can not mixed the types

		//Probelm with Array
		// we can not know array size in realworld
		// dynamic array size --> slice
		// var bookings []string //make empty arrary
		bookings = append(bookings, userName)

		fmt.Printf("The whole array %v\n", bookings)
		// fmt.Printf("The first array %v\n", bookings[0])
		// fmt.Printf("The types array %T\n", bookings)
		// fmt.Printf("The length array %v\n", len(bookings))

		firstNames := []string{}           // := -> var firstNames = []sting{}
		for _, booking := range bookings { //range bookings 를 하면 enumerate와 유사
			var names = strings.Fields(booking) // string.Fields python split과 유사
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The whole array %v\n", firstNames)

		if userTickets > remainingTickets {
			fmt.Printf("we only have %v tickets remaining\n", remainingTickets)
			continue
		}

		remainingTickets = remainingTickets - userTickets

		var noTicketRemaining bool = remainingTickets == 0
		if noTicketRemaining {
			fmt.Println("Sorry we are sold out")
			break
		}
	}
}
