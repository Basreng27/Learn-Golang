// ================================================== Database Pooling ==================================================
// membuat koneksi kedatabase agar bisa di pakai terus menerus
package golang_database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar-golang-database?parseTime=true") //koneksi database

	if err != nil {
		panic(err)
	}

	//wajib ada
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
