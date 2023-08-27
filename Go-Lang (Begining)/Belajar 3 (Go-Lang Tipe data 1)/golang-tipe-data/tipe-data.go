package main

import "fmt"

//============================================= Tipe Data Number =============================================
// func main() {
// 	/*
// 		Tipe data number

// 		2 jenis tipe data number :
// 			integer
// 				integer mempunyai nilai negative
// 				1.	tipe-data = int8
// 					Nilai Minimum = -128
// 					Nilai Maksium = 127
// 				2.	tipe-data = int16
// 					Nilai Minimum = -32768
// 					Nilai Maksium = 32767
// 				3.	tipe-data = int32
// 					Nilai Minimum = -2147483648
// 					Nilai Maksium = 2147483647
// 				4.	tipe-data = int64
// 					Nilai Minimum = -9223372036854775708
// 					Nilai Maksium = 9223372036854775707

// 				integer tidak mempunyai nilai negative
// 				1.	tipe-data = uint8
// 					Nilai Minimum = 0
// 					Nilai Maksium = 255
// 				2.	tipe-data = uint16
// 					Nilai Minimum = 0
// 					Nilai Maksium = 65535
// 				3.	tipe-data = uint32
// 					Nilai Minimum = 0
// 					Nilai Maksium = 4294967295
// 				4.	tipe-data = uint64
// 					Nilai Minimum = 0
// 					Nilai Maksium = 18446744073709551615

// 			floating point
// 				1.	tipe-data = float32
// 					Nilai Minimum = 1.18 x 10(-38) //yang dalam kurung itu pangkat
// 					Nilai Maksium = 3.4 x 10(38) //yang dalam kurung itu pangkat
// 				2.	tipe-data = float64
// 					Nilai Minimum = 2.23 x 10(-308) //yang dalam kurung itu pangkat
// 					Nilai Maksium = 1.80 x 10(308) //yang dalam kurung itu pangkat
// 				3.	tipe-data = complex64
// 					untuk program mate,atik complex
// 				4.	tipe-data = complex128
// 					untuk program mate,atik complex

// 		Alias atau pengalisan
// 			byte = int8
// 			rune = int32
// 			int = minimal int32
// 			uint = minimal int32
// 	*/

// 	fmt.Println("Satu = ", 1)             //koma untuk penghubung parameter
// 	fmt.Println("Dua = ", 2)              //koma untuk penghubung parameter
// 	fmt.Println("Satu Koma Satu = ", 1.1) //koma untuk penghubung parameter
// }

//============================================= Tipe Data Boolean =============================================
// func main() {
// 	/*
// 		Boolean
// 			Nilai : True
// 			Keterangan : Benar

// 			Nilai : False
// 			Keterangan : Salah
// 	*/

// 	fmt.Println("Benar : ", true)
// 	fmt.Println("Salah : ", false)
// }

//============================================= Tipe Data String =============================================
func main() {
	/*
		Function untuk string
			Function : len("string")
			Keterangan : Menghitung jumlah karakter string

			Function : "string"[number]
			Keterangan : Mengambil karakter pada posisi yang ditentukan
	*/

	fmt.Println("Rayandra")
	fmt.Println("Rayandra Wandi")
	fmt.Println("Rayandra Wandi Marselana")

	fmt.Println(len("Rayandra"))               //menghitung total karakter
	fmt.Println("Rayandra Wandi"[3])           //menampilkan karakter "a"
	fmt.Println("Rayandra Wandi Marselana"[9]) //menampilkan karakter "W"
}
