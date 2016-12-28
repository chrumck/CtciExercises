package scratchpad

import (
	"crypto/sha256"
	"io/ioutil"
	"path/filepath"
)

// HashFilesClosure hashes all files and finds file duplicates.
func HashFilesClosure(dirPath string) (hashResults map[string][]string) {
	hashResults = make(map[string][]string)

	var hashFilesClosure func(dirPath string)
	var hashFileClosure func(filePath string)

	hashFilesClosure = func(dirPath string) {
		fileInfos, readDirError := ioutil.ReadDir(dirPath)
		if readDirError != nil {
			return
		}

		for i := range fileInfos {
			filePath := filepath.Join(dirPath, fileInfos[i].Name())
			if fileInfos[i].IsDir() {
				hashFilesClosure(filePath)
				continue
			}
			hashFileClosure(filePath)
		}
	}

	hashFileClosure = func(filePath string) {
		fileBytes, fileReadError := ioutil.ReadFile(filePath)
		if fileReadError != nil {
			return
		}

		checksum := sha256.Sum256(fileBytes)
		hash := string(checksum[:])
		hashPaths := hashResults[hash]
		hashResults[hash] = append(hashPaths, filePath)
	}

	hashFilesClosure(dirPath)

	return
}
