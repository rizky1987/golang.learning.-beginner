package main

import "fmt"

/* contoh inisialisasi array dan cara menempilkan data dalam array */

func main() {

	var arrayMobil [4]string
	arrayMobil = [4]string{"bemo", "avanza", "pajero", "toyota"}

	fmt.Println("jumlah data dalam array : ", len(arrayMobil))

	sliceMobil := []string{"bemo", "avanza", "pajero", "toyota"}
	sliceMobil = append(sliceMobil, "daihatsu")
	fmt.Println(sliceMobil)

}
