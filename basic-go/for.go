package main

import "fmt"

func main() {
	//counter := 1
	//
	//for counter <= 10 {
	//	fmt.Println("Perulangan ke ", counter)
	//	counter++
	//}

	for i := 0; i < 10; i++ {
		fmt.Println("Perulangan ke", i)
	}

	fmt.Println("Selesai")

	names := []string{"Alex", "John", "Paul"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for i, name := range names {
		fmt.Println("index", i, "=", name)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
