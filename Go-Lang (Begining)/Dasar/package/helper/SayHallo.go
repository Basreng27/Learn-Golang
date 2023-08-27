package helper //untuk menandakan ini adalah package. nama package itu sama seperti nama folder tempat file disimpan

import "fmt"

//isi dari package
//untuk pemanggilannya bisa di cek di package.go
func SayHello(nama string) { //function
	fmt.Println("Hallo Nama Saya : ", nama)
}

/*
Catatan :
			Jika membuat function yang sama pada beda file dan folder sama maka akan terjadi error, solusi jangan membuat nama function yang sama!
*/
