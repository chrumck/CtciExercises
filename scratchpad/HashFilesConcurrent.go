package scratchpad

import (
	"crypto/sha256"
	"io/ioutil"
	"path/filepath"
	"sync"
)

// HashFilesConcurrent hashes all files and finds file duplicates.
// Recursion is done with goroutines.
func HashFilesConcurrent(dirPath string) (hashResults map[string][]string) {
	hashResultsConcurrent := &HashFilesConcurrentMap{Map: make(map[string][]string)}
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	go hashFilesConcurrent(dirPath, hashResultsConcurrent, &waitGroup)
	waitGroup.Wait()

	return hashResultsConcurrent.Map
}

func hashFilesConcurrent(dirPath string, hashResults *HashFilesConcurrentMap, waitGroup *sync.WaitGroup) {
	fileInfos, readDirError := ioutil.ReadDir(dirPath)
	if readDirError != nil {
		waitGroup.Done()
		return
	}
	for i := range fileInfos {
		filePath := filepath.Join(dirPath, fileInfos[i].Name())
		if fileInfos[i].IsDir() {
			waitGroup.Add(1)
			go hashFilesConcurrent(filePath, hashResults, waitGroup)
			continue
		}
		hashFileConcurrent(filePath, hashResults)
	}
	waitGroup.Done()
}

func hashFileConcurrent(filePath string, hashResults *HashFilesConcurrentMap) {
	fileBytes, fileReadError := ioutil.ReadFile(filePath)
	if fileReadError != nil {
		return
	}
	checksum := sha256.Sum256(fileBytes)
	hash := string(checksum[:])
	hashResults.Add(hash, filePath)
}
