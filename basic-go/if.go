package main

import "fmt"

func main() {

	name := "Imam"

	if name == "Imam" {
		fmt.Println("Hello Imam")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Hello, boleh kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("name is too long")
	} else {
		fmt.Println("name is too short")
	}

}
