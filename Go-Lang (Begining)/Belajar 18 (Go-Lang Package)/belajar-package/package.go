// ================================================== Package & Import ==================================================
// /*
// package = 	tempat yang digunakan untuk mengorganisir kode program yang dibuat
// 			bisa merapihkan kode yang dibuat / grouping
// 			sebenarnya hanya direktori folder disistem oprasi

// membuat package = 	1. buat folder dengan nama apapun, contoh : helper
// 					2. kemudian buat file dengan nama apapun, contoh : SayHallo.go
// 					3. kemudai liat isi file SayHallo.go untuk lebih lengkap
// */
// package main

// //import = untuk mengakses file lain yang berbeda package atau yang berda di luat package
// // import "golang-package\helper"
// import (
// 	"belajar-package/givehallo"
// )

// func main() {
// 	givehallo.SayHello("Wandi")
// 	// fmt.Println(result)
// }

// ================================================== Access Modifier ==================================================
// /*
// Hak akses variable atau function
// jika nama function atau variable diawali dengan huruf besar maka bisa diakses dari package lain jika dimulai dengan huruf kecil maka tidak bisa diakses oleh package lain
// */
// package main

// func main() {

// }

// ================================================== Package Initialization ==================================================
// /*
// biasanya digunakan untuk function mengkoneksikan ke database (kalau dalam php sama sperti construct__ yang akan dijalankan terlebih dulu sebelum function lain)
// */
// //untuk packagenya menggunakan package database cari saja
// package main

// //blank identifier
// // import _ "belajar-package/database" //tidak digunakan tapi tetap bisa di panggil
// import (
// 	"belajar-package/database"
// 	"fmt"
// )

// func main() {
// 	result := database.GetDatabase()

// 	fmt.Println(result)
// }

// ================================================== Package OS ==================================================
// // Package bawwan golang
// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	arg := os.Args //menagmbil data argumen

// 	fmt.Println("Argument :")
// 	fmt.Println(arg)
// 	// fmt.Println(arg[1])
// 	// fmt.Println(arg[2])
// 	// fmt.Println(arg[3])
// 	//untuk menjalannya perintah "go run package.go Rayandra Wandi Marselana"

// 	hostnam, err := os.Hostname() //untuk mengambil hostname / nama host sistem oprasi yang di pakai
// 	if err == nil {
// 		fmt.Println("Hostname ", hostnam)
// 	} else {
// 		fmt.Println("Error", err.Error())
// 	}

// 	//mensetting username dan password aplikasi
// 	username := os.Getenv("APP_USERNAME")
// 	password := os.Getenv("APP_PASSWORD")

// 	//cara seeting isi username dan password
// 	// export <nama_variabel dari os.Getenv>=<nama username>
// 	// export <nama_variabel dari os.Getenv>=<nama password>

// 	fmt.Println(username)
// 	fmt.Println(password)
// }

// //masih banyak fitur package os bisa di baca di websiter reminya

// ================================================== Package Flag ==================================================
// // berisikan fungsionalitas untuk memparshing command line argument
// package main

// import (
// 	"flag"
// 	"fmt"
// )

// func main() {
// 	host := flag.String("host", "localhost", "Put your database host")
// 	username := flag.String("username", "root", "Put your database username")
// 	password := flag.String("password", "root", "Put your database password")

// 	// flag.Parse()
// 	fmt.Println(*host, *username, *password) //jika ingin hasil pointer hapus * nya
// }

// ================================================== Package String ==================================================
// //package untuk memanipulasi string
// /*
// Function manipulasi string
// 	Function : strings.Trim(string, cutset)
// 	Kegunaan : memotong cutset di awal dan di akhir string

// 	Function : strings.ToLower(string)
// 	Kegunaan : Membuat semua string menjadi lower case atau huruf kecil

// 	Function : strings.ToUpper(string)
// 	Kegunaan : Membuat semua string menjadi upper case atau huruf besar

// 	Function : strings.Split(string, separator)
// 	Kegunaan : Memotong string berdasarkan separator

// 	Function : strings.Contains(string, search)
// 	Kegunaan : Mengecek apakah string mengandung string lain

// 	Function : strings.ReplaceAll(string, from, to)
// 	Kegunaan : Mengubah semua string dari form ke to
// */
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	kata := "     Wandi Marselana Marselana     "

// 	fmt.Println("Kata Awal : ", kata)
// 	fmt.Println("Contains : ", strings.Contains(kata, "Marselana")) //mengecek apakah ada kata marselana atau tidak
// 	fmt.Println("Split : ", strings.Split(kata, " "))
// 	fmt.Println("ToLower : ", strings.ToLower(kata))
// 	fmt.Println("ToUpper : ", strings.ToUpper(kata))
// 	fmt.Println("Trim : ", strings.Trim(kata, " "))
// 	fmt.Println("ReplaceAll : ", strings.ReplaceAll(kata, "Marselana", ""))

// }

// ================================================== Package Strconv ==================================================
// //string konversi
// /*
// Function
// 	Function : strconv.ParseBool(string)
// 	Kegunaan : mengubah string ke boolean

// 	Function : strconv.ParseFloat(string)
// 	Kegunaan : mengubah string ke Floating

// 	Function : strconv.ParseInt(string)
// 	Kegunaan : mengubah string ke int64

// 	Function : strconv.FormatBool(bool)
// 	Kegunaan : mengubah boolean ke string

// 	Function : strconv.FormatFloat(float, ...)
// 	Kegunaan : mengubah float64 ke string

// 	Function : strconv.FormatInt(int, ...)
// 	Kegunaan : mengubah Int64 ke string
// */
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	//boolean
// 	boolean, err := strconv.ParseBool("true")
// 	if err == nil {
// 		fmt.Println(boolean)
// 	} else {
// 		fmt.Println("Error", err.Error())
// 	}

// 	//number
// 	number, err := strconv.ParseInt("1000000", 10, 64) //10 adalah kode desimal, 2 adalah biner, dll
// 	if err == nil {
// 		fmt.Println(number)
// 	} else {
// 		fmt.Println("Error", err.Error())
// 	}

// 	value := strconv.FormatInt(1000000, 2) //2 adalah kode binary
// 	fmt.Println(value)
// }

// //masih banyak fitur package os bisa di baca di websiter reminya

// ================================================== Package Math ==================================================
// // berisikan fungsi matematika
// /*
// Function
// 	Function : math.Round(float64)
// 	Kegunaan : Membulatkan float64, seduai dengan yang paling dekat

// 	Function : math.Floor(float64)
// 	Kegunaan : Membulatkan float64 kebawah

// 	Function : math.Ceil(float64)
// 	Kegunaan : Membulatkan float64 keatas

// 	Function : math.Max(float64, float64)
// 	Kegunaan : mengembalikan nilai float64 paling besar

// 	Function : math.Min(float64, float64)
// 	Kegunaan : mengembalikan nilai float64 paling kecil
// */
// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	fmt.Println(math.Ceil(1.40))
// 	fmt.Println(math.Floor(1.60))
// 	fmt.Println(math.Round(1.60))
// 	fmt.Println(math.Max(1, 2))
// 	fmt.Println(math.Min(1, 2))

// }

// ================================================== Package Container/List ==================================================
// // implementasi struktur data double link list
// // yang mempunyai head and tail
// package main

// import (
// 	"container/list"
// 	"fmt"
// )

// func main() {
// 	data := list.New()         //membuat wadah data list baru
// 	data.PushBack("Rayandra")  //menambahkan data kebelakang
// 	data.PushBack("Wandi")     //menambahkan data kebelakang
// 	data.PushBack("Marselana") //menambahkan data kebelakang
// 	data.PushFront("Hallo : ") //menambahkan data kepaling depan

// 	for i := data.Front(); i != nil; i = i.Next() { //pengulangan pemasukan data. //data.Front()m = mengambil data paling depan. //lalu di cek datanya tidak kosong. //i.Next() = seperti ++
// 		fmt.Println(i.Value) //value = isinya
// 	}

// 	fmt.Println(data.Front())               //mengambil data dari depan
// 	fmt.Println(data.Back())                //mengambil data dari belakang
// 	fmt.Println(data.Front().Next().Next()) //mengambil data Selanjutnya
// 	fmt.Println(data.Back().Prev())         //mengambil data sebelumnya
// }

// ================================================== Package Container/Ring ==================================================
// // implementasi struktur data circular ring
// // dimana di akhir akan kembali ke awal
// // kurang lebih sama dengan list
// package main

// import (
// 	"container/ring"
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	data := ring.New(5) //membuat wadah data ring

// 	for i := 1; i < data.Len(); i++ {
// 		data.Value = "Value " + strconv.FormatInt(int64(i), 10)
// 		data = data.Next()
// 	}

// 	data.Do(func(value interface{}) {//untuk menampilkan
// 		fmt.Println(value)
// 	})
// }

// ================================================== Package Sort ==================================================
// // pemgurutan data
// package main

// import (
// 	"fmt"
// 	"sort"
// )

// type User struct {
// 	nama string
// 	umur int
// }

// type UserSlice []User //pengaliasan

// func (value UserSlice) Len() int { //mengitung panjang
// 	return len(value)
// }

// func (value UserSlice) Less(i, j int) bool {
// 	return value[i].umur <= value[j].umur
// }

// func (value UserSlice) Swap(i, j int) {
// 	value[i], value[j] = value[j], value[i]
// }

// func main() {
// 	data := []User{
// 		{"Rayandra", 12},
// 		{"Wandi", 10},
// 		{"Marselana", 21},
// 	}

// 	sort.Sort(UserSlice(data))
// 	fmt.Println(data)
// }

// ================================================== Package Time ==================================================
// // fungsionalitas unutk management waktu
// /*
// Format
// 	Function : time.Now()
// 	Keterangan : Mendapatkan waktu saat ini

// 	Function : time.Date(...)
// 	Keterangan : Membuat waktu

// 	Function : time.Parse(layout, string)
// 	Keterangan : Memparshing waktu dari string
// */
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	skrng := time.Now() //menampilkan waktu sekarang
// 	fmt.Println(skrng)
// 	fmt.Println(time.Now().Year())       //menampilkan tahun
// 	fmt.Println(time.Now().Month())      //menampilkan bulan
// 	fmt.Println(time.Now().Day())        //menampilkan hari
// 	fmt.Println(time.Now().Hour())       //menampilkan jam
// 	fmt.Println(time.Now().Minute())     //menampilkan menit
// 	fmt.Println(time.Now().Second())     //menampilkan detik
// 	fmt.Println(time.Now().Nanosecond()) //menampilkan nano detik

// 	utc := time.Date(2022, time.January, 10, 23, 0, 0, 0, time.UTC)
// 	fmt.Println(utc)

// 	parse, _ := time.Parse(time.RFC3339, "2010-02-02T15:04:05Z")
// 	fmt.Println(parse)
// }

// ================================================== Package Reflect ==================================================
// // untuk melihat structur code saat berjalan semacam eco exit 1 di dalam php
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type Sample struct {
// 	Name string
// }

// func main() {
// 	sample := Sample{"Udin"}
// 	sampleType := reflect.TypeOf(sample)
// 	structField := sampleType.Field(0)

// 	fmt.Println(structField.Name)
// }

// ================================================== Package Regexp ==================================================
//package ini menggunakna library C yang dibuat google bernama RE2
/*
Format
	function : regexp.MustCompile(string)
	keterangan : membuat regexp

	function : regexp.MustString(string) bool
	keterangan : mengecek apakah regexp match dengan string

	function : regexp.FindAllString(string, max)
	keterangan : mencari string yang match dengan max jumlah hasil
*/
package main

import (
	"fmt"
	"regexp"
)

func main() {
	var rege = regexp.MustCompile("e([a-z])o")

	fmt.Println(rege.MatchString("dandi"))
	fmt.Println(rege.MatchString("wandi"))
	fmt.Println(rege.MatchString("Wandi"))

	fmt.Println(rege.FindAllString("asdasd asdasd asdqw qwdqd", 1))

	// fmt.Println(rege)
}
