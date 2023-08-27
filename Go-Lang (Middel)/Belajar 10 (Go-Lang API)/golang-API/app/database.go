package app

import (
	"database/sql"
	"golangapi/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:8080)/belajar_golang_api") //membuka database dengan mysql
	helper.PanicIfError(err)

	//connection pulling
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
