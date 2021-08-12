package stacks

//Element class
type Element struct {
	Value interface{}
}

// S is a basic LIFO stack that resizes as needed.
type S struct {
	Elements     []*Element
	ElementCount int
}

type Stack interface {
	Push(element *Element)
	Pop() *Element
}
