package Easy


type Stack struct{
	data []byte
}

func (s *Stack)push(e byte) {
	s.data = append(s.data, e)
}

func (s *Stack)pop() byte {
	if len(s.data) > 0 {
		e := s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		return e
	}
	return 0
}

func (s *Stack)size() int {
	return len(s.data)
}

func isValid(s string) bool {
	if len(s) == 0 {
		return false
	}

	stack := new(Stack)
	for _, c := range s {
		if c=='{' || c=='(' || c=='[' {
			stack.push(byte(c))
		}else{
			e := stack.pop()
			diff := byte(c) - e
			if diff != 1 && diff != 2  {
				return false
			}
		}
	}

	return stack.size() == 0
}
