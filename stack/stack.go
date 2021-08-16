package stack

type Stack struct {
	top   int
	max   int
	array []float32
}

func NewStack(size int) *Stack {
	return &Stack{
		top:   -1,
		max:   size,
		array: make([]float32, size),
	}
}

func (s *Stack) Push(n float32) {
	if s.top == s.max-1 {
		return
	}

	s.top++
	s.array[s.top] = n
}

func (s *Stack) Pop() {
	if s.top == -1 {
		return
	}

	s.top--
}

func (s *Stack) Peek() float32 {
	return s.array[s.top]
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) IsFull() bool {
	return s.top == s.max-1
}
