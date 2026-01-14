package lazy_lock

import (
	"sync"
	"sync/atomic"
)

type LazyLock[T any] struct {
	once  sync.Once
	value atomic.Pointer[T]
	init  func() T
}

func New[T any](init func() T) *LazyLock[T] {
	return &LazyLock[T]{init: init}
}

func (l *LazyLock[T]) Get() T {
	l.once.Do(func() {
		val := l.init()
		l.value.Store(&val)
	})
	return *l.value.Load()
}
