package practice

type Stack struct {
	size int
	data [10]int
}

func (s *Stack) push(num int) {
	s.data[s.size] = num
	s.size++
}

func (s *Stack) pop() int {
	s.size--
	return s.data[s.size]
}
