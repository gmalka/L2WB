package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func downloadStite(url *url.URL) error {
	var (
		filename string
	)
	if url.Scheme == "" {
		url.Scheme = "http"
	}

	strs := strings.Split(url.Path, "/")
	if len(strs) == 0 || (len(strs) == 1 && strs[0] == "") {
		filename = "defaultname"
	} else {
		filename = strs[len(strs) - 1]
	}

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0660)
	if err != nil {
		return fmt.Errorf("cant open file with name %v: %v", filename, err)
	}
	defer f.Close()

	resp, err := http.Get(url.String())
	if err != nil {
		return fmt.Errorf("cant make get request: %v", err)
	}
	buf := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			return fmt.Errorf("cant read from source: %v", err)
		}

		f.Write(buf[:n])

		if err == io.EOF || n == 0 {
			break
		}
	}

	return nil
}

func main() {
	if len(os.Args) <= 1 {
		log.Panicln("Incorrect number of arguments")
	}

	for _, v := range os.Args[1:] {
		u, err := url.Parse(v)
		if err != nil {
			log.Panicf("Cant parse url %s: %s\n", v, err)
		}

		err = downloadStite(u)
		if err != nil {
			log.Panicf("Cant download site with url %s: %s\n", u.String(), err)
		}
	}
}
