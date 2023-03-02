package main

import "fmt"

/* contoh penggunaan logika variable */

func main() {

	var nilaiA string = "radya"
	var nilaiB string = "digital"

	/* jika nilaiA bukan 'radya' dan nilaiB bukan digital */
	var testLogika1 = (nilaiA != "radya") && (nilaiB != "digital")
	fmt.Println("hasil logika 1 ", testLogika1)

	/* jika nilaiA sama dengan 'radya' atau nilaiB sama dengan radya */
	var testLogika2 = (nilaiA == "radya") || (nilaiB == "radya")
	fmt.Println("hasil logika 2 ", testLogika2)

}
