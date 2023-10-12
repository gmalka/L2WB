package main

import (
	"bufio"
	"fmt"
	"io"
)

type msorter struct {
	r       *bufio.Reader
	m       map[string]struct{}
	arr     sortstore
	reverse bool
}

func newMsorter(r *bufio.Reader, s sortstore, reverse bool) *msorter {
	return &msorter{
		arr:     s,
		r:       r,
		m:       make(map[string]struct{}, 10),
		reverse: reverse,
	}
}

func (s *msorter) getsorted() ([]string, error) {
	for {
		var (
			result []byte
			l []byte
			ok bool
			err error
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

		if _, ok = s.m[string(result)]; !ok {
			s.m[string(result)] = struct{}{}

			s.arr.add(string(result))
		}
	}
	return s.arr.sort(s.reverse), nil
}