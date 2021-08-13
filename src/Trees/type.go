package Trees

import "sync"

type TN struct {
	K  int32
	V  interface{}
	LN *TN
	RN *TN
}

type BST struct {
	RN *TN
	L  *sync.Mutex
}
