package main

import (
	"fmt"
	"sort"
	"strconv"
)

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

type numsorter struct {
	nums []int
	strs []string
}

func newNumSorter() *numsorter {
	return &numsorter{
		nums: make([]int, 0, 10),
		strs: make([]string, 0, 10),
	}
}

func (n *numsorter) add(str string) {
	num, err := strconv.Atoi(str)
	fmt.Printf("NUM: %d\n", num)
	if err != nil {
		n.strs = append(n.strs, str)
	} else {
		n.nums = append(n.nums, num)
	}
}

func (n *numsorter) sort(reverse bool) []string {
	result := make([]string, 0, len(n.nums)+len(n.strs))

	if !reverse {
		sort.Ints(n.nums)
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(n.nums)))
	}

	for _, v := range n.nums {
		result = append(result, strconv.Itoa(v))
	}

	result = append(result, n.strs...)

	return result
}