package slices

type Stack struct {
	Data []int
}

func (s *Stack) Push(v int) {
	s.Data = append(s.Data, v)
}

func (s *Stack) Len() int {
	return len(s.Data)
}

func (s *Stack) Pop() (int, bool) {

	if len(s.Data) == 0 {
		return 0, false
	}
	val := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return val, true
}
