package database

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func NewDb(userName string, password string, protocol string, address string, database string) *sql.DB {
	db, err := sql.Open("mysql", userName + ":" + password + "@" + protocol + "(" + address + ")/" + database + "?charset=utf8")
	checkErr(err)
	return db
}

func CloseDb(db *sql.DB) {
	err := db.Close()
	checkErr(err)
}

