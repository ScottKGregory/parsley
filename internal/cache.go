package cache

import (
	"github.com/dgraph-io/ristretto/v2"
	"github.com/scottkgregory/parsley/internal/parser/nodes"
)

type Store[K comparable, V any] interface {
	Wait()
	Get(key K) (V, bool)
	Set(key K, value V, cost int64) bool
	Close()
}

func NewRistrettoCache() (Store[string, nodes.Node], error) {
	return ristretto.NewCache(&ristretto.Config[string, nodes.Node]{
		NumCounters: 1e7,
		MaxCost:     1 << 30,
		BufferItems: 64,
	})
}

type NoOpCache[K comparable, V any] struct{}

func NewNoOpCache() Store[string, nodes.Node] {
	return &NoOpCache[string, nodes.Node]{}
}

func (c *NoOpCache[K, V]) Wait() {}

func (c *NoOpCache[K, V]) Get(key K) (V, bool) {
	return *new(V), false
}

func (c *NoOpCache[K, V]) Set(key K, value V, cost int64) bool {
	return true
}

func (c *NoOpCache[K, V]) Close() {}
