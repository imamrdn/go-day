package main

import "fmt"

func getCompleteName() (firstName, lastName string) {
	firstName = "Imam"
	lastName = "Ramadhan"

	return firstName, lastName
}

func main() {
	a, b := getCompleteName()
	fmt.Println(a, b)
}
