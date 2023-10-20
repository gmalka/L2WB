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
		b *bufio.Reader
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
		b = bufio.NewReader(os.Stdin)
	} else {
		b = bufio.NewReader(f)
	}
	result := make([]string, 0, 10)
	for {
		str, err := b.ReadString('\n')

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Cant read from file: %v\n", err)
			return
		}

		if !s || (s && strings.Contains(str, d)) {
			strs := strings.Split(str, d)

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
		//fmt.Printf("|%s %d|\n", v, len(v))
		if v[len(v) - 1] != '\n' {
			fmt.Println(v)
		} else {
			fmt.Print(v)
		}
	}
}