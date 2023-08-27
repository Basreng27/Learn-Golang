package helper

//================================================== Brenckmark ==================================================
// /*
// Benchmark
// 	Benchmark = mekanisme menghitung kecepeatan performa kode aplikasi
// 				ddilakukan secara otomatis
// 				tidak perlu menentukan lamanya iterasi, karena sudah di atur di testing.B bawaan testing package
// 		perintah : "testing.B" = sama seperti "testing.T" ada fail(),failnow(), dll. tetapi yang mebedakannya adalah attribute "N" = digunakan untuk melakukan total iterasi sebuah benchmark
// 		sama seperti di test harus di awali dengan kata "Test" tetapi di benchmark di awali dengan kata "Benchmark" dan memiliki parameter (b *testing.B) tidak boleh ada return value dan filenya juda sama di akhiri dengan "_test.go"
// 		contoh di folder "helper" file "benchmark_test.go"
// 		cara running : 	go test -v bench=. //dengan unitestnya
// 						go test -v -run=NoMathUnitTest -bench=. //benchnya saja
// 						go test -v -run=NoMathUnitTest -bench=<nama benchmark> //untuk memanggil 1 bendhmark
// 						go test -v -bench=. ./... //menjalankan semua banchmark di semua package

// */

// func BenchmarkSayHallo(b *testing.B) { //membuat brenchmark
// 	for i := 1; i < b.N; i++ {
// 		SayHallo("Wandi")
// 	}
// }

//================================================== Sub Brenckmark ==================================================
