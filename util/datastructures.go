package util

type runeStack struct {
	contents []rune
	index    int
}

func NewRuneStack() *runeStack {
	s := new(runeStack)
	s.contents = make([]rune, 0, 20)
	s.index = -1
	return s
}

func (s *runeStack) Push(r rune) {
	s.index++
	if len(s.contents) == s.index {
		s.contents = append(s.contents, r)
	} else {
		s.contents[s.index] = r
	}
}

func (s *runeStack) Pop() rune {
	r := s.contents[s.index]
	s.index--
	return r
}

func (s *runeStack) Peek() rune {
	if s.index < 0 {
		return 0
	}
	return s.contents[s.index]
}

func (s *runeStack) IsEmpty() bool {
	return s.index < 0
}

func (s *runeStack) Reset() {
	s.index = -1
}
