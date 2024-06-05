package main

import "fmt"

func main() {

	//menuliskan menggunakan var
	var name = "Imam Ramadhan"
	fmt.Println(name)

	//tanpa menggunakan var
	age := 20
	fmt.Println(age)

	//multiple variable
	var (
		firstName = "Imam"
		lastName  = "Ramadhan"
	)

	fmt.Println(firstName, lastName)
}
