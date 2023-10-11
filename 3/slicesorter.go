package main

type ssorter struct {
	reverse bool
	arr     sorter
}

func newSSorter(s sorter, reverse bool) *ssorter {
	return &ssorter{
		arr:     s,
		reverse: reverse,
	}
}

func (s *ssorter) add(str string) {
	s.arr.add(str)
}

func (s *ssorter) getsorted() []string {
	return s.arr.sort(s.reverse)
}
