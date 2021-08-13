package Trees

import (
	"log"
	"reflect"
	"sync"
)

func NewBST() *BST {
	bst := &BST{
		RN: nil,
		L:  &sync.Mutex{},
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
	if bst.RN == nil {
		bst.RN = tn
	} else {

	}
}

func insertTN(rn *TN, tn *TN) {
	dType := reflect.TypeOf(rn.V)
	values := []reflect.Value{reflect.ValueOf(rn.V), reflect.ValueOf(tn.V)}
	method, ok := dType.MethodByName("Equals")

	if !ok {
		log.Fatal("you must implement Equals function for your data type")
		return
	}
	result := method.Func.Call(values)
	if result[0].Int() == 1 {
		if rn.LN == nil {
			rn.LN = tn
		} else {
			insertTN(rn.LN, tn)
		}
	} else {
		if result[0].Int() == -1 {
			if rn.RN == nil {
				rn.RN = tn
			} else {
				insertTN(rn.RN, tn)
			}
		}
	}
}
