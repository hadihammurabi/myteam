package model

import (
	. "fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Connection *sql.DB

func Init() {
	db, err := sql.Open("mysql", "root:root@/myteam")

	if err != nil {
		Printf("SQL Error %v\n", err)
		db.Close()
	}

	Connection = db
}