package scratchpad

import (
	"crypto/sha256"
	"io/ioutil"
	"path/filepath"
)

// HashFiles hashes all files and finds file duplicates.
func HashFiles(dirPath string) (hashResults map[string][]string) {
	var debugger = NewDebugger("debug_hashfiles.log", "debug_hashfiles_prof.log")
	defer debugger.Close()

	hashResults = make(map[string][]string)
	hashFiles(dirPath, hashResults)

	debugger.WriteHeapProfile()
	return
}

func hashFiles(dirPath string, hashResults map[string][]string) {
	fileInfos, readDirError := ioutil.ReadDir(dirPath)
	if readDirError != nil {
		return
	}

	for i := range fileInfos {
		filePath := filepath.Join(dirPath, fileInfos[i].Name())
		if fileInfos[i].IsDir() {
			hashFiles(filePath, hashResults)
			continue
		}
		hashFile(filePath, hashResults)
	}
}

func hashFile(filePath string, hashResults map[string][]string) {
	fileBytes, fileReadError := ioutil.ReadFile(filePath)
	if fileReadError != nil {
		return
	}

	checksum := sha256.Sum256(fileBytes)
	hash := string(checksum[:])
	hashPaths := hashResults[hash]
	hashResults[hash] = append(hashPaths, filePath)
}
