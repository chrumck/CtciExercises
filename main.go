package main

import "github.com/chrumck/CtciExercises/ScratchPad"

func main() {
	//fmt.Println(ScratchPad.NewtonSqrt(2))
	//fmt.Println(math.Sqrt(2))

	ScratchPad.Crawl("http://golang.org/", 10, ScratchPad.MakeFakeFetcher())
}
