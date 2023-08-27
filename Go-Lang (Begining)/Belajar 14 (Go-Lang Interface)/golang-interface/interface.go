package main

import (
	"errors"
	"fmt"
)

//======================================== Interface ========================================
// /*
// Interface = Tipe data abstract
// 			sebuah interface berisi definisi method
// 			biasanya digunakan sebagai kontrak
// // angka menandakan urutan prosesnya
// */
// //6.
// type HasName interface { //membuat interface dengan nama HasName
// 	GetName() string //harus ada kontrak yang mempunyai function GetName()
// }

// //4.
// func sayHallo(hasName HasName) { //interface bisa digunakan sebagai parameter seperti struct
// 	//5.
// 	fmt.Println("Hallo", hasName.GetName()) //untuk memanggil GetName harus disebutkan dulu nama variablenya seperti contoh di pinggir //9.
// }

// type Person struct { //2.
// 	nama string
// }

// //bisa dimenerapkan lebih dari 1 struct
// type Animal struct {
// 	nama string
// }

// func (orang Person) GetName() string { //7.
// 	return orang.nama //8.
// }

// //bisa dimenerapkan lebih dari 1 struct
// func (hewan Animal) GetName() string {
// 	return hewan.nama
// }

// func main() {
// 	//implementasi interface
// 	persin := Person{"Rayandra"} //1.
// 	sayHallo(persin)             //3.

// 	hewan := Animal{"Gajah"}
// 	sayHallo(hewan)
// }

//======================================== Interface Kosong ========================================
// /*
// Interface kosong = interface yang tidak memiliki deklarasi method satupun, dan akan membuat semua tipe data menjadi implementasinya
// Penggunaan interface kosong : 	fmt.Println(a ...interface{})
// 								panic(v interface{})
// 								recover() interface{}
// 								dll
// */

// func ups(i int) interface{} { //contoh interface kosong //dapat mengambalikan tipe data apa saja tidak terpatik 1 tipe data
// 	//data interface kosong
// 	if i == 1 {
// 		return 1
// 	} else if i == 2 {
// 		return true
// 	} else {
// 		return "ups"
// 	}
// 	//return 1
// 	//return true
// 	// return "Ups"
// }

// func main() {
// 	var kosong interface{} = ups(1) //tipe data diganti menjadi interface{}
// 	// kosong := ups(1)
// 	fmt.Println(kosong)
// }

//======================================== Nil ========================================
// /*
// nil = data kosong
// 	hanya dapat digunakan di beberapa tipe data seperti :	interface
// 															function
// 															map
// 															slice
// 															pointer
// 															chanel

// nil digunakan untuk pengecekan data kosong
// */

// func newMap(name string) map[string]string {
// 	if name == "" {
// 		return nil
// 	} else {
// 		return map[string]string{
// 			"name": name,
// 		}
// 	}
// }

// func main() {
// 	person := newMap("Wandi")
// 	// var person map[string]string = nil
// 	fmt.Println(person)

// 	nama := newMap("")
// 	if nama == nil {
// 		fmt.Println("Data Kosong")
// 	} else {
// 		fmt.Println(nama)
// 	}
// }

//======================================== Interface Error ========================================
/*
Kontrak untuk membuat error
*/

type error interface { //deklarasi interface
	Error() string
}

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi Dengan 0") //menampilkan pesan error
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	// var contoh error = errors.New("Ups Error")
	// fmt.Println(contoh.Error())

	hasil, err := pembagian(1, 0)

	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Hasil", err.Error())
	}
}
