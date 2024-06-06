package main

import "fmt"

func main() {

	var names [3]string

	names[0] = "Alice"
	names[1] = "Bob"
	names[2] = "Cat"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	names2 := [...]string{"Alice", "Bob", "Cat"}

	fmt.Println(names2)
	fmt.Println(len(names2))
}
