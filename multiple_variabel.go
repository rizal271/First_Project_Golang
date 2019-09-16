package main

import "fmt"
func main(){
	var firstName, lastName string = "Rizal", "Rohman"

	fmt.Printf("halo Rizal Rohman!\n")
	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println("halo", firstName, lastName+"!")
}