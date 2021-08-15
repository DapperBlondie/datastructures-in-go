package dictionary

import "sync"

// Key is a type that designed to hold any type of integer, string, float32 and float64 as a key in Dict
type Key interface{}

// Value is a type for holding your complex objects that a pointer to a struct
type Value *struct{}

type Dict struct {
	E map[Key]Value
	M *sync.RWMutex
}

type Dictionary interface {
	Put(k Key, v Value)
	Remove(k Key) bool
	Contains(k Key) bool
	Reset()
	NumberOfElements() int
	GetKeys() []Key
	GetValues() []Value
}
