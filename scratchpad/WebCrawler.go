package scratchpad

import (
	"fmt"
	"sync"
)

// Crawl uses fetcher to recursively crawl pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	cMap := &ConcurrentMap{mp: make(map[string]int)}
	crawl(url, depth, fetcher, &waitGroup, cMap)
	waitGroup.Wait()
}

func crawl(
	url string,
	depth int,
	fetcher Fetcher,
	waitGroup *sync.WaitGroup,
	cMap *ConcurrentMap) {

	defer waitGroup.Done()

	if depth <= 0 || cMap.Exists(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	cMap.Add(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		waitGroup.Add(1)
		go crawl(u, depth-1, fetcher, waitGroup, cMap)
	}

	return
}
