package main

import "fmt"

//================================================== If ==================================================
// func main() {
// 	nama := "asdfdghfjdgsa"

// 	if nama == "Rayandra" { //if
// 		fmt.Println("Hallo Rayandra")
// 	} else if nama == "Wandi" { //else if
// 		fmt.Println("Hallo Wandi")
// 	} else { //else
// 		fmt.Println("Hallo Marselana")
// 	}

// 	//if short statement
// 	if length := len(nama); length > 5 {
// 		fmt.Println("Nama Terlalu Panjang")
// 	} else {
// 		fmt.Println("Nama Sudah Benar")
// 	}
// }

//================================================== Switch ==================================================
func main() {
	nama := "Wandi"

	switch nama {
	case "Rayandra":
		fmt.Println("Hallo Rayandra")

	case "Wandi":
		fmt.Println("Hallo Wandi")

	default:
		fmt.Println("Hallo Marselana")
	}

	//switch short statement
	switch length := len(nama); length > 5 {
	case true:
		fmt.Println("Nama Kurang")

	case false:
		fmt.Println("Nama Sudah Benar")
	}

	//switch tanpa kondisi
	length := len(nama)
	switch {
	case length > 5:
		fmt.Println("Nama Kurang")

	case length < 5:
		fmt.Println("Nama Sudah Benar")
	}
}
