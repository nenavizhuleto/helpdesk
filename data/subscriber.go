package data

import (
	"log"
	"sync"
)

type Subscriber struct {
	mx sync.RWMutex
	m  map[string]*SubChannel
}

type SubChannel struct {
	Channel chan bool
	count   int
}

func NewSubChannel() *SubChannel {
	return &SubChannel{
		Channel: make(chan bool),
		count:   0,
	}
}

func (sc *SubChannel) makeSub() {
	sc.count++
}

func (sc *SubChannel) removeSub() {
	sc.count--
}

func NewSubscriber() *Subscriber {
	return &Subscriber{
		m: make(map[string]*SubChannel),
	}
}

func (s *Subscriber) Subscribe(key string) *SubChannel {
	s.mx.Lock()
	defer s.mx.Unlock()
	if c, ok := s.m[key]; ok {
		c.makeSub()
	} else {
		c := NewSubChannel()
		c.makeSub()
		s.m[key] = c
	}
	return s.m[key]
}

func (s *Subscriber) Unsubscribe(key string) {
	s.mx.Lock()
	defer s.mx.Unlock()
	if c, ok := s.m[key]; ok {
		c.removeSub()
	}
}

func (s *Subscriber) Notify(key string) {
	s.mx.Lock()
	defer s.mx.Unlock()
	if c, ok := s.m[key]; ok {
		log.Printf("Sub: %v", c)
		for i := 0; i < c.count; i++ {
			c.Channel <- true
		}
	}
}
