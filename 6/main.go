package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type columns []int

func (c *columns) String() string {
	return fmt.Sprintf("%v", *c)
}

func (c *columns) Set(str string) error {
	strs := strings.Split(str, ",")

	for _, v := range strs {
		st := strings.Split(v, "-")
		if len(st) == 2 {
			begin, err := strconv.Atoi(st[0])
			if err != nil {
				return fmt.Errorf("cant parse -f flag: %v", err)
			}
			end, err := strconv.Atoi(st[1])
			if err != nil {
				return fmt.Errorf("cant parse -f flag: %v", err)
			}
			for begin <= end {
				*c = append(*c, begin-1)
				begin++
			}

		} else if len(st) == 1 {
			num, err := strconv.Atoi(st[0])
			if err != nil {
				return fmt.Errorf("cant parse -f flag: %v", err)
			}
			*c = append(*c, num-1)
		} else {
			return errors.New("cant parse -f flag")
		}
	}

	sort.Ints(*c)
	return nil
}

func main() {
	var (
		s bool
		d string
		c columns
	)
	flag.BoolVar(&s, "s", false, `-s - "separated" - только строки с разделителем`)
	flag.StringVar(&d, "d", "\n", `-d - "delimiter" - использовать другой разделитель`)
	flag.Var(&c, "f", `"fields" - выбрать поля (колонки)`)
	flag.Parse()

	if len([]rune(d)) > 1 {
		log.Println("Bad delimiter")
		return
	}
	f, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		log.Printf("Cant open file: %v\n", err)
		return
	}
	b := bufio.NewReader(f)
	result := make([]string, 0, 10)
	for {
		str, err := b.ReadString('\n')

		if err != nil && err != io.EOF {
			log.Printf("Cant read from file: %v\n", err)
			return
		}
		str = str[:len(str)-1]

		strs := strings.Split(str, d)

		if !(s && len(strs) == 1) {
			st := strings.Builder{}

			i := 0
			for _ = i; i < len(strs) - 1 && i < len(c) - 1; i++ {
				st.WriteString(strs[c[i]])
				st.WriteString(d)
			}
			if i < len(c) && c[i] < len(strs) {
				st.WriteString(strs[c[i]])
			}

			result = append(result, st.String())
		}

		if err == io.EOF {
			break
		}
	}

	for _, v :=range result {
		fmt.Println(v)
	}
}
