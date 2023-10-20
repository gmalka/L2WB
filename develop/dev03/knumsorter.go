package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type knum struct {
	arr         [][]interface{}
	left, right int
}

func (s *knum) Len() int {
	return len(s.arr)
}

func (s *knum) Less(i, j int) bool {
	for t := s.left; t < len(s.arr[i]) && t < s.right; t++ {
		if t >= len(s.arr[j]) {
			return false
		}

		left, ok1 := s.arr[i][t].(int)
		right, ok2 := s.arr[j][t].(int)
		if !ok1 && !ok2 {
			continue
		} else if !ok1 {
			return true
		} else if !ok2 {
			return true 
		} else if left == right {
			continue
		} else {
			return left <= right
		}
	}
	
	return len(s.arr[j]) <= len(s.arr[i])
}

func (s *knum) Swap(i, j int) {
	s.arr[i], s.arr[j] = s.arr[j], s.arr[i]
}

func (s *knum) String() string {
	return fmt.Sprintf("%v", *s)
}

func newKnumsorter(left, right int) *knumsorter {
	return &knumsorter{
		strs: knum{
			left:  left - 1,
			right: right,
			arr:   make([][]interface{}, 0, 10),
		},
	}
}

type knumsorter struct {
	strs knum
}

func (n *knumsorter) add(str string) {
	s := strings.Split(str, " ")

	result := make([]interface{}, len(s))

	for k, v := range s {
		num, err := strconv.Atoi(v)
		if err != nil {
			result[k] = v
		} else {
			result[k] = num
		}
	}

	n.strs.arr = append(n.strs.arr, result)
}

func (n *knumsorter) sort(reverse bool) []string {
	if !reverse {
		sort.Sort(&n.strs)
	} else {
		sort.Sort(sort.Reverse(&n.strs))
	}

	result := make([]string, len(n.strs.arr))
	for k, v := range n.strs.arr {
		b := strings.Builder{}

		for _, val := range v {
			if st, ok := val.(string); ok {
				b.WriteString(st)
				b.WriteByte(' ')
			} else {
				b.WriteString(strconv.Itoa(val.(int)))
				b.WriteByte(' ')
			}
		}

		result[k] = b.String()[:len(b.String())-1]
	}

	return result
}