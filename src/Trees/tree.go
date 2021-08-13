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

// InOrderTraversTree use for traversing tree Left,Root,Right
func (bst *BST) InOrderTraversTree(f func(v interface{})) {
	bst.L.Lock()
	defer bst.L.Unlock()
	inOrderTraverseTree(bst.RootN, f)
}

// inOrderTraverseTree a helper function for BST.InOrderTraversTree
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

// PreOrderTraverseTree use for traversing tree Root,Left,Right
func (bst *BST) PreOrderTraverseTree(f func(v interface{})) {
	bst.L.Lock()
	defer bst.L.Unlock()
	preOrderTraverseTree(bst.RootN, f)
}

// preOrderTraverseTree a helper function for BST.PreOrderTraverseTree
func preOrderTraverseTree(rootN *TN, f func(v interface{})) {
	if rootN != nil {
		f(rootN.V)
		inOrderTraverseTree(rootN.LN, f)
		inOrderTraverseTree(rootN.RN, f)
	} else if rootN == nil {
		f(rootN.V)
	}
	return
}

// PostOrderTraverseTree use for traversing tree Root,Left,Right
func (bst *BST) PostOrderTraverseTree(f func(v interface{})) {
	bst.L.Lock()
	defer bst.L.Unlock()
	postOrderTraverseTree(bst.RootN, f)
}

// postOrderTraverseTree a helper function for BST.PostOrderTraverseTree
func postOrderTraverseTree(rootN *TN, f func(v interface{})) {
	if rootN != nil {
		inOrderTraverseTree(rootN.LN, f)
		inOrderTraverseTree(rootN.RN, f)
		f(rootN.V)
	} else if rootN == nil {
		f(rootN.V)
	}
	return
}

// MinTN use for finding the deepest LeftNode in BST
func (bst *BST) MinTN() interface{} {
	bst.L.Lock()
	defer bst.L.Unlock()

	tn := bst.RootN
	if tn == nil {
		return tn.V
	}
	for {
		if tn.LN == nil {
			return tn.V
		}
		tn = tn.LN
	}
}

// MaxTN use for finding the deepest RightNode in BST
func (bst *BST) MaxTN() interface{} {
	bst.L.Lock()
	defer bst.L.Unlock()

	tn := bst.RootN
	if tn == nil {
		return tn.V
	}
	for {
		if tn.RN == nil {
			return tn.RN
		}
		tn = tn.RN
	}
}
