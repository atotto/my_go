package main

import (
	"code.google.com/p/go.net/websocket"
	"log"
	"net/http"
)

type Connection struct {
	ws          *websocket.Conn
	sendMessage chan string
}

type Serv struct {
	conns      map[*Connection]bool
	register   chan *Connection
	unregister chan *Connection
	message    chan string
	name       string
}

func New(servname string) *Serv {
	s := new(Serv)

	s.conns = make(map[*Connection]bool)
	s.register = make(chan *Connection)
	s.unregister = make(chan *Connection)
	s.message = make(chan string)

	s.name = servname

	return s
}

func (s *Serv) Run() {
	log.Println("server start")
	go func() {
		for {
			select {
			case c := <-s.register:
				s.conns[c] = true
				log.Println("register")
			case c := <-s.unregister:
				delete(s.conns, c)
				log.Println("unregister")
			case msg := <-s.message:
				for c := range s.conns {
					c.sendMessage <- msg
				}
			}
		}
	}()
}

func (s *Serv) clientHandler(ws *websocket.Conn) {
	sendMessage := make(chan string)
	c := &Connection{ws: ws, sendMessage: sendMessage}

	s.register <- c
	defer func() {
		s.unregister <- c
		close(c.sendMessage)
	}()

	go sender(c)
	receiver(c, s.message)
}

func sender(c *Connection) {
	for {
		buf := <-c.sendMessage
		err := websocket.Message.Send(c.ws, buf)
		if err != nil {
			log.Println(err)
			break
		}
		//log.Printf("send:%q\n", buf)
	}
}

func receiver(c *Connection, message chan string) {
	var buf string
	for {
		err := websocket.Message.Receive(c.ws, &buf)
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("recv:%q\n", buf)
		message <- buf
	}
}

func main() {
	serv := New("chatserver")
	serv.Run()

	http.Handle("/chat", websocket.Handler(serv.clientHandler))
	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(":8800", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
