package main

import (
	"bufio"
	"fmt"
	"io"
)

type ssorter struct {
	r       *bufio.Reader
	reverse bool
	arr     sortstore
}

func newSSorter(r *bufio.Reader, s sortstore, reverse bool) *ssorter {
	return &ssorter{
		r:       r,
		arr:     s,
		reverse: reverse,
	}
}

func (s *ssorter) getsorted() ([]string, error) {
	for {
		var (
			result []byte
			l      []byte
			ok     bool
			err    error
		)

		for l, ok, err = s.r.ReadLine(); ok && err == nil; l, ok, err = s.r.ReadLine() {
			result = append(result, l...)
		}
		result = append(result, l...)

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("cant read line: %v", err)
		}

		s.arr.add(string(result))
	}
	return s.arr.sort(s.reverse), nil
}
