package test

import (
	"designpattern/pkg/behavioral/strategy"
	"testing"
)

func TestCacheEvictionFIFO(t *testing.T) {
	fifo := strategy.NewFIFO[string, string]()
	cache, err := strategy.NewCache[string, string](fifo, 3)
	if err != nil {
		t.Fatalf("failed to create cache: %v", err)
	}

	cache.Add("k1", "v1")
	cache.Add("k2", "v2")
	cache.Add("k3", "v3")

	// Capacity is reached. Adding "k4" should evict the oldest "k1".
	evicted := cache.Add("k4", "v4")
	if evicted != "k1" {
		t.Errorf("expected k1 to be evicted, got %s", evicted)
	}

	if _, ok := cache.Get("k1"); ok {
		t.Errorf("k1 should have been evicted")
	}
}

func TestCacheEvictionLRU(t *testing.T) {
	lru := strategy.NewLRU[string, string]()
	cache, err := strategy.NewCache[string, string](lru, 3)
	if err != nil {
		t.Fatalf("failed to create cache: %v", err)
	}

	cache.Add("k1", "v1")
	cache.Add("k2", "v2")
	cache.Add("k3", "v3")

	// Access k1, making it the most recently used
	cache.Get("k1")

	// Capacity is reached. Adding "k4" should evict "k2" (the least recently used)
	evicted := cache.Add("k4", "v4")
	if evicted != "k2" {
		t.Errorf("expected k2 to be evicted, got %s", evicted)
	}

	if _, ok := cache.Get("k2"); ok {
		t.Errorf("k2 should have been evicted")
	}
}

func TestCacheEvictionLFU(t *testing.T) {
	lfu := strategy.NewLFU[string, string]()
	cache, err := strategy.NewCache[string, string](lfu, 3)
	if err != nil {
		t.Fatalf("failed to create cache: %v", err)
	}

	cache.Add("k1", "v1") // freq(k1) = 1
	cache.Add("k2", "v2") // freq(k2) = 1
	cache.Add("k3", "v3") // freq(k3) = 1

	// Access k2 and k3 to increase their frequencies
	cache.Get("k2") // freq(k2) = 2
	cache.Get("k3") // freq(k3) = 2
	cache.Get("k3") // freq(k3) = 3

	// Capacity is reached. Adding "k4" should evict "k1" (frequency 1)
	evicted := cache.Add("k4", "v4")
	if evicted != "k1" {
		t.Errorf("expected k1 to be evicted, got %s", evicted)
	}

	if _, ok := cache.Get("k1"); ok {
		t.Errorf("k1 should have been evicted")
	}
}

func TestDynamicStrategySwitching(t *testing.T) {
	fifo := strategy.NewFIFO[string, string]()
	cache, _ := strategy.NewCache[string, string](fifo, 3)

	cache.Add("k1", "v1")
	cache.Add("k2", "v2")
	cache.Add("k3", "v3")

	// Switch strategy dynamically to LRU
	lru := strategy.NewLRU[string, string]()
	// Sync existing keys to the new strategy for correctness
	cache.SetEvictionStrategy(lru)
	lru.OnPut("k1")
	lru.OnPut("k2")
	lru.OnPut("k3")

	// Access k1 (LRU tracking: k2 -> k3 -> k1)
	cache.Get("k1")

	// Adding k4 should evict k2 (least recently used)
	evicted := cache.Add("k4", "v4")
	if evicted != "k2" {
		t.Errorf("expected k2 to be evicted, got %s", evicted)
	}
}


