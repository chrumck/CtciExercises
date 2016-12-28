package scratchpad

// DuplicatesCount returns the count of duplicate file occurrence.
func DuplicatesCount(hashResults map[string][]string) (count int) {
	for _, filePaths := range hashResults {
		if len(filePaths) > 1 {
			count++
		}
	}
	return
}

// PrintDuplicates prints duplicates to the printer given as parameter.
func PrintDuplicates(hashResults map[string][]string, printer fPrinter) {
	for fileHash, filePaths := range hashResults {
		fileCount := len(filePaths)
		if fileCount < 2 {
			continue
		}
		printer.Printf("Hash:%X - file count: %v\n", fileHash, fileCount)
		for _, filePath := range filePaths {
			printer.Printf("\t%v\n", filePath)
		}
	}
}

type fPrinter interface {
	Printf(format string, v ...interface{})
}
