package main

import "fmt"

func main() {

	name := "imam"

	switch name {
	case "imam":
		fmt.Println("I'm imam")
	case "joko":
		fmt.Println("I'm joko")
	default:
		fmt.Println("Hi, what's your name?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	length := len(name)
	switch {
	case length > 5:
		fmt.Println("Nama terlalu panjang")
	case length > 6:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
