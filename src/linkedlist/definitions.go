package linkedlist

import "fmt"

type LinkedList interface {
	AddToHead(data struct{}) *Node
	IterateAllOver()
	SpecificNode(i uint) *Node
	LastNode() *Node
	AddToEnd(data struct{})
	NodeWithData(data struct{})
	AddAfterSpecificNode(i int, data struct{})

}

type Node struct {
	Data struct{}
	Next *Node
}

// CreateHead use for creating the first node for LinkedList
func CreateHead(data struct{}) *Node {
	hNode := &Node{
		Data: data,
		Next: nil,
	}
	return hNode
}

// AddToHead use for adding a node to head of LinkedList
func (h *Node) AddToHead(data struct{}) *Node {
	n := &Node{
		Data: data,
		Next: nil,
	}
	n.Next = h

	return n
}

func (h *Node) IterateAllOver() {
	for n := h.Next ;n != nil;n = n.Next {
		fmt.Println(n)
	}
}

func (h *Node) SpecificNode(i uint) *Node {
	var n *Node
	for n = h.Next ;i != 0;n = n.Next {
		i -= 1
	}

	return n
}

func (h *Node) LastNode() *Node {
	var n *Node
	for n = h.Next ;  ;n = n.Next {
		if n.Next == nil {
			return n
		}
	}
}

func (h *Node) AddToEnd(data struct{}) {
	ln := h.LastNode()
	n := &Node{
		Data: data,
		Next: nil,
	}

	ln.Next = n
}

func (h *Node) NodeWithData(data struct{}) {}

func (h *Node) AddAfterSpecificNode(i int, data struct{}) {}