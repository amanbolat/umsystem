package pkg

import (
	"sync"
)

type Map[T any] struct {
	internalMap sync.Map
}

func NewMap[T any]() *Map[T] {
	return &Map[T]{}
}

func (m *Map[T]) Set(key string, val T) {
	m.internalMap.Store(key, val)
}

func (m *Map[T]) Delete(key string) {
	m.internalMap.Delete(key)
}

func (m *Map[T]) Exists(key string) bool {
	_, ok := m.internalMap.Load(key)
	return ok
}

func (m *Map[T]) Get(key string) (T, bool) {
	val, ok := m.internalMap.Load(key)
	if !ok {
		return *new(T), false
	}

	return val.(T), true
}
