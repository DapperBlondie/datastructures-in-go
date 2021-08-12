package stacks

// NewStack returns a new stack.
func NewStack() *S {
	s := &S{
		Elements:     []*Element{},
		ElementCount: 0,
	}

	return s
}

// String method on Element class
func (element *Element) String() string {

	return ""
}

// Push use for push Element into S
func (s *S) Push(elem *Element) {
	s.Elements = append(s.Elements, elem)
	s.ElementCount += 1
}

// Pop use for pop Element from S
func (s *S) Pop() *Element {
	if s.ElementCount == 0 {
		return nil
	}
	var length int = len(s.Elements)
	var element *Element = s.Elements[length-1]

	if length > 1 {
		s.Elements = s.Elements[:length-1]
	} else {
		s.Elements = s.Elements[0:]
	}
	s.ElementCount = len(s.Elements)

	return element
}
