package main

import (
	"log"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/saitadikonda99/Golang-Docker-Mysql-JWT/cmd/api"
	"github.com/saitadikonda99/Golang-Docker-Mysql-JWT/config"
	"github.com/saitadikonda99/Golang-Docker-Mysql-JWT/db"
)

func main() {

	cfg := config.GetConfig

	db, err := db.NewMySQLstorage(mysql.Config {
		User: cfg.DBUser,
		Passwd: cfg.DBPasswd,
		Addr: cfg.DBAddr,
		DBName: cfg.DBName,
		Net: "tcp",
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	initStorage(db, cfg.DBName)

	server := api.NewAPIserver("localhost:8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
		return
	} 
}

func initStorage(db *sql.DB, dbName string) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}