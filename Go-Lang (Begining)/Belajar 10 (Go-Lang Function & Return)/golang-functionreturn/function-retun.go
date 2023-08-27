package main

import "fmt"

//==================================================== Function ====================================================
// /*
// func = untuk mendeklarasikan/membuat function dan di ikuti nana functionnya
// <nama_function>() = untuk memanggil function
// */
// func Hallo() {
// 	fmt.Println("Hallo Gaizz")
// }

// func main() { //function inti
// 	//menampilkan function Hallo 10x
// 	for i := 1; i <= 10; i++ { //perulangan
// 		Hallo() //memanggil funtion Hallo
// 	}
// }

//==================================================== Function Parameter ====================================================
// func Param(awal string, tengah string, akhir string) { //menerima parameter
// 	fmt.Println("Hallo", awal, tengah, akhir)
// }

// func main() {
// 	awl := "Rayandra"
// 	tgh := "Wandi"
// 	akh := "Marselana"
// 	Param(awl, tgh, akh)
// 	// Param("Rayandra", "Wandi", "Marselana") //mengirim parameter ke funtion param
// }

//==================================================== Function Return Value ====================================================
// func ReturnValue(nama string) string { //string diluar kurung itu adalah nilai/tipe data retun/nilai kembali
// 	// return "Hallo" + nama //untuk return dan return selalu di simpan di paling bawah

// 	if nama == "" {
// 		return "Hallo Gaiz"
// 	} else {
// 		return "Hallo" + nama
// 	}
// }

// func main() {
// 	retur := ReturnValue("")
// 	fmt.Println(retur)
// 	// fmt.Println(ReturnValue("Rayandra")) //memanggil dan mengirim data ke function dan menerima kembalian
// }

//==================================================== Function Returning Multiple Values ====================================================
// /*
// Mengembalikan data lebih dari satu / return lebih dari satu
// */
// func ReturnMultipleValues() (string, string, string) { //kurung buka kedua untuk return multiple digunakan kurung buka tambahan dan bisa lebih dari 2 return
// 	return "Rayandra", "Wandi", "Marselana" //mereturn 2 return
// }

// func main() {
// 	// _ //untuk menghiraukan variable

// 	awal, mid, akhir := ReturnMultipleValues() //variable awal akan mengambil data dari retun pertama, dan akhir akan mengambil data dari return ke dua

// 	fmt.Println(awal, mid, akhir) //memanggil variable yang sudah diisi return pada function
// }

//==================================================== Named Return Values ====================================================
// /*
// Retun yang di buat dengan variable
// */
// func NamedReturnValues() (first, mid, last string) {
// 	first = "Rayandra"
// 	mid = "Wandi"
// 	last = "Marselana"

// 	return
// 	// return first, mid, last // bisa menggunakan return saja tanpa di panggil ulang
// }

// func main() {
// 	first, mid, last := NamedReturnValues()

// 	fmt.Println(first, mid, last)
// }

//==================================================== Variadic Function ====================================================
// /*
// bisa di anggap parrameternya seperti array atau hampir sama
// */
// func VariadicFunction(numbers ...int) int { //dianggap variadic function karena parameternya bersifat variable argument cirinya ada titik 3x. yang artinya bisa memasukan number atau isi lebih dari 1x
// 	total := 0
// 	for _, number := range numbers {
// 		total += number
// 	}

// 	return total
// }

// func main() {
// 	nilai := VariadicFunction(10, 10, 10, 10, 10, 10)
// 	fmt.Println(nilai)

// 	numbers := []int{10, 20, 30} //bisa mengambil data dari slice jika sudah memounyai slice
// 	totall := VariadicFunction(numbers...)
// 	fmt.Println(totall)
// }

//==================================================== Function As Value ====================================================
// func FunctionValue(name string) string {
// 	return "Hallo " + name
// }

// func main() {
// 	nama := FunctionValue //membuat function menjadi value

// 	fmt.Println(nama("Rayandra Wandi Marselana")) //memanggil nama variable yang isinya function
// }

//==================================================== Function As Parameter ====================================================
// func HalloFilter(nama string, filter func(string) string) {
// 	namafilter := filter(nama)
// 	fmt.Println("Hello", filter(namafilter))
// }

// func SpamFilter(nama string) string {
// 	if nama == "Babi" {
// 		return "Kasar"
// 	} else {
// 		return nama
// 	}
// }

// func main() {
// 	HalloFilter("Wandi", SpamFilter)

// 	filter := SpamFilter
// 	HalloFilter("Babi", filter)
// }

// //Function Type Declaration bisa membuat alias sebagai function
// type Filter func(string) string //pengaliasan function dengan nama Filter

// func HF(name string, filter Filter) {
// 	nf := filter(name)
// 	fmt.Println("Hello", filter(nf))
// }

//==================================================== Anonymous Function ====================================================
// /*
// Function yang tidak memiliki nama function
// */
// type Blacklist func(string) bool //membuat alias Blacklist

// func RegisUser(nama string, blacklish Blacklist) {
// 	if blacklish(nama) { //mengecek jika return value true maka masuk ke sini. tru di ambil dari alias Blaclist
// 		fmt.Println(nama, "Kamu diBlok")
// 	} else {
// 		fmt.Println("Welcome", nama)
// 	}
// }

// // func blacklishAdmin(nama string) bool {
// // 	return nama == "admin"
// // }

// // func blacklishRoot(nama string) bool {
// // 	return nama == "root"
// // }

// func main() {
// 	blacklist := func(nama string) bool { //pendeklarasian function lebih singkat dari sebelumnya atua function tanpa nama
// 		return nama == "Babi"
// 	}

// 	RegisUser("Wandi", blacklist)
// 	RegisUser("Babi", func(nama string) bool {
// 		return nama == "Babi"
// 	})
// }

//==================================================== Recursive Function ====================================================
/*
function yang memanggil dirinya sendiri
*/
func Factorial(value int) int {
	// result := 1

	// for i := value; i > 0; i-- {
	// 	result *= i
	// }
	// return result

	if value == 1 {
		return 1
	} else {
		return value * Factorial(value-1)
	}
}

func main() {
	fmt.Println(Factorial(15))
}
