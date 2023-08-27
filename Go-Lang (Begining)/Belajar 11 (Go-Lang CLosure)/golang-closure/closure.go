package main

import "fmt"

/*
Closure harus digunakan dengan bijak saat membuat aplikasi
bisa disebut dengan berinteraksi dengan data di sekitarnya
*/

func main() {
	counter := 0

	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment() //memanggil dan menambahkan nilai 1 pada counter
	increment() //memanggil dan menambahkan nilai 1 pada counter
	fmt.Println(counter)
}
