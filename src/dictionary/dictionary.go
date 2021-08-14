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
