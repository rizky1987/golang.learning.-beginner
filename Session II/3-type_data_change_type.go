package main

import "fmt"

/* contoh kasus jika kita mengisi data yang salah */
func main() {

	var variableInt int = 10

	// pada baris ini aplikasi akan error
	variableInt = 64.9

	fmt.Println("hasil ", variableInt)
}
