package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Can't connect !")
			conn.Close()
			continue
		}
		fmt.Println("Connected")
		bufReader := bufio.NewReader(conn)
		fmt.Println("Start")
		go func(conn net.Conn) {
			defer conn.Close()
			for {
				rbyte, err := bufReader.ReadByte()
				if err != nil {
					fmt.Println("Can't read!", err)
					break
				}
				fmt.Printf(string(rbyte))
			}
		}(conn)

		<-time.After(time.Second * 10)

		fmt.Println("Writing to console")
		_, err = conn.Write([]byte("Hello from server\n"))
		if err != nil {
			fmt.Println("Can't write!", err)
		}
	}
}
