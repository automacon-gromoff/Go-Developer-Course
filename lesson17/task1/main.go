package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	wg := sync.WaitGroup{}
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			handleConn(conn)
		}()
	}
	log.Println("завершение работы")
}

func handleConn(c net.Conn) {
	defer c.Close()
	var i int
	for {
		_, err := io.WriteString(c, fmt.Sprintf("%d\n", i))
		if err != nil {
			log.Println(err)
			return
		}
		i++
		time.Sleep(time.Second)
	}
}
