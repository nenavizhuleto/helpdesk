package data

import (
	"sync"
)

type Db struct {
	mx sync.RWMutex
	m  map[string]map[string]Issue
}

func GetDb() map[string]map[string]Issue {
	D.mx.Lock()
	defer D.mx.Unlock()

	cp := make(map[string]map[string]Issue)
	for k, v := range D.m {
		for k1, v1 := range v {
			cp[k] = make(map[string]Issue)
			cp[k][k1] = v1
		}
	}

	return cp
}

func NewDb() *Db {
	return &Db{
		m: make(map[string]map[string]Issue),
	}
}

func (d *Db) Store(key string, value Issue) {
	d.mx.Lock()
	defer d.mx.Unlock()
	if _, ok := d.m[key]; !ok {
		d.m[key] = make(map[string]Issue)
	}
	d.m[key][value.ID] = value
}

func (d *Db) Get(key string) map[string]Issue {
	d.mx.Lock()
	defer d.mx.Unlock()
	return d.m[key]
}

var D = NewDb()
