package scratchpad

import (
	"sync"
)

// ConcurrentMap allows access to a map in a thread safe way.
type ConcurrentMap struct {
	mp  map[string]int
	mux sync.Mutex
}

// Add item to concurrent map in a safe way.
func (cMap *ConcurrentMap) Add(item string) {
	defer cMap.mux.Unlock()
	cMap.mux.Lock()
	cMap.mp[item]++
}

// Exists Checks if item exists in the concurrent map.
func (cMap *ConcurrentMap) Exists(item string) bool {
	defer cMap.mux.Unlock()
	cMap.mux.Lock()
	var _, ok = cMap.mp[item]
	return ok
}