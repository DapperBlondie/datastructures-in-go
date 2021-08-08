package sets

import (
	"errors"
	zerolog "github.com/rs/zerolog/log"
	"reflect"
)

func NewSet() *Sets {
	set := &Sets{Values: make(map[interface{}]bool)}

	return set
}

// AddElement use for adding element, but you should implement the Equals function for comparison between structures
func (set *Sets) AddElement(data interface{}) (error, bool) {
	td := reflect.TypeOf(data)
	if td.Kind() == reflect.PtrTo(reflect.TypeOf(reflect.Struct)).Kind() {
		method, ok := td.MethodByName("Equals")
		if !ok {
			return errors.New("you are not implement Equals method for comparison"), false
		}
		for key, _ := range set.Values {
			values := []reflect.Value{reflect.ValueOf(data), reflect.ValueOf(key)}
			result := method.Func.Call(values)
			if result[0].Bool() {
				return errors.New("this data is exists in set"), false
			}
		}

		set.Values[data] = true
		return nil, true
	}

	if td.Kind() == reflect.Struct {
		method, ok := td.MethodByName("Equals")
		if !ok {
			zerolog.Error().Msg("your structure does not have method for checking equality")
			return errors.New("your structure does not have method for checking equality"), false
		}
		for key, _ := range set.Values {
			values := []reflect.Value{reflect.ValueOf(key), reflect.ValueOf(data)}
			result := method.Func.Call(values)
			if result[0].Bool() {
				return errors.New("your set have this data"), false
			}
		}
		set.Values[data] = true
		return nil, true
	} else {
		exists := set.ContainsElement(data)
		if exists {
			return errors.New("this data exists in your set"), false
		}

		set.Values[data] = true
		return nil, true
	}
}

// ContainsElement use for checking a data is exists in your set or not just use it any types of Integers, floats or string
func (set *Sets) ContainsElement(data interface{}) bool {
	_, exists := set.Values[data]

	return exists
}

// DeleteElement for deleting a element from set
func (set *Sets) DeleteElement(data interface{}) (error, bool) {
	td := reflect.TypeOf(data)
	if td.Kind() == reflect.PtrTo(reflect.TypeOf(reflect.Struct)).Kind() {
		method, ok := td.MethodByName("Equals")
		if !ok {
			return errors.New("you are not implement Equals method for comparison"), false
		}
		for key, _ := range set.Values {
			values := []reflect.Value{reflect.ValueOf(data), reflect.ValueOf(key)}
			result := method.Func.Call(values)
			if result[0].Bool() {
				delete(set.Values, data)
				return nil, true
			}
		}
	} else if td.Kind() == reflect.Struct {
		method, ok := td.MethodByName("Equals")
		if !ok {
			return errors.New("you are not implement Equals method for comparison"), false
		}
		for key, _ := range set.Values {
			values := []reflect.Value{reflect.ValueOf(data), reflect.ValueOf(key)}
			result := method.Func.Call(values)
			if result[0].Bool() {
				delete(set.Values, data)
				return nil, true
			}
		}
	}

	delete(set.Values, data)
	return nil, true
}
