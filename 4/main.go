package main

import (
	"fmt"
	"sort"
	"unicode"
)

func main() {
	fmt.Println(anagram(&[]string{"пятак", "пятка", "тяпка", "столик", "листок", "слиток"}))
}

func anagram(strs *[]string) *map[string]*[]string{
	result := make(map[string]*[]string, 0)
	inuse := make(map[string]string, 0)

	sort.Strings(*strs)
	for _, v := range *strs {
		newv := sortandlow(v)
		if val, ok := inuse[newv]; ok {
			inuse[newv] = v
			st := result[newv]
			*st = append(*st, v)
		} else {
			inuse[newv] = v
			sl := make([]string, 1)
			sl[0] = v

			result[newv] = &sl
		}
	}

	return &result
}

func sortandlow(str string) string {
	r := []rune(str)

	for k, v := range r {
		r[k] = unicode.ToLower(v)
	}

	sort.Slice(r, func(i, j int) bool{
		return r[i] < r[j]
	})

	return string(r)
}