package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"       // <-- GANTI SESUAI PUNYA KAMU
	dbname   = "tugas13db" // <-- GANTI NAMA DATABASE KAMU
)

func ConnectDB() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic("Gagal connect ke database: " + err.Error())
	}

	fmt.Println("Berhasil connect ke database!")
}
