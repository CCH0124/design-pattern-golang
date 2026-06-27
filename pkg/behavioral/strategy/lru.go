package strategy

// LRU implements the generic EvictionStrategy interface using an LRU (Least Recently Used) algorithm.
type LRU[K comparable, V any] struct {
	keys []K
}

// NewLRU creates and returns a new LRU eviction strategy.
func NewLRU[K comparable, V any]() *LRU[K, V] {
	return &LRU[K, V]{
		keys: make([]K, 0),
	}
}

// OnPut updates the key's position to mark it as most recently used.
func (l *LRU[K, V]) OnPut(key K) {
	l.moveToEnd(key)
}

// OnGet updates the key's position to mark it as most recently used.
func (l *LRU[K, V]) OnGet(key K) {
	l.moveToEnd(key)
}

// OnEvict removes the key from the tracking list.
func (l *LRU[K, V]) OnEvict(key K) {
	for i, k := range l.keys {
		if k == key {
			l.keys = append(l.keys[:i], l.keys[i+1:]...)
			break
		}
	}
}

// Evict returns the least recently used key (which is at index 0).
func (l *LRU[K, V]) Evict(cache *Cache[K, V]) K {
	if len(l.keys) == 0 {
		return *new(K)
	}
	return l.keys[0]
}

// moveToEnd removes the key from its current position and appends it to the end of the slice.
func (l *LRU[K, V]) moveToEnd(key K) {
	for i, k := range l.keys {
		if k == key {
			l.keys = append(l.keys[:i], l.keys[i+1:]...)
			break
		}
	}
	l.keys = append(l.keys, key)
}
