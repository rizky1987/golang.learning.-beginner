package main

import (
	"fmt"
)

/* contoh penulisan variable */
func main() {

	// === mulai penulisan variable dengan type data
	var namaDepan string = "radya"

	var namaBelakang string
	namaBelakang = "digital"

	fmt.Printf("halo %s %s!\n", namaDepan, namaBelakang)
	// === akhir penulisan variable denngan type data

	// === mulai penulisan 2 variable atau lebih dengan type data dalam 1 baris code

	var namaDepanLagi, namaBelakangLagi string = "rizky", "unyu"
	fmt.Printf("halo %s %s!\n", namaDepanLagi, namaBelakangLagi)

	// === akhir penulisan 2 variable atau lebih dalam 1 baris code

}
