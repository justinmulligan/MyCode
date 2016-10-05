package main

import "fmt"

func main() {

	// First some variables
	var usersName string
	var usersAge int
	computersAge := 10
	var differenceInAge int

	// 2 Questions to ask

	fmt.Print("Hello, What is your name? ")
	fmt.Scan(&usersName)
	fmt.Println("Hello", usersName)

	fmt.Print("How old are you? ")
	fmt.Scan(&usersAge)

	// Calculate age difference between user and compuer

	differenceInAge = usersAge - computersAge

	// If statement needed when the users age is lower than the computers age

	if usersAge <= computersAge-1 {
		differenceInAge *= -1
		fmt.Println("I am", differenceInAge, "years older than you")
	} else {
		fmt.Println("You are", differenceInAge, "years older than me")
	}
}
