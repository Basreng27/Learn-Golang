package entity

import (
	"database/sql"
	"time"
)

// untuk table comment di database
type Comment struct {
	Id             int32
	Email, Comment string
}

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
