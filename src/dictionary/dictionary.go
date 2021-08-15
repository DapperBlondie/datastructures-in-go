package dictionary

import (
	"github.com/rs/zerolog/log"
	"sync"
)

func NewDict() *Dict {
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
