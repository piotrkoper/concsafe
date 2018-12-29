// Package concsafe provides wrapper for concurrency-safe variables
package concsafe

import (
	"fmt"
	"strings"
	"sync"
)

// Slice is guarded by RWMutex
type Slice struct {
	sync.RWMutex // Guards items
	items        []interface{}
}

// Count of items
func (e *Slice) Count() int {
	e.RLock()
	defer e.RUnlock()
	return len(e.items)
}

// Get returns all items
func (e *Slice) Get() []interface{} {
	e.RLock()
	defer e.RUnlock()
	return e.items
}

// Add to slice
func (e *Slice) Add(item interface{}) {
	e.Lock()
	defer e.Unlock()
	e.items = append(e.items, item)
}

// List of errors
func (e *Slice) List() (string, error) {
	e.RLock()
	defer e.RUnlock()

	if e.Count() == 0 {
		return "", nil
	}
	// Use string builder
	var sb strings.Builder
	for i, e := range e.items {
		_, err := sb.WriteString(fmt.Sprintf("\n%6d | %s", i, e))
		if err != nil {
			return "", err
		}
	}
	return sb.String(), nil
}
