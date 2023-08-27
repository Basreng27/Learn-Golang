package helper

//================================================== Unit testing ==================================================
// /*
// Unit Testing
// 	Aturan File test /testing unit
// 		untuk membuat file unit test maka harus menggunakan "_test" diakhiran
// 		contoh : jika membuat file "hello.go" maka untuk membuat testnya menggunakan nama file "hello_test.go" (contoh pada package "helper")
// 			"hello_test.go" digunakan hanya untuk mengetest file "hello.go" saja

// 	Aturan Fucntion unit test
// 		untuk membuat function unit test maka harus diawali dengan nama "Test"
// 		contoh : jika membuat function "SayHallo" untuk mengetest functionnya harus di buat unit test dengan nama "TestSayHallo"
// 			untuk lanjutannya bebas tapi harus diawali deng "Test"
// 				lalu harus memiliki parameter (t *testing.T) dan tidak mengembalikan value (contoh bisa di liat di package "helper" file "hello_test.go")

// 	Aturan untuk menjalankan Unit test
// 		sebelum melakukan perintah di bawah masuk dulu ke folder yang ada file testnya
// 		perintah : 	"go test" = untuk menjalankan test
// 					"go test -v" = jika ingin melihat lebih detail function apasaja yang di test
// 					"go test -v -run <nama_function>" = untuk memilih function mana yang akan di test / di running
// 		untuk mengetest tanpa harus masuk folder yang memiliki file test
// 		perintah : 	"go test ./..." = menjalankan test semua folder/package (semua perintah sama dengan atas tinggal tambah "./..." di akhirnya)

// 	Menggagalkan test
// 		mengagalkan test menggunakan panic bukan hal yang bagus, karena jika bertemu panic di awal maka program selanjutnya tidak tereksekusi
// 		perintah : 	function "Fail()" = akan mengagalkan unit test, tapi tetap melanjutkan unit test programnya. namun di akhir ketika selesai maka unit test akan dianggap gagal
// 					function "FailNow()" = akan mengagaglkan unit tesu saat itu juga tanpa melanjutkan test unit selanjutnya
// 					function "Error()" = akan memberi pesan eror dan langsung menjalankan function "Fail()"
// 					function "Fatal()" = mirip dengan "Error()" tapi akan langsung memanggil "FailNow()"
// */

// //testing
// func TestSayHallo(t *testing.T) { //membuat function testing unit
// 	result := SayHallo("Wandi")  //mengeksekusi function "SayHallo"
// 	if result != "Hallo Wandi" { //di cek sama tidak hallonya
// 		//unit test gagal
// 		panic("Hasilnya bukan 'Hallo Wandi'")
// 	}
// }

// //testing 2 function
// func TestSayHalloLele(t *testing.T) {
// 	result := SayHallo("Lele")
// 	if result != "Hallo Lele" {
// 		//unit test gagal
// 		panic("Hasilnya bukan 'Hallo Wandi'")
// 	}
// }

// // testing Fail()
// func TestSayHallo(t *testing.T) {
// 	result := SayHallo("Wandi")
// 	if result != "Hallo Wandi" {
// 		//unit test gagal
// 		// panic("Hasilnya bukan 'Hallo Wandi'")
// 		t.Fail() //function test gagal (bisa dilihat penjelasannya di "main.go")
// 	}

// 	fmt.Println("TestSayHallo Selesai") //untuk memberitahu bahwa program tetap jalan //ini akan jalan
// }

// // testing FailNow()
// func TestSayHalloLele(t *testing.T) {
// 	result := SayHallo("Lele")
// 	if result != "Hallo Lele" {
// 		//unit test gagal
// 		// panic("Hasilnya bukan 'Hallo Wandi'")
// 		t.FailNow() //functin test gagal
// 	}

// 	fmt.Println("TestSayHallo Selesai") //untuk memberitahu bahwa program tetap jalan //ini tidak akan jalan
// }

// testing Fail()
// func TestSayHallo(t *testing.T) {
// 	result := SayHallo("Wandi")
// 	if result != "Hallo Wandi" {
// 		//unit test gagal
// 		// panic("Hasilnya bukan 'Hallo Wandi'")
// 		// t.Fail()
// 		t.Error("Harusnya Hallo Wandi")
// 	}

// 	fmt.Println("TestSayHallo Selesai") //untuk memberitahu bahwa program tetap jalan //ini akan jalan
// }

// // testing FailNow()
// func TestSayHalloLele(t *testing.T) {
// 	result := SayHallo("Lele")
// 	if result != "Hallo Lele" {
// 		//unit test gagal
// 		// panic("Hasilnya bukan 'Hallo Wandi'")
// 		// t.FailNow()
// 		t.Fatal("Harusnya Hallo Lele")
// 	}

// 	fmt.Println("TestSayHallo Selesai") //untuk memberitahu bahwa program tetap jalan //ini tidak akan jalan
// }

// ================================================== Assert ==================================================
// /*
// Assertion
// 	disarankan menggunakan assertion untuk pengecekan banyak data, tidak disarankan menggunakan if else
// 	Go-Lang tidak menyediakan package untuk assertion, jadi harus menggunakan libarary
// 		library yang sering di pakai : Testify = untuk melakukan assertion terhadap result data di uni test (https://github.com/stretchr/testify)
// 		untuk menginstallnya meggunakan perintah : go get github.com/stretchr/testify
// 	untuk lebih lengkap bisa baca di linknya tentang penjelasannya
// 	library testyfy menyediakan 2 package untuk assertion yaitu :	assert = jika menggunakan assert akan memanggil Fail()
// 																	require = jika menggunakan require akan memanggil FailNow()
// */

// /assert
// func TestHalloAssertionAssert(t *testing.T) {
// 	result := SayHallo("Wandi")

// 	//equal = sama dengan (==)(catatan perintah sama seperti opraotr aritatika tetapi menggunakan b.inggris sma seperti "equal" artinya sama dengan(==))
// 	assert.Equal(t, "Hello Wandi", result, "Pesan isi harus 'Hallo Wandi'") //mengecek apakan result == "Hello Wandi". sama seperti if else tapi lebih singkat

// 	fmt.Println("Perintah dijalankan") //memberi tahu perintah di jalankan
// }

// // require
// func TestHalloAssertionRequire(t *testing.T) {
// 	hsl := SayHallo("Lele")

// 	require.Equal(t, "Hello Lele", hsl, "Pesan isi harus 'Hallo Lele'")

// 	fmt.Println("Perintah dijalankan")
// }

// ================================================== Skip Test ==================================================
// /*
// Skip Test
// 	membatalkan test
// 		perintah : Skip() = untuk membatalkan
// */

// func TestSkip(t *testing.T) {
// 	if runtime.GOOS == "windows" { //mengecek os yang di gunakan
// 		t.Skip("Unit Test tidak bisa di jalankan di Windows") //jika ada skip maka program selanjutnya tidak di eksekusi
// 	}

// 	result := SayHallo("Wandi")                                         //tidak di eksekusi
// 	require.Equal(t, "Hallo Wandi", result, "Harus sama 'Hallo Wandi'") //tidak di eksekusi
// }

// ================================================== Before After Test ==================================================
// /*
// Before After Test
// 	perintah : testing.M = digunakan untuk mengatur eksekusi unit test namun hal ini juga bisa digunakan untuk melakukan before aftr unit test
// 	untuk mengatur eksekusi unit test, maka buat function bernama "TestMain(m testing.M)" jika ada function ini maka akan menjalankan ini saat menjalankan unit test di sebuah package
// 	dengan ini kita bisa mengatur before after seperti yang diinginkan
// 	hanya di eksekusi 1x dalam 1 package
// */

// func TestMain(m *testing.M) { //seperti construct__ di php
// 	//before / biasanya di isikan koneksi ke database
// 	fmt.Println("Sebelum Unit Test")

// 	m.Run() //mengeksekusi/menjalankan semua unit test

// 	//after
// 	fmt.Println("Setelah Unit Test")
// }

// func TestSkip(t *testing.T) {
// 	if runtime.GOOS == "windows" {
// 		t.Skip("Unit Test tidak bisa di jalankan di Windows")
// 	}

// 	result := SayHallo("Wandi")
// 	require.Equal(t, "Hallo Wandi", result, "Harus sama 'Hallo Wandi'")
// }

// ================================================== Sub Test ==================================================
// /*
// Sub Test
// 	adalah function unit test di dalam function / function di dalam function
// 	untuk membuat sub test bisa gunakan perintah : function "Run()"
// 		hanya menjalankan subtest saja perintah :	go test -run <TestFUnctionnya>/<NamaSubTest>
// 													go test -run /<NamaSubTest>
// */

// func TestSubTest(t *testing.T) {
// 	t.Run("Wandi", func(t *testing.T) { //buat sub unit test, "Wandi" = nama dari subtestnya
// 		result := SayHallo("Wandi")
// 		require.Equal(t, "Hallo Wandi", result, "Isi Harus 'Hallo Wandi'")
// 	})

// 	t.Run("Marselana", func(t *testing.T) {
// 		resul := SayHallo("Marselana")
// 		require.Equal(t, "Hallo Marselana", resul)
// 	})
// }

// ================================================== Table Test ==================================================
// /*
// Table Test
// 	sub test tadi biasanya digunakan uuntuk membuat test dengan konsep tabel
// 	tabel test yaitu dimana kita menyediakan data berubah slice yang berisi parameter dan ekpetasi hasil dari unit test
// 	lalu slice di iterasi menggunakan sub test
// */

// func TestHalloTable(t *testing.T) {
// 	//definisikan isi tabelnya
// 	tests := []struct { //membuat structnya
// 		name     string //isi structnya
// 		request  string //isi structnya
// 		expected string //isi structnya
// 	}{
// 		{
// 			name:     "SayHallo(Wandi)",
// 			request:  "Wandi",
// 			expected: "Hallo Wandi",
// 		}, {
// 			name:     "SayHallo(Marselana)",
// 			request:  "Marselana",
// 			expected: "Hallo Marselana",
// 		},
// 	}

// 	// ulangi untuk mengecek
// 	for _, test := range tests { //mengulangi sebanyak data yang ada pada "tests"
// 		t.Run(test.name, func(t *testing.T) { //mengambil nama untuk sub testnya dari struct //dan buat subtest dengan "t.Run()"
// 			result := SayHallo(test.request)        //mengambil request dari struct
// 			require.Equal(t, test.expected, result) //mengambil expected dari struct
// 		})
// 	}
// }
