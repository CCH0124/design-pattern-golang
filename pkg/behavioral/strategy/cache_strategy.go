package strategy

import "errors"

// EvictionStrategy defines the interface for cache eviction algorithms.
// K is the comparable type of key, V is the type of value.
type EvictionStrategy[K comparable, V any] interface {
	// OnPut is called when a key is added/updated in the cache.
	OnPut(key K)
	// OnGet is called when a key is accessed.
	OnGet(key K)
	// OnEvict is called when a key is removed from the cache.
	OnEvict(key K)
	// Evict selects and returns a key to be evicted from the cache.
	Evict(cache *Cache[K, V]) K
}

// Cache is the context struct that delegates eviction behavior to EvictionStrategy.
type Cache[K comparable, V any] struct {
	storage      map[K]V
	evictionAlgo EvictionStrategy[K, V]
	capacity     int
	maxCapacity  int
}

// NewCache initializes and returns a generic Cache context.
func NewCache[K comparable, V any](algo EvictionStrategy[K, V], maxCapacity int) (*Cache[K, V], error) {
	if maxCapacity <= 0 {
		return nil, errors.New("max capacity must be greater than 0")
	}
	return &Cache[K, V]{
		storage:      make(map[K]V),
		evictionAlgo: algo,
		capacity:     0,
		maxCapacity:  maxCapacity,
	}, nil
}

// SetEvictionStrategy dynamically switches the eviction strategy at runtime.
func (c *Cache[K, V]) SetEvictionStrategy(algo EvictionStrategy[K, V]) {
	c.evictionAlgo = algo
}

// Add inserts a key-value pair into the cache, triggering eviction if capacity is reached.
// It returns the evicted key (if any).
func (c *Cache[K, V]) Add(key K, value V) K {
	var evictedKey K
	if c.capacity >= c.maxCapacity {
		evictedKey = c.evict()
	}

	c.storage[key] = value
	c.capacity++
	if c.evictionAlgo != nil {
		c.evictionAlgo.OnPut(key)
	}
	return evictedKey
}

// Get retrieves a value from the cache.
func (c *Cache[K, V]) Get(key K) (V, bool) {
	val, ok := c.storage[key]
	if ok && c.evictionAlgo != nil {
		c.evictionAlgo.OnGet(key)
	}
	return val, ok
}

// Keys returns all keys currently in the cache storage.
func (c *Cache[K, V]) Keys() map[K]bool {
	keys := make(map[K]bool)
	for k := range c.storage {
		keys[k] = true
	}
	return keys
}

// evict triggers the eviction process using the configured strategy.
func (c *Cache[K, V]) evict() K {
	if c.evictionAlgo == nil {
		return *new(K)
	}
	keyToEvict := c.evictionAlgo.Evict(c)
	delete(c.storage, keyToEvict)
	c.capacity--
	c.evictionAlgo.OnEvict(keyToEvict)
	return keyToEvict
}
