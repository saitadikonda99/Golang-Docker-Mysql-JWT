package api

import (
	"database/sql"
	"net/http"
	"log"

	"github.com/gorilla/mux"
)

type APIserver struct {
	address string
	db      *sql.DB
}

func NewAPIserver(address string, db *sql.DB) *APIserver {
	return &APIserver{
		address: address,
		db: db,
	}
}

func (c *APIserver) Run() error {

	router := mux.NewRouter();
	// subRouter := router.PathPrefix("/api/v1").Subrouter()

	// userHandler := user.NewHandler()
	// userHandler.RegisterRoutes(subRouter)

	log.Println("Server is running on ", c.address)
	log.Println("Connected to database: ")

	return http.ListenAndServe(c.address, router)
}
