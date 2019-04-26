package immutable

type Stack struct {
	top    interface{}
	bottom *Stack
}

func (s *Stack) IsEmpty() bool {
	return s.bottom == nil
}

func (s *Stack) Peek() interface{} {
	return s.top
}

func (s *Stack) Pop() *Stack {
	return s.bottom
}

func (s *Stack) Push(value interface{}) *Stack {
	return &Stack{
		top:    value,
		bottom: s,
	}
}
