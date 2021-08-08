package stacks

//Element class
type Element struct {
	Value interface{}
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	elements []*Element
	elementCount int
}
