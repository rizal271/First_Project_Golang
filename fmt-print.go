package main

import "fmt"

func main() {
	var firstName string = "Rizal"

	var lastName string
	lastName = "Rohman"

	fmt.Printf("halo Rizal Rohman!\n")
	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println("halo", firstName, lastName+"!")
}