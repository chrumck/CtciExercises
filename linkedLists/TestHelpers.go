package linkedLists

import (
	"container/list"
)

// SampleT is a sample struct to be used for testing.
type SampleT struct {
	sampleInt    int
	sampleString string
}

// GetExampleLists returns a slice of example *list.List items
func GetExampleLists() (exampleLists []*list.List) {
	l := list.New()
	l.PushFront(SampleT{3, "foo"})
	l.PushFront(SampleT{4, "bar"})
	l.PushFront(SampleT{5, "baz"})
	l.PushFront(SampleT{5, "bar"})
	l.PushFront(SampleT{4, "bar"})
	l.PushFront(SampleT{3, "foo"})
	exampleLists = append(exampleLists, l)

	l = list.New()
	l.PushFront(SampleT{4, "bar"})
	l.PushFront(SampleT{5, "baz"})
	l.PushFront(SampleT{5, "bar"})
	l.PushFront(SampleT{4, "foo"})
	exampleLists = append(exampleLists, l)

	l = list.New()
	l.PushFront(SampleT{4, "foo"})
	l.PushFront(SampleT{4, "foo"})
	exampleLists = append(exampleLists, l)

	l = list.New()
	l.PushFront(SampleT{4, "foo"})
	l.PushFront(SampleT{4, "bar"})
	exampleLists = append(exampleLists, l)

	l = list.New()
	l.PushFront(SampleT{4, "foo"})
	exampleLists = append(exampleLists, l)

	l = list.New()
	exampleLists = append(exampleLists, l)

	return
}

// ContainsDupes checks if there are duplicates in the given list.List
func ContainsDupes(l *list.List) bool {
	foundValues := map[interface{}]bool{}
	for {
		element := l.Front()
		if element == nil {
			break
		}
		elementValue := l.Remove(element)
		if foundValues[elementValue] == true {
			return true
		}
		foundValues[elementValue] = true
	}
	return false
}
