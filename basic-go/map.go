package main

import "fmt"

func main() {

	person := map[string]string{
		"name":   "John",
		"course": "Go",
	}

	fmt.Println(person["name"])
	fmt.Println(person["course"])
	fmt.Println(person)

	book := make(map[string]string)
	book["name"] = "Jack"
	book["course"] = "Go"

	fmt.Println(book)

	delete(book, "course")

	fmt.Println(book)

}
