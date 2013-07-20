package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	Port string
}

var server Server = Server{}

func init() {
	var port int
	flag.IntVar(&port, "port", 8000, "port number")

	server.Port = strconv.Itoa(port)
}

func main() {

	http.Handle("/", http.FileServer(http.Dir("./")))

	log.Println("Starting server on:" + server.Port)
	err := http.ListenAndServe("0.0.0.0:"+server.Port, nil)

	if err != nil {
		log.Printf("Server failed: ", err.Error())
	}
}
