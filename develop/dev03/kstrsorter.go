package main

import (
	"fmt"
	"sort"
	"strings"
)

type kstring struct {
	arr         [][]string
	left, right int
}

func (s *kstring) Len() int {
	return len(s.arr)
}

func (s *kstring) Less(i, j int) bool {
	var (
		l, r string
	)
	if len(s.arr[i]) > s.left {
		if len(s.arr[i]) < s.right {
			l = strings.Join(s.arr[i][s.left:len(s.arr[i])], " ")
		} else {
			l = strings.Join(s.arr[i][s.left:s.right], " ")
		}
	} else {
		l = ""
	}

	if len(s.arr[j]) > s.left {
		if len(s.arr[j]) < s.right {
			r = strings.Join(s.arr[j][s.left:len(s.arr[j])], " ")
		} else {
			r = strings.Join(s.arr[j][s.left:s.right], " ")
		}
	} else {
		r = ""
	}

	return l <= r
	// for t := s.left; t < len(s.arr[i]) && t < s.right; t++ {
	// 	if t > len(s.arr[j]) {
	// 		return true
	// 	}

	// 	if s.arr[i][t] == s.arr[j][t] {
	// 		continue
	// 	} else {
	// 		return s.arr[i][t] > s.arr[j][t]
	// 	}
	// }

	// return true
}

func (s *kstring) Swap(i, j int) {
	s.arr[i], s.arr[j] = s.arr[j], s.arr[i]
}

func (s *kstring) String() string {
	return fmt.Sprintf("%v", *s)
}

func newkstrsorter(left, right int) *kstrsorter {
	return &kstrsorter{
		strs: kstring{
			left:  left - 1,
			right: right,
			arr:   make([][]string, 0, 10),
		},
	}
}

type kstrsorter struct {
	strs kstring
}

func (n *kstrsorter) add(str string) {
	s := strings.Split(str, " ")

	n.strs.arr = append(n.strs.arr, s)
}

func (n *kstrsorter) sort(reverse bool) []string {
	if !reverse {
		sort.Sort(&n.strs)
	} else {
		sort.Sort(sort.Reverse(&n.strs))
	}

	result := make([]string, len(n.strs.arr))
	for k, v := range n.strs.arr {
		str := strings.Join(v, " ")

		result[k] = str
	}

	return result
}
