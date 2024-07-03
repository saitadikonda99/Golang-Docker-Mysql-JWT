package db

import (
	"database/sql"
	"log"
	"github.com/go-sql-driver/mysql"
)


func NewMySQLstorage(c mysql.Config) (*sql.DB, error) {

	db, err := sql.Open("mysql", c.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}
	
	return db, nil
}
