package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type sorter interface {
	add(str string)
	sort(bool) []string
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
	var (
		left StringSliceFlag
	)

	flag.Var(&left, "k", "Multiple values for flag -k")

	flag.Parse()

	fmt.Println(left)
}