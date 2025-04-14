// Package cache is used to caches parsed expressions ready to be re-evaluated with different data
package cache

import (
	"fmt"

	"github.com/dgraph-io/ristretto/v2"
	"github.com/scottkgregory/parsley/internal/helpers"
	"github.com/scottkgregory/parsley/internal/nodes"
)

// ErrCacheSetup is returned when setting up the cache fails
const ErrCacheSetup = helpers.ConstError("error setting up cache")

// Store provides caching functions
type Store[K string, V any] interface {
	Get(key K) (V, bool)
	Set(key K, value V)
	Close()
}

// NewCache returns a new cache backed by github.com/dgraph-io/ristretto/v2
func NewCache() (Store[string, nodes.Node], error) {
	inner, err := ristretto.NewCache(&ristretto.Config[string, nodes.Node]{
		NumCounters: 1e7,
		MaxCost:     1 << 30,
		BufferItems: 64,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCacheSetup, err)
	}

	return &cache[string, nodes.Node]{inner}, nil
}

type cache[K string, V any] struct {
	inner *ristretto.Cache[K, V]
}

// Get gets a value from the cache, if available
func (c *cache[K, V]) Get(key K) (V, bool) {
	return c.inner.Get(key)
}

// Set sets a value in the cache
func (c *cache[K, V]) Set(key K, value V) {
	c.inner.Set(key, value, 0)
	c.inner.Wait()
}

// Close releases any underlying resources
func (c *cache[K, V]) Close() {
	c.inner.Close()
}

// NewNoOpCache returns an implementation of the cache interface which doesn't do any caching
func NewNoOpCache() Store[string, nodes.Node] {
	return &noOpCache[string, nodes.Node]{}
}

type noOpCache[K string, V any] struct{}

// Get gets a value from the cache, if available
func (c *noOpCache[K, V]) Get(_ K) (V, bool) {
	return *new(V), false
}

// Set sets a value in the cache
func (c *noOpCache[K, V]) Set(_ K, _ V) {}

// Close releases any underlying resources
func (c *noOpCache[K, V]) Close() {}
