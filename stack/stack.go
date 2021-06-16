package stack

type Stack struct {
	data []interface{}
}

func New() *Stack {
	return &Stack{
		data: []interface{}{},
	}
}

func (s *Stack) Push(v interface{}) {
	s.data = append(s.data, v)
}

func (s *Stack) Peek() interface{} {
	return s.data[len(s.data)-1]
}

func (s *Stack) Pop() {
	if len(s.data) < 1 {
		return
	}

	s.data = s.data[:len(s.data)-1]
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}
