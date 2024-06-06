package main

import "fmt"

func sayHello() {
	fmt.Println("hello world")
}

func person(firstName string, lastName string) {
	fmt.Println(firstName, lastName)
}

func getHello(name string) string {
	return "Hello " + name
}

func getFullName() (string, string) {
	return "Mr. Jack", "Jack"
}

func getFullName2() (string, string) {
	return "imam", "rama"
}

func main() {
	sayHello()

	person("Imam", "Ramadhan")

	result := getHello("imam")
	fmt.Println(result)

	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	firstName, _ = getFullName2()
	fmt.Println(firstName)
}
