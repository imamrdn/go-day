package main

import "fmt"

func main() {

	type noKTP string

	var ktpImam noKTP = "11111111111"

	var contoh string = "22222222222"
	var contohKtp noKTP = noKTP(contoh)

	fmt.Println(ktpImam)
	fmt.Println(contohKtp)
}
