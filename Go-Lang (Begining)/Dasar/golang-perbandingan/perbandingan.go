package main

import "fmt"

//================================================== Oprator Perbandingan ==================================================
// func main() {
// 	/*
// 		oprator perbandingan
// 			> = lebih dari
// 			< = kurang dari
// 			>= = lebih dari sama dengan
// 			<= = kurang dari sama dengan
// 			== = sama dengan
// 			!= tidak sama dengan
// 	*/

// 	a := "Wandi"
// 	b := "Wandi"
// 	var c bool = a == b //hasil true karena a sama dengan b
// 	var d bool = a != b //hasil false kariena a tidak sama dengan b
// 	fmt.Println(c)
// 	fmt.Println(d)

// 	var v1 = 100
// 	var v2 = 100
// 	fmt.Println(v1 > v2)  //hasil lebih besar
// 	fmt.Println(v1 < v2)  //hasil lebih kecil
// 	fmt.Println(v1 == v2) //hasil sama dengan
// 	fmt.Println(v1 != v2) //hasil tidak sama dengan

// }

//================================================== Oprator Boolean ==================================================
func main() {
	/*
		oprator boolean
			&& = dan
			|| = atau
			! = kebalikan
	*/

	var Nakhir = 90
	var absen = 80

	var lulusna bool = Nakhir > 80
	var lulusa bool = absen > 80

	var hsl bool = lulusna && lulusa

	fmt.Println(hsl)
}
