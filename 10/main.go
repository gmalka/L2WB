package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	var (
		timeout time.Duration
		conn    net.Conn
		err     error
	)
	flag.DurationVar(&timeout, "timeout", time.Second*3, "Timeout duration")

	flag.Parse()

	args := flag.Args()

	if len(args) < 2 {
		log.Println("Incorrect number of arguments")
		return
	}

	port := args[len(args)-1]
	host := args[len(args)-2]

	t := time.After(timeout)
	for {
		conn, err = net.DialTimeout("tcp", fmt.Sprintf("%s:%s", host, port), timeout)
		if err != nil {
			select {
			case <-t:
				log.Printf("Cant connect to %s", fmt.Sprintf("%s:%s", host, port))
				return
			default:
				continue
			}
		} else {
			fmt.Printf("Connected to %s\n", fmt.Sprintf("%s:%s", host, port))
			break
		}
	}
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		b := make([]byte, 1024)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n, err := conn.Read(b)
				if n > 0 {
					fmt.Printf("#From Server#: %s", string(b[:n]))
				}
				if err != nil {
					log.Printf("Cant red from remote: %v", err)
					cancel()
					return
				}
			}
			time.Sleep(time.Second)
		}
	}()

	go func() {
		r := bufio.NewReader(os.Stdin)
		for {
			buf, err := r.ReadBytes('\n')
			if err == io.EOF {
				fmt.Println("Exiting...")
				cancel()

				return
			}

			if err != nil {
				log.Printf("Cant read bytes: %s\n", err)
				return
			}

			_, err = conn.Write(buf)
			if err != nil {
				fmt.Println("Server stopped...")
				return
			}
		}
	}()

	<-ctx.Done()
	fmt.Println("Server stopped")
	conn.Close()
}
