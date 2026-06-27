package strategy

// FIFO implements the generic EvictionStrategy interface using a FIFO (First-In, First-Out) algorithm.
type FIFO[K comparable, V any] struct {
	keys []K
}

// NewFIFO creates and returns a new FIFO eviction strategy.
func NewFIFO[K comparable, V any]() *FIFO[K, V] {
	return &FIFO[K, V]{
		keys: make([]K, 0),
	}
}

// OnPut tracks the insertion order of keys.
func (f *FIFO[K, V]) OnPut(key K) {
	for _, k := range f.keys {
		if k == key {
			return // Already tracked
		}
	}
	f.keys = append(f.keys, key)
}

// OnGet does nothing for FIFO strategy.
func (f *FIFO[K, V]) OnGet(key K) {}

// OnEvict removes the key from the tracking list.
func (f *FIFO[K, V]) OnEvict(key K) {
	for i, k := range f.keys {
		if k == key {
			f.keys = append(f.keys[:i], f.keys[i+1:]...)
			break
		}
	}
}

// Evict returns the oldest key in the cache.
func (f *FIFO[K, V]) Evict(cache *Cache[K, V]) K {
	if len(f.keys) == 0 {
		return *new(K)
	}
	return f.keys[0]
}
