package dictionary

import "sync"

type Key interface{}

type Value *struct{}

type Dict struct {
	E map[Key]Value
	M *sync.RWMutex
}

type Dictionary interface {
	Put(k Key, v Value)
}
