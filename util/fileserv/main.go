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
var port int

func init() {
	flag.IntVar(&port, "port", 8000, "port number")
}

func main() {
	flag.Parse()

	server.Port = strconv.Itoa(port)

	http.Handle("/", http.FileServer(http.Dir("./")))

	log.Println("Starting server on:" + server.Port)
	err := http.ListenAndServe("0.0.0.0:"+server.Port, nil)

	if err != nil {
		log.Printf("Server failed: ", err.Error())
	}
}
