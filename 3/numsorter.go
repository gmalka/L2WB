package main

import (
	"sort"
	"strconv"
)

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
