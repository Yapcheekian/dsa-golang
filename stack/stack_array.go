package stack

type StackArray struct {
	arr []string
}

func NewStackArray() *StackArray {
	return &StackArray{
		arr: make([]string, 0, 5),
	}
}

func (s *StackArray) Push(item string) {
	s.arr = append(s.arr, item)
}

func (s *StackArray) Pop() string {
	if s.IsEmpty() {
		return ""
	}
	item := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return item
}

func (s *StackArray) IsEmpty() bool {
	return len(s.arr) == 0
}
