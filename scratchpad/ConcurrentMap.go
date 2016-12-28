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
	cMap.mux.Lock()
	defer cMap.mux.Unlock()
	cMap.mp[item]++
}

// Exists Checks if item exists in the concurrent map.
func (cMap *ConcurrentMap) Exists(item string) bool {
	cMap.mux.Lock()
	defer cMap.mux.Unlock()
	var _, ok = cMap.mp[item]
	return ok
}