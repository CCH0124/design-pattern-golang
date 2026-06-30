package prototype

import (
	"fmt"
	"sync"
)

type PrototypeRegistry[T Cloneable[T]] struct {
	mu        sync.RWMutex
	templates map[string]T
}

func NewPrototypeRegistry[T Cloneable[T]]() *PrototypeRegistry[T] {
	return &PrototypeRegistry[T]{
		templates: make(map[string]T),
	}
}

func (pr *PrototypeRegistry[T]) Register(name string, tpl T) {
	pr.mu.Lock()
	defer pr.mu.Unlock()
	pr.templates[name] = tpl
}
func (r *PrototypeRegistry[T]) Get(name string) (T, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tpl, exists := r.templates[name]
	if !exists {
		var zero T
		return zero, fmt.Errorf("template '%s' not found", name)
	}

	return tpl.Clone(), nil
}
