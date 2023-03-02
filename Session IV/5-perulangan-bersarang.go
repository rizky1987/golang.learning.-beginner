package main

import "fmt"

/* contoh penggunaan perulangan versi 2 */

func main() {

	for i := 0; i < 5; i++ {

		fmt.Print(i)
		for j := 0; j < 5; j++ {
			fmt.Print(j)
		}
		fmt.Println("")
	}

}
