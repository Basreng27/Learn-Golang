package main

import "fmt"

/*
untuk merubah tipe data seperti yang diinginan
sering digunakan ketika bertemu interface kosong
*/

func rand() interface{} {
	return 2
}

func main() {
	// result := rand()                //memasukan function rank kedalam variable
	// resultString := result.(string) //merubah isi tipe data kembalian menjadi string
	// fmt.Println(resultString)

	//error
	// resultInt := result.(int) //panic //meribah menjadi int
	// fmt.Println(resultInt)

	//menggunakan switch
	resultt := rand()
	switch value := resultt.(type) {
	case string:
		fmt.Println("String", value)

	case int:
		fmt.Println("Integer", value)

	default:
		fmt.Println("Tidak diketahui")
	}
}
