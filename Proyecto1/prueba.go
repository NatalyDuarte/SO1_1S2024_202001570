package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:2000@tcp(192.168.0.15:3306)/Proyecto")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("insert into raminfo(freeram,boundram,time) values (?,?,?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(8000000, 9128912, time.Now())
	if err != nil {
		panic(err)
	}

	println("Usuario insertado correctamente")
	db.Close()
}
