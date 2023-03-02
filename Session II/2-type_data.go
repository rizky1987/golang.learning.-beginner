package main

import "fmt"

/* contoh kasus jika kita menggabungkan atau menambahkan 2 type data berbeda. */
func main() {

	var variableInt int = 10
	var variableDecimal float64 = 64.9

	// pada baris ini aplikasi akan error
	result := variableInt + variableDecimal

	//solusi
	//result := float64(variableInt) + variableDecimal

	fmt.Println("hasil ", result)

}
