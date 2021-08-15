package dictionary

import (
	"github.com/rs/zerolog/log"
	"reflect"
	"sync"
)

func NewDict(v Value) *Dict {
	method, ok := reflect.TypeOf(v).MethodByName("Equals")
	if !ok {
		log.Error().Msg("you must implement the equals method")
		return nil
	}
	METHOD = method
	dict := &Dict{
		E: make(map[Key]Value),
		M: &sync.RWMutex{},
	}

	return dict
}

func (d *Dict) Put(k Key, v Value) {
	d.M.Lock()
	defer d.M.Unlock()
	if d.E == nil {
		log.Error().Msg("your dictionary is not initialized yet")
		return
	}

	d.E[k] = v
	return
}

func (d *Dict) Remove(k Key) bool {
	d.M.Lock()
	defer d.M.Unlock()
	_, exists := d.E[k]
	if exists {
		delete(d.E, k)
		return true
	}

	return false
}

func (d *Dict) Contains(k Key) bool {
	d.M.RLock()
	defer d.M.RUnlock()
	_, exists := d.E[k]

	return exists
}

func (d *Dict) FindByKey(k Key) Value {
	d.M.RLock()
	defer d.M.RUnlock()
	return d.E[k]
}

func (d *Dict) Reset() {
	d.M.Lock()
	defer d.M.Unlock()
	d.E = make(map[Key]Value)
	return
}

func (d *Dict) NumberOfElements() int {
	d.M.RLock()
	defer d.M.RUnlock()
	return len(d.E)
}

func (d *Dict) GetKeys() []Key {
	d.M.RLock()
	defer d.M.RUnlock()
	var dictKeys []Key = []Key{}
	for K := range d.E {
		dictKeys = append(dictKeys, K)
	}

	return dictKeys
}

func (d *Dict) GetValues() []Value {
	d.M.RLock()
	defer d.M.RUnlock()
	var dictValues []Value = []Value{}
	for _, V := range d.E {
		dictValues = append(dictValues, V)
	}

	return dictValues
}

func (d *Dict) GetKeyByValue(v Value) Key {
	var values []reflect.Value = []reflect.Value{reflect.ValueOf(v), reflect.ValueOf(v)}
	for K, V := range d.E {
		values[1] = reflect.ValueOf(V)
		result := METHOD.Func.Call(values)
		if result[0].Bool() {
			return K
		}
	}

	return false
}
