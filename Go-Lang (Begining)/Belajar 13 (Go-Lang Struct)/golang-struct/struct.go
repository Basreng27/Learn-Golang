package main

import "fmt"

//================================================== Struct ==================================================
// /*
// Struct kumpulan dari field
// */

// type Orang struct { //membuat struct
// 	nama, alamat string
// 	umur         int
// }

// //membuat data struct
// func main() {
// 	//struct literals /penulisan isi struct
// 	var wandi Orang                         //membuat variabel untuk struct
// 	wandi.nama = "Rayandra Wandi Marselana" //untuk mengisi struct Orang pada field nama
// 	wandi.alamat = "Jln Cilengkrang II"
// 	wandi.umur = 21

// 	fmt.Println(wandi) //memanggil variable wandi

// 	Raya := Orang{
// 		nama:   "Raya",
// 		alamat: "Jln Cilengkrang II",
// 		umur:   15,
// 	}

// 	fmt.Println(Raya)

// 	Yayu := Orang{"Yayu", "Cilengkrang", 40}
// 	fmt.Println(Yayu)
// }

//================================================== Struct Method ==================================================
/*
method adalah function
struct yang mempunyai function
*/

type Orang struct { //membuat stryct
	nama, alamat string
	umur         int
}

func (costumer Orang) dataDiri(nama string) { //membuat struct yang mempunyai function //menggunakan parameter
	// func (costumer Orang) dataDiri() { //membuat struct yang mempunyai function //tidak menggunakan parameter
	fmt.Println("Hallo ", nama, " Data Diri Anda : ", costumer.nama, costumer.alamat, costumer.umur)
}

func main() {
	wandi := Orang{
		nama:   "Rayandra Wandi Marselana",
		alamat: "Jln Cilengkrang 2",
		umur:   21,
	}

	wandi.dataDiri("Gaiz") //untuk memanggil struct function
}
