/*
Pengenalan testing package : packagenya bernama testing

	perintah :	testing.T = struct ini digunakan untuk unit test di golang
				testing.M = untuk mengatur life cycle testing / menentukan alur testing
				testing.B = untuk melakukan benchmarking(mengukur kecepatan program)
*/
package main

import (
	"fmt"
	"golang-unit-test/helper"
)

/*
Unit Testing
	Aturan File test /testing unit
		untuk membuat file unit test maka harus menggunakan "_test" diakhiran
		contoh : jika membuat file "hello.go" maka untuk membuat testnya menggunakan nama file "hello_test.go" (contoh pada package "helper")
			"hello_test.go" digunakan hanya untuk mengetest file "hello.go" saja

	Aturan Fucntion unit test
		untuk membuat function unit test maka harus diawali dengan nama "Test"
		contoh : jika membuat function "SayHallo" untuk mengetest functionnya harus di buat unit test dengan nama "TestSayHallo"
			untuk lanjutannya bebas tapi harus diawali deng "Test"
				lalu harus memiliki parameter (t *testing.T) dan tidak mengembalikan value (contoh bisa di liat di package "helper" file "hello_test.go")

	Aturan untuk menjalankan Unit test
		sebelum melakukan perintah di bawah masuk dulu ke folder yang ada file testnya
		perintah : 	"go test" = untuk menjalankan test
					"go test -v" = jika ingin melihat lebih detail function apasaja yang di test
					"go test -v -run <nama_function>" = untuk memilih function mana yang akan di test / di running
		untuk mengetest tanpa harus masuk folder yang memiliki file test
		perintah : 	"go test ./..." = menjalankan test semua folder/package (semua perintah sama dengan atas tinggal tambah "./..." di akhirnya)

	Menggagalkan test
		mengagalkan test menggunakan panic bukan hal yang bagus, karena jika bertemu panic di awal maka program selanjutnya tidak tereksekusi
		perintah : 	function "Fail()" = akan mengagalkan unit test, tapi tetap melanjutkan unit test programnya. namun di akhir ketika selesai maka unit test akan dianggap gagal
					function "FailNow()" = akan mengagaglkan unit tesu saat itu juga tanpa melanjutkan test unit selanjutnya
					function "Error()" = akan memberi pesan eror dan langsung menjalankan function "Fail()"
					function "Fatal()" = mirip dengan "Error()" tapi akan langsung memanggil "FailNow()"
*/

/*
Assertion
	disarankan menggunakan assertion untuk pengecekan banyak data, tidak disarankan menggunakan if else
	Go-Lang tidak menyediakan package untuk assertion, jadi harus menggunakan libarary
		library yang sering di pakai : Testify = untuk melakukan assertion terhadap result data di uni test (https://github.com/stretchr/testify)
		untuk menginstallnya meggunakan perintah : go get github.com/stretchr/testify
	untuk lebih lengkap bisa baca di linknya tentang penjelasannya
	library testyfy menyediakan 2 package untuk assertion yaitu :	assert = jika menggunakan assert akan memanggil Fail()
																	require = jika menggunakan require akan memanggil FailNow()
*/

/*
Skip Test
	membatalkan test
		perintah : Skip() = untuk membatalkan
*/

/*
Before After Test
	perintah : testing.M = digunakan untuk mengatur eksekusi unit test namun hal ini juga bisa digunakan untuk melakukan before aftr unit test
	untuk mengatur eksekusi unit test, maka buat function bernama "TestMain(m testing.M)" jika ada function ini maka akan menjalankan ini saat menjalankan unit test di sebuah package
	dengan ini kita bisa mengatur before after seperti yang diinginkan
	hanya di eksekusi 1x dalam 1 package
*/

/*
Sub Test
	adalah function unit test di dalam function / function di dalam function
	untuk membuat sub test bisa gunakan perintah : function "Run()"
		hanya menjalankan subtest saja perintah :	go test -run <TestFUnctionnya>/<NamaSubTest>
													go test -run /<NamaSubTest>
*/

/*
Table Test
	sub test tadi biasanya digunakan uuntuk membuat test dengan konsep tabel
	tabel test yaitu dimana kita menyediakan data berubah slice yang berisi parameter dan ekpetasi hasil dari unit test
	lalu slice di iterasi menggunakan sub test
*/

/*
Mock
	mock = 	object yang sudah kita program dengan expectasi tertentu sehingga saat di panggil, akan menghasilkan data yang sudah di program di awal
			salah satu teknik dalam unit testing, dimana kita bisa membuat mock object dari suatu object yang sulit untuk di testing
		contoh :	saat ingin membuat unit test, akan tetapi ada kode program yang harus dipanggil "API Call ke Third party service". akan sulit untuk di test, karena harus selalu memanggil third party service, dan belum tentu responsenya sesuai dengan apa yang diinginkan
				untuk kasus seperti itu cocok menggunakan "mock object"
		menggunakan library testify sama seperti assertion
		catatan : harus desain coding yang baik
		untuk contohnya bisa buka package "entity", "repository", "service"

	(catatan : maih eror harus di coba lagi)
*/

/*
Benchmark
	Benchmark = mekanisme menghitung kecepeatan performa kode aplikasi
				ddilakukan secara otomatis
				tidak perlu menentukan lamanya iterasi, karena sudah di atur di testing.B bawaan testing package
		perintah : "testing.B" = sama seperti "testing.T" ada fail(),failnow(), dll. tetapi yang mebedakannya adalah attribute "N" = digunakan untuk melakukan total iterasi sebuah benchmark
		sama seperti di test harus di awali dengan kata "Test" tetapi di benchmark di awali dengan kata "Benchmark" dan memiliki parameter (b *testing.B) tidak boleh ada return value dan filenya juda sama di akhiri dengan "_test.go"
		contoh di folder "helper" file "benchmark_test.go"
		cara running : 	go test -v bench=. //dengan unitestnya
						go test -v -run=NoMathUnitTest -bench=. //benchnya saja
						go test -v -run=NoMathUnitTest -bench=<nama benchmark> //untuk memanggil 1 bendhmark
						go test -v -bench=. ./... //menjalankan semua banchmark di semua package

*/

/*
Sub Benchmark
	sama saja dengan subtest b.Run()
*/

/*
Table Benchmark
	sama saja dengan table benchmark b.Run()
*/

func main() {
	result := helper.SayHallo("Wandi")
	fmt.Println(result)
}
