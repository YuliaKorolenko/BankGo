package main

import (
	"awesomeProject/base"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {

	db, err := base.NewPostgresDB(base.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "postgres",
	})

	connection := &base.PostgresMethods{Db: db}
	hadler := &Handler{postgresMethods: *connection}

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err)
	}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/create", hadler.CreateAccount)
	router.HandleFunc("/deposit", hadler.DepositMoney)
	router.HandleFunc("/reserve", hadler.Reserve)
	router.HandleFunc("/balance", hadler.PersonalBalance)

	log.Fatal(http.ListenAndServe(":8080", router))
}
