package cache

import (
	"github.com/dgraph-io/ristretto/v2"
	"github.com/scottkgregory/parsley/internal/parser/nodes"
)

type Store[K string, V any] interface {
	Get(key K) (V, bool)
	Set(key K, value V)
	Close()
}

func NewCache() (Store[string, nodes.Node], error) {
	inner, err := ristretto.NewCache(&ristretto.Config[string, nodes.Node]{
		NumCounters: 1e7,
		MaxCost:     1 << 30,
		BufferItems: 64,
	})
	return &cache[string, nodes.Node]{inner}, err
}

type cache[K string, V any] struct {
	inner *ristretto.Cache[K, V]
}

func (c *cache[K, V]) Get(key K) (V, bool) {
	return c.inner.Get(key)
}

func (c *cache[K, V]) Set(key K, value V) {
	c.inner.Set(key, value, 0)
	c.inner.Wait()
}

func (c *cache[K, V]) Close() {
	c.inner.Close()
}

func NewNoOpCache() Store[string, nodes.Node] {
	return &noOpCache[string, nodes.Node]{}
}

type noOpCache[K string, V any] struct{}

func (c *noOpCache[K, V]) Get(key K) (V, bool) {
	return *new(V), false
}

func (c *noOpCache[K, V]) Set(key K, value V) {}

func (c *noOpCache[K, V]) Close() {}
