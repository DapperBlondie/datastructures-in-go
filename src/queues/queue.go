package queues

// NewQueue use for creating new Q from scratch
func NewQueue(priority int, data interface{}) *Q {
	order := &Order{
		Priority: priority,
		Data:     data,
	}

	q := Q{}
	q = append(q, order)

	return &q
}

// AddOrder use for adding Order into Q
func (q *Q) AddOrder(order *Order) {
	if len(*q) == 0 {
		*q = append(*q, order)
	} else {
		var appended bool
		appended = false
		var i int
		var addedOrder *Order
		for i, addedOrder = range *q {
			if order.Priority > addedOrder.Priority {
				*q = append((*q)[:i], append(Q{order}, (*q)[i:]...)...)
				appended = true
				break
			}
		}
		if !appended {
			*q = append(*q, order)
		}
	}
}
