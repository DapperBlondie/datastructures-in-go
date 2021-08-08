package linkedlist

type LinkedList interface {
	AddToHead(data *struct{}) *Node
	IterateAllOver()
	SpecificNode(i uint) (*Node, error)
	LastNode() *Node
	AddToEnd(data *struct{})
	NodeWithData(data *struct{}) (*Node, error)
	AddAfterSpecificNode(i uint, data *struct{}) error
}

// Node use for holding LinkedList data
type Node struct {
	Data *struct{}
	Next *Node
}
