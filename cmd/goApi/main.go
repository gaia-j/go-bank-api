package main

import (
	"github.com/gaia-j/go-bank-api/api"
	"github.com/gaia-j/go-bank-api/internal/database"
	"log"
)

func main() {

	database.ConnectDatabase()

	err := api.Start()

	if err != nil {
		log.Fatal(err)
	}

}
