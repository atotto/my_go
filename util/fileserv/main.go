package main

import (
	"log"
	"net/http"
)

const PORT = "8000"

func main() {

	http.Handle("/", http.FileServer(http.Dir("./")))

	log.Println("Starting server on:", PORT)
	err := http.ListenAndServe("0.0.0.0:"+PORT, nil)

	if err != nil {
		log.Printf("Server failed: ", err.Error())
	}
}
