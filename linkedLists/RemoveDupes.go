package linkedLists

import (
	"container/list"
)

// RemoveDupes removes duplicates from container/list.List
// uses map[container/list.Element]*list.Element to store the unique elements and returns a filtered
// list.List constructed from those.
func RemoveDupes(l *list.List) (filteredList *list.List) {
	foundElements := make(map[interface{}]*list.Element)
	for {
		element := l.Front()
		if element == nil {
			break
		}
		foundElements[element.Value] = element
		l.Remove(element)
	}
	filteredList = list.New()
	for _, element := range foundElements {
		filteredList.PushFront(element)
	}
	return
}

// RemoveDupes2 removes duplicates from container/list.List
// Brute force - compare each element to each other - running time O(n^2/2)
func RemoveDupes2(l *list.List) {
	i := l.Front()
	for i != nil {
		j := i.Next()
		for j != nil {
			if i.Value == j.Value {
				nextJ := j.Next()
				l.Remove(j)
				j = nextJ
				continue
			}
			j = j.Next()
		}
		i = i.Next()
	}
}
