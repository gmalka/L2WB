package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {
	m := *anagram(&[]string{"пятак", "пятка", "тяпка", "столик", "листок", "слиток", "салфетка", "графин", "нифгра", "нифгра", "ниФгра", "Кслито"})
	for k, v := range m {
		fmt.Printf("|%v = %v|\n", k, *v)
	}
}

func anagram(strs *[]string) *map[string]*[]string{
	result := make(map[string]*[]string, 0)
	inuse := make(map[string]string, 0)
	singleword := make(map[string]struct{}, 0)
	for k, v := range *strs {
		(*strs)[k] = strings.ToLower(v)
	}

	sort.Strings(*strs)
	for _, v := range *strs {
		if _, ok := singleword[v]; !ok {
			singleword[v] = struct{}{}
			newv := sortandlow(v)
			if val, ok := inuse[newv]; ok {
				st := result[val]
				*st = append(*st, v)
			} else {
				inuse[newv] = v
				sl := make([]string, 1)
				sl[0] = v
	
				result[v] = &sl
			}
		}
	}
	deletesingleone(&result)

	return &result
}

func deletesingleone(m *map[string]*[]string) {
	for k, v := range *m {
		if len(*v) <= 1 {
			delete(*m, k)
		}
	}
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