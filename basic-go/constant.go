package main

import "fmt"

func main() {

	//variable yang tidak dapat diubah lagi atau reassign
	const firstName = "Imam Aja"

	fmt.Println(firstName)

	const (
		namaDepan    = "Imam"
		namaBelakang = "Ramadhan"
	)

	fmt.Println(namaDepan, namaBelakang)

}
