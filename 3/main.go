package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"math"
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

type StringSliceFlag []int

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

	*s = result
	return nil
}

func main() {
	t := time.Now()
	var (
		left StringSliceFlag
		m    bool
		r    bool
		n    bool
	)

	flag.Var(&left, "k", "")
	flag.BoolVar(&m, "u", false, "")
	flag.BoolVar(&r, "r", false, "")
	flag.BoolVar(&n, "n", false, "")

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

	if len(left) != 0 {
		if len(left) == 1 {
			left = append(left, math.MaxInt)
		}
		if n {
			ss = newKnumsorter(left[0], left[1])
		} else {
			ss = newkstrsorter(left[0], left[1])
		}
	} else {
		left = append(left, 1)
		if n {
			ss = newNumSorter()
		} else {
			ss = newStrSorter()
		}
	}

	if n && m {
		so = newMNsorter(b, ss, r, left)
	} else if m {
		so = newMsorter(b, ss, r, left)
	} else {
		so = newSSorter(b, ss, r)
	}

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
