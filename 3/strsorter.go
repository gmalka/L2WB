package main

import (
	"sort"
)

type strsorter struct {
	strs []string
}

func newStrSorter() *numsorter {
	return &numsorter{
		nums: make([]int, 0, 10),
		strs: make([]string, 0, 10),
	}
}

func (n *strsorter) add(str string) {
	n.strs = append(n.strs, str)
}

func (n *strsorter) sort(reverse bool) []string {
	if !reverse {
		sort.Strings(n.strs)
	} else {
		sort.Sort(sort.Reverse(sort.StringSlice(n.strs)))
	}

	return n.strs
}
