package main

import "fmt"

//================================================== Variable ==================================================
// func main() {
// 	//variable harus dipakai kalau tidak dipakai akan eror
// 	var nama string //pendeklarasian variable

// 	nama = "Rayandra Wandi" //penampungan variable
// 	fmt.Println(nama)       //pemanggilan variable

// 	nama = "Rayandra Wandi Marselana" //penimpaan variable sebelumnya
// 	fmt.Println(nama)                 //pemanggilan variable

// 	var hewan = "Lele" //bisa langsung isikan variable seperti ini
// 	fmt.Println(hewan)

// 	umur := 22 //pendeklarasian variable tanpa menggunakan kata var hanya menggunakan :=
// 	fmt.Println(umur)

// 	//Deklarasi Multiple variable
// 	var (
// 		NamaAwal   = "Rayandra"
// 		NamaTengah = "Wandi"
// 		NamaAkhir  = "Marselana"
// 	)

// 	fmt.Println(NamaAwal, NamaTengah, NamaAkhir) //pemanggilan secara lebih dari 1 variable
// }

// ================================================== Constant ==================================================
// func main() {
// 	//constanta adalah variable yang tidak bisa diubah saat pertama kali di deklarasikan
// 	const umur int32 = 22   //deklarasi variable constant
// 	const nama = "Rayandra" //deklarasi variable constant

// 	fmt.Println("Nama : ", nama, "Umur : ", umur)

// 	//Deklarasi Multiple variable
// 	const (
// 		NamaAwal   string = "Rayandra"
// 		NamaTengah        = "Wandi"
// 		NamaAkhir         = "Marselana"
// 	)
// }

// ================================================== Konversi Tipe Data ==================================================
// func main() {
// 	var nilai32 int32 = 100000
// 	var nilai64 int64 = int64(nilai32) //mengonversikan nilai32 int32 menjadi int64
// 	var nilai16 int16 = int16(nilai32) //mengonversikan nilai32 int32 menjadi int16

// 	fmt.Println("Nilai int32 : ", nilai32)
// 	fmt.Println("Nilai int64 : ", nilai64)
// 	fmt.Println("Nilai int16 : ", nilai16)

// 	//komversi dari angka ke huruf
// 	var name = "Wandi"
// 	var e byte = name[0]
// 	var estring string = string(e)

// 	fmt.Println(name)
// 	fmt.Println(estring)
// }

// ================================================== Tipe Data Declaration ==================================================
//bisa disebut peng aliasan variable
func main() {
	type NoKK string //didekalarasikan NoKK adalah string atau string alias NoKK

	var KK NoKK = "453412313"
	fmt.Println(KK)
	fmt.Println(NoKK("1234321"))
}
