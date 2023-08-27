package main

import "fmt"

//================================================== Pointer ==================================================
// type Address struct {
// 	kota, provinsi, negara string
// }

// func main() {
// 	//pass by value
// 	address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
// 	address2 := address1 //address1 mengcopy dirinya dan menyimpan di address2

// 	address2.kota = "Cilengkrang" //merubah address2

// 	fmt.Println(address1) //address 1 tidak berubah
// 	fmt.Println(address2) //address 2 berubah kotanya

// 	//pointer dapat membuat pass by reference / pass by reference dengan pointer
// 	//oprator pointer = &<nama_variabke>
// 	alamat1 := Address{"Cipatik", "Jawa Barat", "Indonesia"}
// 	alamat2 := &alamat1 //mengubah data alamat 1 menggunakan pointer

// 	alamat2.kota = "Cililin" //merubah data addres 1

// 	fmt.Println(alamat1) //data berubah
// 	fmt.Println(alamat2) //data berubah

// 	//oprator bintang = *<nama_variable>
// 	add1 := Address{"Tegal", "Jawa Tengah", "Indonesia"}
// 	add2 := &add1

// 	add2.kota = "Banyumas"

// 	*add2 = Address{"Surabaya", "Jawa Timur", "Indonesia"}

// 	fmt.Println(add1) //ikut berubah dengan paksa
// 	fmt.Println(add2)

// 	//pointer dengan isi data kosong = new
// 	var add3 *Address = new(Address) //langsung membuat pointer ke addres baru
// 	fmt.Println(add3)                //data kosong

// 	add3.negara = "Indonesia" //mengisi field negara
// 	fmt.Println(add3)         //data ada
// }

//================================================== Pointer Di Function ==================================================
// type Alamat struct {
// 	kota, provinsi, negara string
// }

// func changeAddres(alamat *Alamat) { //menggunakan pointer
// 	// func changeAddres(alamat Alamat) {//tidak menggunakan pointer
// 	alamat.negara = "Indonesia"
// }

// func main() {
// 	//tidak mengugnakan pointer
// 	alt := Alamat{"Bandung", "Jawa Barat", "Nigeria"}
// 	changeAddres(&alt) //menggunakan pointer
// 	// changeAddres(alt)//tidak menggunakan pointer

// 	fmt.Println(alt)
// }

//================================================== Pointer Di Method ==================================================
/*
disarankan menggunakan pointer jika mengguanakan method sehingga tidak boros memory
*/

type Org struct {
	nama string
}

// func (orng Org) Status() {//tidak menggunakan pointer
func (orng *Org) Status() { //mengguanakan pointer
	orng.nama = "Mr." + orng.nama
	fmt.Println("Method ", orng.nama)
}

func main() {
	wandi := Org{"Wandi"}
	wandi.Status()
	fmt.Println(wandi.nama)
}
