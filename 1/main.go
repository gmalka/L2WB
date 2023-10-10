package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	addr := "pool.ntp.org"
	t, err := ntp.Time(addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cant get time from %s: %v", addr, err)
		os.Exit(1)
	}

	fmt.Println(t.Format("2006-01-02 15:04:05"))
}