package main

import "github.com/chrumck/CtciExercises/scratchpad"
import "fmt"

func main() {
	//fmt.Println(scratchpad.NewtonSqrt(2))
	//fmt.Println(math.Sqrt(2))
	//scratchpad.Crawl("http://golang.org/", 10, scratchpad.MakeFakeFetcher())

	var debugger = scratchpad.NewDebugger("debug.log", "debugProf.log")
	defer debugger.Close()

	debugger.Log("Starting...\n")

	var hashedFiles = scratchpad.HashFiles("/media/tomasz/Docs/SoftwareDocs")

	debugger.WriteHeapProfile()

	debugger.Log(fmt.Sprintf("Total unique files: %v\n", len(hashedFiles)))

	duplicateCount := 0
	for _, filePaths := range hashedFiles {
		if len(filePaths) > 1 {
			duplicateCount++
		}
	}

	debugger.Log(fmt.Sprintf("Total duplicate occurrences count: %v\n", duplicateCount))

	for fileHash, filePaths := range hashedFiles {
		if len(filePaths) < 2 {
			continue
		}
		debugger.Log(fmt.Sprintf("Hash:%X - file count: %v\n", fileHash, len(filePaths)))
		for _, filePath := range filePaths {
			debugger.Log(fmt.Sprintln("\t" + filePath))
		}
	}
}
