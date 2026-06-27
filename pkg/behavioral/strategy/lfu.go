package strategy

// LFU implements the generic EvictionStrategy interface using an LFU (Least Frequently Used) algorithm.
type LFU[K comparable, V any] struct {
	frequencies map[K]int
}

// NewLFU creates and returns a new LFU eviction strategy.
func NewLFU[K comparable, V any]() *LFU[K, V] {
	return &LFU[K, V]{
		frequencies: make(map[K]int),
	}
}

// OnPut increments the access frequency of the key.
func (l *LFU[K, V]) OnPut(key K) {
	l.frequencies[key]++
}

// OnGet increments the access frequency of the key.
func (l *LFU[K, V]) OnGet(key K) {
	l.frequencies[key]++
}

// OnEvict deletes the frequency information of the key.
func (l *LFU[K, V]) OnEvict(key K) {
	delete(l.frequencies, key)
}

// Evict returns the key with the lowest frequency that is currently present in the cache.
func (l *LFU[K, V]) Evict(cache *Cache[K, V]) K {
	var minKey K
	minFreq := -1

	// Only evict keys currently present in the cache's storage.
	for k := range cache.storage {
		freq := l.frequencies[k]
		if minFreq == -1 || freq < minFreq {
			minFreq = freq
			minKey = k
		}
	}
	return minKey
}
