package main

import "fmt"

//================================================== For ==================================================
// func main() {
// 	counter := 1 //inisialisasi variable

// 	for counter <= 10 { //inisialisai for
// 		fmt.Println("Perulangan Ke-", counter)
// 		counter++ //penambahan 1 setiap di ulang
// 	}

// 	//for dengan statement
// 	for count := 1; count <= 10; count++ {
// 		fmt.Println("Perulangan statement ke-", count)
// 	}

// 	//for range untuk array
// 	slice := []string{
// 		"Rayandra",
// 		"Wandi",
// 		"Marselana",
// 	}

// 	for i := 0; i < len(slice); i++ {
// 		fmt.Println(slice[i])
// 	}

// 	//mengambil semua array menggunakan for range
// 	for i, name := range slice {
// 		fmt.Println("index", i, "=", name)
// 		// fmt.Println(name) //tinggal di ganti i menjadi _. "_" = adalahariable yang tidak di butuhkan atau tidak usah di print
// 	}

// 	//array map
// 	person := make(map[string]string)
// 	person["nama"] = "Rayandra Wandi Marselana"
// 	person["title"] = "Begginer Traner"

// 	for key, isi := range person {
// 		fmt.Println(key, "=", isi)
// 	}
// }

//================================================== Break & Continue ==================================================
func main() {
	/*
		jika ada break maka pengulanagan akan berhenti
		jika ada continueelanjutkan ke program selanjutnya
	*/

	//Break
	// for i := 0; 1 < 10; i+ {
	// 	if i ==  {
	// 		break //ketika i == 5 maka program akan di hentikan sampai sni
	//	}
	// 	fmt.Println("Pengulanagna Ke-",i)
	// }

	//continue
	for a := 0; 1 < 10; a++ {
		if a%2 == 0 {
			break //ketika a == 5 maka program akan di hentikan dan di lanjutkan lagi melewatiang selanjutnya
		}
		fmt.Println("Pengulanagna Ke-", a)
	}
}
