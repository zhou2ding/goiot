package pool

import (
	"errors"
	"goiot/pkg/defs"
	"golang.org/x/sync/semaphore"
	"sync"
)

var (
	once         sync.Once
	poolInstance *Pool
)

type Pool struct {
	w *semaphore.Weighted
}

func GetPool() *Pool {
	once.Do(func() {
		poolInstance = &Pool{semaphore.NewWeighted(defs.PoolMaxSize)}
	})
	return poolInstance
}

func (p *Pool) AddTask(f func()) error {
	if !p.w.TryAcquire(1) {
		return errors.New("pool is full")
	}
	go func() {
		defer p.w.Release(1)
		f()
	}()
	return nil
}
