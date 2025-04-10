package cache

import (
	"testing"

	"github.com/scottkgregory/parsley/internal/assert"
	"github.com/scottkgregory/parsley/internal/nodes"
)

func TestCache(t *testing.T) {
	c, err := NewCache()
	assert.Nil(t, err)

	node := nodes.NewNumberNode(12)

	c.Set("a", node)

	actual, present := c.Get("a")
	assert.Equal(t, node, actual)
	assert.Equal(t, true, present)

	actual, present = c.Get("b")
	assert.Equal(t, nil, actual)
	assert.Equal(t, false, present)

	c.Close()
}

func TestNoOpCacheGet(t *testing.T) {
	node, present := NewNoOpCache().Get("key")
	assert.Equal(t, false, present)
	assert.Nil(t, node)
}

func TestNoOpCacheSet(t *testing.T) {
	NewNoOpCache().Set("key", nodes.NewNumberNode(12))
}

func TestNoOpCacheClose(t *testing.T) {
	NewNoOpCache().Close()
}
