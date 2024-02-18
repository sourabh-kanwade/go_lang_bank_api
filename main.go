package main

import (
	"log"
)

func main() {
	store, err := NewMongoStore()
	if err != nil {
		log.Fatal(err)
	}
	server := NewAPIServer(":3000", store)
	server.Run()
}
