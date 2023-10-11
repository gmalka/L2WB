package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(decode(``))
}

func decode(str string) (string, error) {
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
			if l > 1 {
				for i := l % n; i > 0; i-- {
					result = append(result, symbol)
				}
			} else if symbol != 0 {
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
	if l > 1 {
		for i := l % n; i > 0; i-- {
			result = append(result, symbol)
		}
	} else if symbol != 0 {
		result = append(result, symbol)
	}

	// fmt.Printf("|%v|", result)
	return string(result), nil
}

/*

	aaaabccddddde
	aaaabccddddde

*/

// 4gg

// 14 100
// 1   1
// 4   10