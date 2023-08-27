// ================================================== Driver Database ==================================================
// package golang_database

// import (
// 	"testing"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func TestEmpty(t *testing.T) {

// }

// ================================================== Koneksi Ke Database ==================================================
// package golang_database

// import (
// 	"database/sql"
// 	"testing"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func TestConnection(t *testing.T) {
// 	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar-golang-database") //membuat koneksi kedatabase //3306 default localhost

// 	if err != nil { //mengecek eror atau tidak
// 		panic(err)
// 	}

// 	defer db.Close() //menutup database
// }

package golang_database