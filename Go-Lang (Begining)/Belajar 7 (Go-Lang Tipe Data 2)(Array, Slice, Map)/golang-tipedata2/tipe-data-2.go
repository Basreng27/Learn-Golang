package main

import "fmt"

//======================================================= Tipe Data Array =======================================================
// func main() {
// 	/*
// 		array[Rayandra, Wandi, Marselana]
// 			index : 0
// 			data : Rayandra

// 			index : 1
// 			data : Wandi

// 			index : 2
// 			data : Marselana
// 	*/

// 	//Membuat array secara manual
// 	var name [3]string
// 	name[0] = "Rayandra" //inisialisai data
// 	name[1] = "Wandi"
// 	name[2] = "Marselana"

// 	fmt.Println(name[0], " = Index ke-0") //mengambil data array
// 	fmt.Println(name[1], " = Index ke-1")
// 	fmt.Println(name[2], " = Index ke-2")

// 	//Membuat array secara langsung
// 	var value = [3]int{
// 		100,
// 		99,
// 		98,
// 	}

// 	fmt.Println(value) //mengambil data array

// 	/*
// 		Function array
// 			oprasi : len(array)
// 			keterangan : Untuk mendapatkan panjang array

// 			oprasi : array[index]
// 			keterangan : Mendapatkan data di posisi index

// 			oprasi : array[index] = value
// 			keterangan : Mengubah data di posisi index
// 	*/

// 	fmt.Println(len(name))
// 	fmt.Println(len(value))
// }

//======================================================= Tipe Data Slice =======================================================
// func main() {
// 	/*
// 		Tipe data slice = potongan array
// 		memiliki 3 data yaitu pointer, length dan capacity
// 			pointer = penunjuk data pertama
// 			length = panjang dari slice
// 			capacity = kapasitas slice dengan catatan length tidak lebih dari capacity

// 		membuat slice dari array
// 			membuat slice : array[low:high]
// 			keterangan : Membuat slice dari array dimulai dari index low sampai index sebelum high (low = terkecil, hight = terbesar)

// 			membuat slice : array[low:]
// 			keterangan : Membuat slice dari array dimulai dari index low sampai index akhir

// 			membuat slice : array[:high]
// 			keterangan : Membuat slice dari array dimulai dari index 0 sampai index sebelum high

// 			membuat slice : array[:]
// 			keterangan : Membuat slice dari array dimulai dari index 0 sampai index akhir di array

// 		function
// 			oprasi : len(slice)
// 			keterangan : untuk pendapat panjang

// 			oprasi : cap(slice)
// 			keterangan : untuk pendapat kapasitas

// 			oprasi : append(slice, data)
// 			keterangan : membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas sudah penuh, maka akan membuat array data baru

// 			oprasi : make([]TypeData, length, capacity)
// 			keterangan : Membuat slice baru

// 			oprasi : copy(destination, source)
// 			keterangan : menyalin slice dari source ke destination
// 	*/

// 	//len, cap
// 	var bulan = [...]string{ //[...] = untuk jumlah array tidak diketahui
// 		"Januari",
// 		"Februari",
// 		"Maret",
// 		"April",
// 		"Mei",
// 		"Juni",
// 		"Juli",
// 		"Agustus",
// 		"September",
// 		"Oktober",
// 		"November",
// 		"Desember",
// 	}

// 	slice1 := bulan[4:7]

// 	fmt.Println(slice1)      //mengambil data slice1
// 	fmt.Println(len(slice1)) //mengambil panjang data slice1
// 	fmt.Println(cap(slice1)) //mengambil kapasitas data slice1

// 	slice1[0] = "Mei Lagi" //[0] mengambil data dari index slice1 bukan bulan
// 	fmt.Println(bulan)

// 	//append
// 	hari := [...]string{
// 		"Senin",
// 		"Selasa",
// 		"Rabu",
// 		"Kamis",
// 		"Jumat",
// 		"Sabtu",
// 		"Minggu",
// 	}

// 	slicehari := hari[5:]
// 	slicehari[0] = "Sabtu Baru"
// 	slicehari[1] = "Minggu Baru"
// 	fmt.Println(hari) //senin-jumat dan sabtu-minggu baru

// 	slicehari1 := append(slicehari, "Libur Gaiz")
// 	slicehari1[0] = "Ups"
// 	fmt.Println(slicehari1) //[ups, minggu, libur gaiz]
// 	fmt.Println(hari)

// 	//make
// 	newslice := make([]string, 2, 5) //membuat slice baru
// 	newslice[0] = "Rayandra"         //mengisikan data slice
// 	newslice[1] = "Wandi"

// 	fmt.Println(newslice)
// 	fmt.Println(len(newslice))
// 	fmt.Println(cap(newslice))

// 	//copy
// 	dari_slice := hari[:]
// 	ke_slice := make([]string, len(dari_slice), cap(dari_slice))

// 	copy(ke_slice, dari_slice)

// 	fmt.Println(ke_slice)

// 	//perbedaan pembuatan array dan slice
// 	iniarray := [...]int{1, 2, 3, 4, 5} //Array
// 	inislice := []int{1, 2, 3, 4}       //Slice

// 	fmt.Println(iniarray)
// 	fmt.Println(inislice)
// }

//======================================================= Tipe Data Map =======================================================
func main() {
	/*
		Map = sama seperti primarykey di mysql
		map[tipdedata]tipedata
		"key" : "value"
	*/

	person := map[string]string{ //membuat map
		"nama": "Rayandra",
		"jk":   "laki",
	}

	person["title"] = "Beginer programer" //membuat atau menambahkan isi

	fmt.Println(person)
	fmt.Println(person["nama"]) //pemanggilan
	fmt.Println(person["jk"])

	/*
		Function Map
			operasi : len(map)
			keterangan : mendapatkan jumlah data di map

			operasi : map[key]
			keterangan : Mengambil data di map menggunakan key tertentu

			operasi : map[key] = value
			keterangan : Mengubah data di map dengan key

			operasi : make(map[TypeKey]TypeValue)
			keterangan : Membuat map baru

			operasi : delete(map, key)
			keterangan : menghapus data di map dengan key
	*/

	buku := make(map[string]string)
	buku["no"] = "satu"
	buku["judul"] = "buku baru"
	buku["dapus"] = "dafatr pustaka"

	delete(buku, "dapus")

	fmt.Println(buku)

}
