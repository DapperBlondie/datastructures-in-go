package sets

type Set interface {
	AddElement(data interface{}) (error, bool)
	ContainsElement(data interface{}) bool
	DeleteElement(data interface{}) (error, bool)
}

type Sets struct {
	Values map[interface{}]bool
}
