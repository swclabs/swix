package tools

import (
	"sync"
)

type ITasks interface {
	Run()
}

type Pool struct {
	channel chan interface{}
}

type Handler func(<-chan interface{}, *sync.WaitGroup)

type Tasks struct {
	handler    Handler
	concurrent int
	wg         *sync.WaitGroup
	Process    *Pool
}

func NewPool() *Pool {
	return &Pool{
		channel: make(chan interface{}),
	}
}

func (sp *Pool) Process(model interface{}) {
	// _data, _ := json.Marshal(model)
	sp.channel <- model
}

func (sp *Pool) Close() {
	close(sp.channel)
}

func NewTask(pool *Pool, handler Handler, concurrent int) *Tasks {
	var _wg sync.WaitGroup
	_wg.Add(concurrent)
	return &Tasks{
		Process:    pool,
		handler:    handler,
		concurrent: concurrent,
		wg:         &_wg,
	}
}

func (t *Tasks) Run() {
	for i := 0; i < t.concurrent; i++ {
		go t.handler(t.Process.channel, t.wg)
	}
}
