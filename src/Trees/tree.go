package Trees

import (
	"log"
	"reflect"
	"sync"
)

func NewBST() *BST {
	bst := &BST{
		RootN: nil,
		L:     &sync.Mutex{},
	}

	return bst
}

// InsertElement for inserting a elem int our BST
func (bst *BST) InsertElement(k int32, v interface{}) {
	bst.L.Lock()
	defer bst.L.Unlock()

	tn := &TN{
		K:  k,
		V:  v,
		LN: nil,
		RN: nil,
	}
	if bst.RootN == nil {
		bst.RootN = tn
	} else {
		insertTN(bst.RootN, tn)
	}
}

// insertTN a helper function for adding node to BST
func insertTN(rootN *TN, tn *TN) {
	dType := reflect.TypeOf(rootN.V)
	values := []reflect.Value{reflect.ValueOf(rootN.V), reflect.ValueOf(tn.V)}
	method, ok := dType.MethodByName("Equals")

	if !ok {
		log.Fatal("you must implement Equals function for your data type")
		return
	}
	result := method.Func.Call(values)
	if result[0].Int() == 1 {
		if rootN.LN == nil {
			rootN.LN = tn
		} else {
			insertTN(rootN.LN, tn)
		}
	} else {
		if result[0].Int() == -1 {
			if rootN.RN == nil {
				rootN.RN = tn
			} else {
				insertTN(rootN.RN, tn)
			}
		}
	}
}

func (bst *BST) InOrderTraversTree(f func(v interface{})) {
	bst.L.Lock()
	defer bst.L.Unlock()
	inOrderTraverseTree(bst.RootN, f)
}

func inOrderTraverseTree(rootN *TN, f func(v interface{})) {
	if rootN != nil {
		inOrderTraverseTree(rootN.LN, f)
		f(rootN.V)
		inOrderTraverseTree(rootN.RN, f)
	} else if rootN == nil {
		f(rootN.V)
	}

	return
}
