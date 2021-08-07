package linkedlist

import (
	"errors"
	"fmt"
	"reflect"
)

type LinkedList interface {
	AddToHead(data *struct{}) *Node
	IterateAllOver()
	SpecificNode(i uint) (*Node, error)
	LastNode() *Node
	AddToEnd(data *struct{})
	NodeWithData(data *struct{}) (*Node, error)
	AddAfterSpecificNode(i uint, data *struct{}) error
}

type Node struct {
	Data *struct{}
	Next *Node
}

// CreateHead use for creating the first node for LinkedList
func CreateHead(data *struct{}) *Node {
	hNode := &Node{
		Data: data,
		Next: nil,
	}
	return hNode
}

// AddToHead use for adding a node to head of LinkedList
func (h *Node) AddToHead(data *struct{}) *Node {
	n := &Node{
		Data: data,
		Next: nil,
	}
	n.Next = h

	return n
}

// IterateAllOver use for iterating over all of existing nodes
func (h *Node) IterateAllOver() {
	for n := h.Next; n != nil; n = n.Next {
		fmt.Println(n)
	}
}

// SpecificNode use for finding a specific node by its position
func (h *Node) SpecificNode(i uint) (*Node, error) {
	var n *Node
	for n = h.Next; i != 0; n = n.Next {
		if n.Next == nil {
			return n, errors.New("i is so bigger than length of linked list")
		}

		i -= 1
	}

	return n, nil
}

// LastNode use for returning the last node of linked list
func (h *Node) LastNode() *Node {
	var n *Node
	for n = h.Next; ; n = n.Next {
		if n.Next == nil {
			return n
		}
	}
}

// AddToEnd use for adding a node to the end of the linked list
func (h *Node) AddToEnd(data *struct{}) {
	ln := h.LastNode()
	n := &Node{
		Data: data,
		Next: nil,
	}

	ln.Next = n
}

// NodeWithData use for finding a specific node by comparing the all associated fields with its Equals function
func (h *Node) NodeWithData(data *struct{}) (*Node, error) {
	method, ok := reflect.TypeOf(data).MethodByName("Equals")
	if !ok {
		return nil, errors.New("your data structure does not implement the Equals function")
	}

	for n := h.Next; n != nil; n = n.Next {
		values := []reflect.Value{reflect.ValueOf(data).Elem(), reflect.ValueOf(n.Data).Elem()}
		result := method.Func.Call(values) // It takes the first as the receiver of function

		if result[0].Bool() {
			return n, nil
		}
	}

	return nil, errors.New("could not be able to find node")
}

// AddAfterSpecificNode use for adding a node after specific node
func (h *Node) AddAfterSpecificNode(i uint, data *struct{}) error {
	n, err := h.SpecificNode(i)
	var nn *Node

	if err != nil {
		nn = &Node{
			Data: data,
			Next: nil,
		}

		n.Next = nn
		return err
	}

	nn = &Node{
		Data: data,
		Next: nil,
	}

	nn.Next = n.Next
	n.Next = nn

	return nil
}
