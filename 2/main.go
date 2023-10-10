package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(Decode(`a4bc2d3e`))
}

func Decode(str string) (string, error) {
	var (
		symbol rune
		esc    bool
	)
	result := make([]rune, 0, len(str))

	l := 1
	n := 1
	for _, v := range str {
		if v == '\\' && !esc {
			esc = true
		} else if v >= '0' && v <= '9' && !esc {
			if symbol == 0 {
				return "", errors.New("incorrect string")
			}
			n *= 10
			l *= 10
			l += int(v - '0')
		} else {
			fmt.Println(l % n, l, n)
			for i := l % n; i > 0; i-- {
				result = append(result, symbol)
			}
			symbol = v
			l = 1
			n = 1
			esc = false
		}
	}
	if esc {
		return "", errors.New("incorrect string")
	}
	for i := l % n; i > 0; i-- {
		result = append(result, symbol)
	}

	return string(result), nil
}