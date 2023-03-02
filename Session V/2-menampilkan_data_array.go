package main

import "fmt"

/* contoh inisialisasi array dan cara menempilkan data dalam array */

func main() {

	var arrayMobil [4]string
	arrayMobil = [4]string{"bemo", "avanza", "pajero", "toyota"}

	/* menampilkan dengan for */
	for i := 0; i < len(arrayMobil); i++ {
		fmt.Printf("index ke %d mempunya isi %s", i, arrayMobil[i])
	}

	/* menampilkan dengan range */
	for i, dataMobil := range arrayMobil {
		fmt.Printf("index ke %d mempunya isi %s", i, dataMobil)
	}

}
