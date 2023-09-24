package data

import "sync"

type Issue struct {
    ID string
    Company string `form:"company"`
    Department string `form:"department"`
    Name string `form:"name"`
    PhoneNumber string `form:"phonenumber"`
    InnerNumber string `form:"innernumber"`
    Description string `form:"description"`
    Status string
}

type Subscribers struct {
    mx sync.RWMutex
    m map[string]chan string
}

var Subs = NewSubscribers()

func NewSubscribers() *Subscribers {
    return &Subscribers{
        m: make(map[string]chan string),
    }
}

func (s *Subscribers) Subscribe(key string) chan string {
    s.mx.Lock()
    defer s.mx.Unlock()
    s.m[key] = make(chan string)
    return s.m[key]
}

func (s *Subscribers) Notify(key string) {
    s.mx.Lock()
    defer s.mx.Unlock()
    if c, ok := s.m[key]; ok {
        c <- "notify"
    }
}

func StoreIssue(id string, i *Issue) error {
    D.Store(id, *i)
    Subs.Notify(id)
    return nil
}

func GetIssuesById(id string) (map[string]Issue, error) {
    issues := D.Get(id)
    return issues, nil
}

