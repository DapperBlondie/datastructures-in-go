package Trees

import "sync"

type TN struct {
	K  int32
	V  interface{}
	LN *TN
	RN *TN
}

type BST struct {
	RootN *TN
	L     *sync.Mutex
}

type Trees interface {
	InsertElement(k int32, v interface{})
	InOrderTraversTree(f func(v interface{}))
	PreOrderTraverseTree(f func(v interface{}))
	PostOrderTraverseTree(f func(v interface{}))
	MinTN() interface{}
	MaxTN() interface{}
	SearchTN(v interface{}) bool
}
