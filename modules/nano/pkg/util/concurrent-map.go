package util

import "sync"

type Map[K, V any] interface {
	Set(key K, value V)
	Get(key K) V
	Find(key K) (V, bool)
	Delete(key K)
	Keys() []string
	Values() []V
	Len() int
}

type ConcurrentSyncMap[K any, V any] struct {
	nilValue V
	m        sync.Map
}

func (c *ConcurrentSyncMap[K, V]) Set(key K, value V) {
	c.m.Store(key, value)
}

func (c *ConcurrentSyncMap[K, V]) Get(key K) V {
	v, _ := c.Find(key)

	return v
}

func (c *ConcurrentSyncMap[K, V]) Find(key K) (V, bool) {
	v, ok := c.m.Load(key)
	if !ok {
		return c.nilValue, false
	}
	return v.(V), true
}

func (c *ConcurrentSyncMap[K, V]) Delete(key K) {
	c.m.Delete(key)
}

func (c *ConcurrentSyncMap[K, V]) Keys() []string {
	var keys []string

	c.m.Range(func(key, value interface{}) bool {
		keys = append(keys, key.(string))
		return true
	})

	return keys
}

func (c *ConcurrentSyncMap[K, V]) Values() []V {
	var values []V

	c.m.Range(func(key, value interface{}) bool {
		values = append(values, value.(V))
		return true
	})

	return values
}

func (c *ConcurrentSyncMap[K, V]) Len() int {
	result := 0
	c.m.Range(func(key, value interface{}) bool {
		result++
		return true
	})

	return result
}

func NewConcurrentSyncMap[K any, V any]() Map[K, V] {
	return &ConcurrentSyncMap[K, V]{}
}
