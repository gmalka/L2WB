package main

import "fmt"

type Sorter interface {
	Sort([]int)[]int
}

type MergerSort struct {

}

func (m MergerSort)Sort(mas []int) []int {
	fmt.Println("Sorting by merge sort")
	return mas
}

type QuickSort struct {

}

func (m QuickSort)Sort(mas []int) []int {
	fmt.Println("Sorting by quick sort")
	return mas
}

type SortedStore struct {
	s Sorter
	mas []int
}

func (st *SortedStore) Add(num int) {
	st.mas = append(st.mas, num)
	st.mas = st.s.Sort(st.mas)
}

func (st *SortedStore) SetSort(s Sorter) {
	st.s = s
}

func NewStore(s Sorter) *SortedStore {
	st := &SortedStore{
		s: s,
		mas: make([]int, 0),
	}

	return st
}