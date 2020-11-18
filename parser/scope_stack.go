package parser

import "fmt"

type Scope int

const (
	ScopePipe = iota
	ScopeVariableNode
	ScopeStreamNode
	ScopeStreamSplitter
	ScopeStreamCollector
	ScopePipeNodeArray
)

type ScopeStack struct {
	data []Scope
	len  int
}

func (s *ScopeStack) Len() int {
	return s.len
}

func (s *ScopeStack) isEmpty() bool {
	return s.len == 0
}

func (s *ScopeStack) Pop() (el Scope) {
	if s.len <= 0 {
		panic("try to pop from empty stack")
	}
	s.len--
	el, s.data = s.data[s.len], s.data[:s.len]
	return
}

func (s *ScopeStack) Push(el Scope) {
	s.data = append(s.data, el)
	s.len++
}

func (s *ScopeStack) Peek() Scope {
	if s.len <= 0 {
		panic("try to peek from empty stack")
	}
	return s.data[s.len-1]
}

func (s *ScopeStack) AtTop(idx int) Scope {
	if idx >= s.len {
		panic("index out of bound")
	}
	return s.data[s.len-idx-1]
}

func (s *ScopeStack) String() string {
	return fmt.Sprintf("ScopeStack%v", s.data)
}
