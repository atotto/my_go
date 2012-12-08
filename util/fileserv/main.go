package main

import (
	"net/http"
	"log"
)

const PORT = "8000"

func main() {

	http.Handle("/", http.FileServer(http.Dir("./")))

	log.Println("Starting server on:", PORT)
	err := http.ListenAndServe(":" + PORT, nil)

	if err != nil {
		log.Printf("Server failed: ", err.Error())
	}
}