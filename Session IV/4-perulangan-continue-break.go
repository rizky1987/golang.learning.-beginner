package main

import "fmt"

/* contoh penggunaan perulangan versi 2 */

func main() {

	for i := 0; i <= 10; i++ {
		fmt.Println("awal iterasi : ", i)
		fmt.Println("Angka", i)
		if i == 3 {
			fmt.Println("akhir iterasi dikarenakan perintah 'continue' ", i)
			fmt.Println("")
			continue
		}
		if i == 7 {
			fmt.Println("akhir iterasi dikarenakan perintah 'break' ", i)
			fmt.Println("")
			break
		}
		fmt.Println("akhir iterasi normal ", i)
		fmt.Println("")
	}
}
