package data

import "sync"

type Subscriber struct {
    mx sync.RWMutex
    m map[string]chan bool
}

func NewSubscriber() *Subscriber {
    return &Subscriber{
        m: make(map[string]chan bool),
    }
}

func (s *Subscriber) Subscribe(key string) chan bool {
    s.mx.Lock()
    defer s.mx.Unlock()
    s.m[key] = make(chan bool)
    return s.m[key]
}

func (s *Subscriber) Notify(key string) {
    s.mx.Lock()
    defer s.mx.Unlock()
    if c, ok := s.m[key]; ok {
        c <- true
    }
}
