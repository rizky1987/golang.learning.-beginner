package main

import "fmt"

/* contoh penggunaan perulangan versi 2 */

func main() {

	i := 2

	for {
		fmt.Println("angka ", i)
		i++

		if i == 3 {
			break
		}
	}
}
