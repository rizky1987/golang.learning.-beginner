package main

import "fmt"

/* contoh inisialisasi array dan cara menempilkan data dalam array */

func main() {

	//array 1 dimensi
	var arrayMobil [4]string

	arrayMobil = [4]string{"bemo", "avanza", "pajero", "toyota"}

	fmt.Println(arrayMobil[0]) //bemo
	fmt.Println(arrayMobil[1]) //avanza
	fmt.Println(arrayMobil[2]) //pajero
	fmt.Println(arrayMobil[3]) //toyota
	fmt.Println(arrayMobil[4]) // aplikasi error

	//array 1 multidimensi
	var arrayMultidimensi = [2][3]int{[3]int{3, 2, 1}, [3]int{6, 4, 5}}
	fmt.Println(arrayMultidimensi[0][0]) //1

}
