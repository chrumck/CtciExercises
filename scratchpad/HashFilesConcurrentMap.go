package scratchpad

import (
	"sync"
)

// HashFilesConcurrentMap allows access to a map in a thread safe way.
type HashFilesConcurrentMap struct {
	Map map[string][]string
	mux sync.Mutex
}

// Add filePath to concurrent map in a safe way.
func (cMap *HashFilesConcurrentMap) Add(hash, filePath string) {
	cMap.mux.Lock()
	defer cMap.mux.Unlock()

	hashPaths := cMap.Map[hash]
	cMap.Map[hash] = append(hashPaths, filePath)
}
