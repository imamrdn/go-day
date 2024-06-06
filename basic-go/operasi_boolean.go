package main

import "fmt"

func main() {

	nilaiAkhir := 90
	nilaiAbsensi := 80

	lulusNilaiAkhir := nilaiAkhir > 80
	lulusAbsensi := nilaiAbsensi > 80

	lulus := lulusNilaiAkhir && lulusAbsensi
	fmt.Println(lulus)
}
