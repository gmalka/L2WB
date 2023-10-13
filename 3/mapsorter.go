package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type msorter struct {
	r       *bufio.Reader
	m       map[string]struct{}
	arr     sortstore
	borders []int
	reverse bool
}

func newMsorter(r *bufio.Reader, s sortstore, reverse bool, borders []int) *msorter {
	borders[0] -= 1
	return &msorter{
		arr:     s,
		r:       r,
		borders: borders,
		m:       make(map[string]struct{}, 10),
		reverse: reverse,
	}
}

func (s *msorter) getsorted() ([]string, error) {
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
		b := append([]int{}, s.borders...)

		if len(b) == 1 {
			b = append(b, len(bb))
		} else if b[1] > len(bb) {
			b[1] = len(bb)
		}

		if s.borders[0] < len(bb) {
			sl := make([]string, 0, b[1] - b[0])
			for i := b[0]; i < b[1]; i++ {
				sl = append(sl, string(bb[i]))
			}

			j := strings.Join([]string(sl), " ")
			if _, ok = s.m[j]; !ok {
				s.m[j] = struct{}{}

				s.arr.add(string(result))
			}
		} else if _, ok = s.m["\x01"]; !ok {
			s.m["\x01"] = struct{}{}

			s.arr.add(string(result))
		}
	}

	return s.arr.sort(s.reverse), nil
}