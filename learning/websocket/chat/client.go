package main

import (
	"bufio"
	"code.google.com/p/go.net/websocket"
	"fmt"
	"log"
	"os"
)

func main() {
	origin := "http://localhost/"
	url := "ws://localhost:8800/chat"
	ws, err := websocket.Dial(url, "", origin)

	if err != nil {
		log.Fatal(err)
	}

	var msg = make([]byte, 1024)
	go func() {
		for {
			var n int
			if n, err = ws.Read(msg); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Received: %s", msg[:n])
		}
	}()

	var reader = bufio.NewReader(os.Stdin)
	var buf = make([]byte, 1024)
	for {
		var n int
		if n, err = reader.Read(buf); err != nil {
			log.Fatal(err)
		}
		if _, err := ws.Write(buf[:n]); err != nil {
			log.Fatal(err)
		}
	}
}
