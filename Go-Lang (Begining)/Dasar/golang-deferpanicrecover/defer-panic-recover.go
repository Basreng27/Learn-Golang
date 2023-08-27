package main

//================================================== Defer ==================================================
// /*
// Defer = function yang bisa kita jadwlkan untuk di eksekusi setelah sebuah function selesai di ekseskusi
// 		akan selalu di eksekusi walaupun ada error di functionnya
// */
// //tidak error
// func logging() {
// 	fmt.Println("Selesai memanggil function") //menajalankan isi dari function 4.
// }

// func RunApp() {
// 	defer logging()                    //setelah semua funcction dijalankan maka lanjut ke function berikutnya 3.
// 	fmt.Println("Aplikasi dijalankan") //menjalankan isi dari function 2.
// }

// func main() {
// 	RunApp() //menjalankan aplikasi pertama 1.
// }

// //error
// func logging() {
// 	fmt.Println("Selesai memanggil function") //menajalankan isi dari function 4.
// }

// func RunApp(value int) {
// 	// defer logging()                    //setelah semua funcction dijalankan maka lanjut ke function berikutnya 3.
// 	fmt.Println("Aplikasi dijalankan") //menjalankan isi dari function 2.

// 	result := 10 / value
// 	fmt.Println(result)
// 	logging() //harus mengguanakan defer

// }

// func main() {
// 	RunApp(0) //menjalankan aplikasi pertama 1.
// }

//================================================== Panic ==================================================
// /*
// Panic = function yang bisa digunakan untuk menghentikan program
// 		biasanya digunakan saat terjadi error
// 		saat di panggil program akan berhenti tapi defer tetap jalan
// */
// func EndApp() {
// 	fmt.Println("End App")
// }

// func RunApp(eror bool) {
// 	defer EndApp()
// 	if eror {
// 		panic("ERROR")
// 	}

// 	fmt.Println("Aplikasi Berjalan")
// }

// func main() {
// 	RunApp(false)
// }

//================================================== Recover ==================================================
// /*
// Recover = function yang bisa digunakan untuk menangkan data panic
// 		dengan recover proses panic akan terhenti dan program tetap berjalan
// */

// func EndApp() {
// 	massage := recover() //recover //mengambil massage dari panic
// 	if massage != nil {
// 		fmt.Println("Terjadi Error", massage)
// 	}
// 	fmt.Println("End App")
// }

// func RunApp(eror bool) {
// 	defer EndApp()

// 	if eror {
// 		panic("ERROR")
// 	}

// 	fmt.Println("Aplikasi Jalan")
// }

// func main() {
// 	RunApp(true)
// }
