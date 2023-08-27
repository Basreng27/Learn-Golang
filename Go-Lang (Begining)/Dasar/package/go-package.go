// ================================================== Package & Import ==================================================
/*
package = 	tempat yang digunakan untuk mengorganisir kode program yang dibuat
			bisa merapihkan kode yang dibuat / grouping
			sebenarnya hanya direktori folder disistem oprasi

membuat package = 	1. buat folder dengan nama apapun, contoh : helper
					2. kemudian buat file dengan nama apapun, contoh : SayHallo.go
					3. kemudai liat isi file SayHallo.go untuk lebih lengkap
*/
package main

import (
	"fmt"
	"package\helper"
)

//import = untuk mengakses file lain yang berbeda package atau yang berda di luat package
// import "golang-package\helper"

func main() {
	helper.SayHallo("Wandi")
	fmt.Println("BLok")
}
