package queues

// Q our queue type
type Q []*Order

// Order a structure for holding data
type Order struct {
	Priority int
	Data     interface{}
}

type Queue interface {
	AddOrder(order *Order)
}
