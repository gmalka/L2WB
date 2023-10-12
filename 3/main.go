package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type sortstore interface {
	add(str string)
	sort(bool) []string
}

type sorter interface {
	getsorted() ([]string, error)
}

type StringSliceFlag [][]int

func (s *StringSliceFlag) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *StringSliceFlag) Set(value string) error {
	st := strings.Split(value, ",")
	if len(st) > 2 || len(st) == 0 {
		return errors.New("incorrect number of arguments")
	}
	result := []int{}
	for _, v := range st {
		val, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("cant convert value %v to integer: %v", v, err)
		}

		result = append(result, val)
	}

	*s = append(*s, result)
	return nil
}

func main() {
	t := time.Now()
	var (
		left StringSliceFlag
	)

	flag.Var(&left, "k", "Multiple values for flag -k")

	flag.Parse()

	var (
		so sorter
		ss sortstore
	)
	f, err := os.Open("file")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cant open file: %v", err)
		return
	}
	b := bufio.NewReader(f)
	ss = newNumSorter()
	so = newMNsorter(b, ss, false, 0)
	fmt.Println(time.Since(t).Microseconds())

	result, err := so.getsorted()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cant sort file: %v", err)
		return
	}
	fmt.Println(time.Since(t).Microseconds())

	for k, v := range result {
		fmt.Printf("|%d: %v|\n", k, v)
	}
	fmt.Println(time.Since(t).Microseconds())
}