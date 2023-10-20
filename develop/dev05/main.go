package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

type flags struct {
	A    int
	B    int
	C    int
	c    bool
	i    bool
	v    bool
	F    bool
	n    bool
	reg  string
	file string
}

func main() {
	var (
		b *bufio.Reader
	)
	strs := make([]string, 0, 10)

	f, err := parsFlags()
	if err != nil {
		log.Fatalf("Can't parse flags: %v", err)
	}

	if f.file != "" {
		file, err := os.Open(f.file)
		if err != nil {
			log.Fatalf("Can't open file: %v", err)
		}
		b = bufio.NewReader(file)
	} else {
		b = bufio.NewReader(os.Stdin)
	}

	for {
		str, err := b.ReadBytes('\n')
		if len(str) > 0 {
			if str[len(str) - 1] == '\n' {
				str = str[:len(str) - 1]
			}
		}

		strs = append(strs, string(str))

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Can't read string: %v", err)
		}
	}

	r, err := regexp.Compile(f.reg)
	if err != nil {
		log.Fatalf("Can't create regexp: %v", err)
	}
	count := 0
	result := make([]string, 0, 10)
	for i := 0; i < len(strs); i++ {
		if f.c && r.MatchString(strs[i]) {
			count++
		}

		if f.A > 0 && r.MatchString(strs[i]) {
			begin := f.A + i
			if begin > len(strs)-1 {
				begin = len(strs) - 1
			}
			for start := i; start <= begin; start++ {
				result = append(result, strs[start])
			}
			result = append(result, "--")
		} else if f.B > 0 && r.MatchString(strs[i]) {
			result = append(result, "--")
			begin := i - f.B
			if begin < 0 {
				begin = 0
			}
			for end := i; begin <= end; begin++ {
				result = append(result, strs[begin])
			}
		} else if f.C > 0 && r.MatchString(strs[i]) {
			result = append(result, "--")
			begin := i - f.C
			if begin < 0 {
				begin = 0
			}
			end := f.C + i
			if end > len(strs)-1 {
				end = len(strs) - 1
			}
			for ; begin <= end; begin++ {
				result = append(result, strs[begin])
			}
			result = append(result, "--")
		} else if f.n && r.MatchString(strs[i]) {
			result = append(result, fmt.Sprintf("%d:%s", i+1, strs[i]))
		} else if f.F && strings.Contains(strs[i], f.reg) {
			result = append(result, strs[i])
		} else if f.v && !r.MatchString(strs[i]) {
			result = append(result, strs[i])
		} else if f.i && r.MatchString(strs[i]) {
			result = append(result, strs[i])
		} else if r.MatchString(strs[i]) && !f.v {
			result = append(result, strs[i])
		}
	}
	if f.c {
		fmt.Println(count)
		return
	}
	if len(result) > 0 && result[len(result)-1] == "--" {
		result = result[:len(result)-1]
	}
	if len(result) != 0 && result[0] == "--" {
		result = result[1:]
	}
	for _, v := range result {
		fmt.Println(v)
	}
}

func parsFlags() (*flags, error) {
	var (
		reg, file string
	)

	if len(os.Args) < 2 {
		return nil, errors.New("there is not reg expression ")
	}

	A := flag.Int("A", 0, "печатать +N строк после совпадения")
	B := flag.Int("B", 0, "печатать +N строк до совпадения")
	C := flag.Int("C", 0, "печатать ±N строк вокруг совпадения")
	c := flag.Bool("c", false, "(количество строк)")
	i := flag.Bool("i", false, "(игнорировать регистр)")
	v := flag.Bool("v", false, "(вместо совпадения, исключать)")
	F := flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	n := flag.Bool("n", false, "напечатать номер строки")

	flag.Parse()

	if len(flag.Args()) == 1 {
		reg = flag.Args()[len(flag.Args())-1]
	} else {
		reg = flag.Args()[len(flag.Args())-2]
		file = flag.Args()[len(flag.Args())-1]
	}
	return &flags{*A, *B, *C, *c, *i, *v, *F, *n, reg, file}, nil
}