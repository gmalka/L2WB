package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type mnsorter struct {
	r       *bufio.Reader
	m       map[string]struct{}
	arr     sortstore
	left    int
	reverse bool
}

func newMNsorter(r *bufio.Reader, s sortstore, reverse bool, left int) *mnsorter {
	return &mnsorter{
		arr:     s,
		r:       r,
		left:    left,
		m:       make(map[string]struct{}, 10),
		reverse: reverse,
	}
}

func (s *mnsorter) getsorted() ([]string, error) {
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
		bb := bytes.Split(result, []byte{' '})

		if s.left < len(bb) {
			if checkfornonnum(bb[s.left]) {
				if _, ok = s.m[""]; !ok {
					s.m[""] = struct{}{}

					s.arr.add(string(result))
				}
			} else if _, ok = s.m[string(bb[s.left])]; !ok {
				s.m[string(bb[s.left])] = struct{}{}

				s.arr.add(string(result))
			}
		}
	}

	return s.arr.sort(s.reverse), nil
}

func checkfornonnum(b []byte) bool {
	for _, v := range b {
		if v < '0' || v > '9' {
			return true
		}
	}

	return false
}