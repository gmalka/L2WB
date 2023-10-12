package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// func newNumSorter() *numsorter {
// 	return &numsorter{
// 		strs: num{
// 			arr: make([][]interface{}, 0, 10),
// 		},
// 	}
// }

// type numsorter struct {
// 	strs num
// }

// func (n *numsorter) add(str string) {
// 	s := strings.Split(str, " ")

// 	result := make([]interface{}, len(s))

// 	for k, v := range s {
// 		num, err := strconv.Atoi(v)
// 		if err != nil {
// 			result[k] = v
// 		} else {
// 			result[k] = num
// 		}
// 	}

// 	n.strs.arr = append(n.strs.arr, result)
// }

// func (n *numsorter) sort(reverse bool) []string {
// 	if !reverse {
// 		sort.Sort(&n.strs)
// 	} else {
// 		sort.Sort(sort.Reverse(&n.strs))
// 	}

// 	result := make([]string, len(n.strs.arr))
// 	for k, v := range n.strs.arr {
// 		b := strings.Builder{}

// 		for _, val := range v {
// 			if st, ok := val.(string); ok {
// 				b.WriteString(st)
// 				b.WriteByte(' ')
// 			} else {
// 				b.WriteString(strconv.Itoa(val.(int)))
// 				b.WriteByte(' ')
// 			}
// 		}

// 		result[k] = b.String()[:len(b.String())-1]
// 	}

// 	return result
// }

type num struct {
	arr [][]interface{}
}

func (s *num) Len() int {
	return len(s.arr)
}

func (s *num) Less(i, j int) bool {
	for t := 0; t < len(s.arr[i]); t++ {
		if t > len(s.arr[j]) {
			return true
		}

		left, ok1 := s.arr[i][t].(int)
		right, ok2 := s.arr[j][t].(int)
		if !ok1 && !ok2 {
			continue
		} else if !ok1 {
			return false
		} else if !ok2 {
			return true
		} else if left == right {
			continue
		} else {
			return left < right
		}
	}

	return len(s.arr[j]) <= len(s.arr[i])
}

func (s *num) Swap(i, j int) {
	s.arr[i], s.arr[j] = s.arr[j], s.arr[i]
}

func (s *num) String() string {
	return fmt.Sprintf("%v", *s)
}

// func (s *num) add(num int, str string) {
// 	s.arr = append(s.arr, num)
// 	s.strs = append(s.strs, str)
// }

// func (s *num) get() []string {
// 	return s.strs
// }

// func (s *num) Len() int {
// 	return len(s.arr)
// }

// func (s *num) Less(i, j int) bool {
// 	return s.arr[i] > s.arr[j]
// }

// func (s *num) Swap(i, j int) {
// 	s.arr[i], s.arr[j] = s.arr[j], s.arr[i]
// 	s.strs[i], s.strs[j] = s.strs[j], s.strs[i]
// }

// func (s *num) String() string {
// 	return fmt.Sprintf("%v", s.strs)
// }

type numsorter struct {
	nums num
}

func newNumSorter() *numsorter {
	return &numsorter{
		nums: num{
			arr: make([][]interface{}, 0, 10),
		},
	}
}

func (n *numsorter) add(str string) {
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

	n.nums.arr = append(n.nums.arr, result)
}

func (n *numsorter) sort(reverse bool) []string {
	if !reverse {
		sort.Sort(&n.nums)
	} else {
		sort.Sort(sort.Reverse(&n.nums))
	}

	result := make([]string, len(n.nums.arr))
	for k, v := range n.nums.arr {
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
