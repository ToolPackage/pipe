package color

import "fmt"

type Stack struct {
	data []interface{}
	len  int
}

func NewStack() *Stack {
	s := &Stack{}
	s.data = make([]interface{}, 0)
	s.len = 0

	return s
}

func (s *Stack) Len() int {
	return s.len
}

func (s *Stack) isEmpty() bool {
	return s.len == 0
}

func (s *Stack) Pop() (el interface{}) {
	if s.len <= 0 {
		panic("try to pop from empty stack")
	}
	s.len--
	el, s.data = s.data[s.len], s.data[:s.len]
	return
}

func (s *Stack) Push(el interface{}) {
	s.data = append(s.data, el)
	s.len++
}

func (s *Stack) Peek() interface{} {
	if s.len <= 0 {
		panic("try to peek from empty stack")
	}
	return s.data[s.len-1]
}

func (s *Stack) AtTop(idx int) interface{} {
	if idx >= s.len {
		panic("index out of bound")
	}
	return s.data[s.len-idx-1]
}

func (s *Stack) String() string {
	return fmt.Sprintf("Stack%v", s.data)
}
