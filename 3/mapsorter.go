package main

type msorter struct {
	m   map[string]struct{}
	arr sorter
	reverse bool
}

func newmsorter(s sorter, reverse bool) *msorter {
	return &msorter{
		arr: s,
		m:   make(map[string]struct{}, 10),
		reverse: reverse,
	}
}

func (s *msorter) add(str string) {
	if _, ok := s.m[str]; !ok {
		s.add(str)
		s.m[str] = struct{}{}
	}
}

func (s *msorter) getsorted() []string {
	return s.arr.sort(s.reverse)
}
