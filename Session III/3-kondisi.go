package main

import "fmt"

/* contoh penggunaan kondisi */

func main() {

	var nilai = 7

	/* seleksi kondisi dengan method switch */

	switch nilai {
	case 10:
		fmt.Println("Cakeeeuuuup")
	case 6, 7, 8, 9:
		fmt.Println("lumayan lah")
	default:
		fmt.Println("belajar lagi cuy")
	}

	/* seleksi kondisi dengan method if-else */

	nilai = 10
	if nilai == 10 {
		fmt.Println("Cakeeeuuuup")
	} else if nilai > 6 {
		fmt.Println("lumayan lah")
	} else {
		fmt.Println("belajar lagi cuy")
	}

}
