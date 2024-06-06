package main

import "fmt"

func main() {

	names := [...]string{"imam", "rama", "dhan", "imamrama", "imamramad", "imamramadhan"}

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[2:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Minggu"}

	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Imam"
	newSlice[1] = "Rama"
	//newSlice[2] = "dhan" // error, harusnya menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice)) // len untuk menghitung panjang data
	fmt.Println(cap(newSlice)) // cap atau capacity untuk mengetahui banyak kapasitas data

	newSlice2 := append(newSlice, "Imam")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]                                      // asal slice atau asal data
	toSlice := make([]string, len(fromSlice), cap(fromSlice)) //slice baru

	copy(toSlice, fromSlice) //menyalin ke slice mana, dari slice mana

	fmt.Println(toSlice)
	fmt.Println(fromSlice)

	//MEMBEDAKAN SLICE DAN ARRAY

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
