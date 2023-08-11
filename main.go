package main

import (
	"gopnsorders/store"
	"log"
)

func main() {
	conn, err := store.GetConnection("./db.db")

	if err != nil {
		log.Fatal(err)
	}

	if err := conn.InitializeTables(); err != nil {
		log.Fatal(err)
	}

	log.Println("Got Connected!")
}
