package main

import (
	"github.com/chrumck/CtciExercises/scratchpad"
	"github.com/chrumck/CtciExercises/testlogger"
)

func main() {
	//fmt.Println(scratchpad.NewtonSqrt(2))
	//fmt.Println(math.Sqrt(2))
	//scratchpad.Crawl("http://golang.org/", 10, scratchpad.MakeFakeFetcher())

	testLogger := testlogger.New("debug.log", "debug_prof.log")
	defer testLogger.Close()

	testLogger.Println("Starting...")

	//hashedFiles := scratchpad.HashFilesConcurrent("/media/tomasz/Docs/SoftwareDocs/SoftDesign")
	hashedFiles := scratchpad.HashFilesConcurrent("C:/_private/SoftwareDocs")

	testLogger.WriteHeapProfile()
	testLogger.Printf("Total unique files: %v\n", len(hashedFiles))
	testLogger.Printf("Total duplicate occurrences count: %v\n", scratchpad.DuplicatesCount(hashedFiles))
	// scratchpad.PrintDuplicates(hashedFiles, testLogger)
}
