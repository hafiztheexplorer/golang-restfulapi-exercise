package app

import (
	"database/sql"
	"golang-restfulapi-exercise/helper"
	"time"
)

func NewDB() *sql.DB {
	// masukkan drivernya "mysql"
	// rumusnya seperti ini Open(driverName string, dataSourceName string)
	// datasource kalau bingung coba buka web github dari sql.DB nanti
	// akan dikasih tahu cara menulisnya
	db, error := sql.Open("mysql", "root@tcp(localhost:3306)/exercise_golang_restfulapi")
	helper.PanicIfError(error)

	db.SetConnMaxIdleTime(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}
