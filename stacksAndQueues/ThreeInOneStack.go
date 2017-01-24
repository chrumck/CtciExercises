package stacksAndQueues

import (
	"fmt"
)

const stackCount = 3

// ThreeInOneStack implement three stacks held in one array.
type ThreeInOneStack struct {
	stackSlice []interface{}
	itemCount  [stackCount]int
}

// Push adds an item to the given stack.
func (tio *ThreeInOneStack) Push(stackNumber int, value interface{}) {
	panicIfStackNumberWrong(stackNumber)

	itemIndex := tio.itemCount[stackNumber]*stackCount + stackNumber
	tio.growSlice(itemIndex)
	tio.stackSlice[itemIndex] = value
	tio.itemCount[stackNumber]++
}

// Pop removes the item from the stack and returns its value.
func (tio *ThreeInOneStack) Pop(stackNumber int) (value interface{}) {
	panicIfStackNumberWrong(stackNumber)
	tio.panicIfStackEmpty(stackNumber)

	tio.itemCount[stackNumber]--
	itemIndex := tio.itemCount[stackNumber]*stackCount + stackNumber
	value = tio.stackSlice[itemIndex]
	tio.shrinkSlice()
	return
}

// TODO: implement IsEmpty(stackNumber int)

func panicIfStackNumberWrong(stackNumber int) {
	if stackNumber >= stackCount || stackNumber < 0 {
		panic("The stack number has to be between 0 and 2")
	}
}

func (tio *ThreeInOneStack) panicIfStackEmpty(stackNumber int) {
	if tio.itemCount[stackNumber] <= 0 {
		panic(fmt.Sprintf("Stack number %v is empty", stackNumber))
	}
}

func (tio *ThreeInOneStack) growSlice(itemIndex int) {
	if len(tio.stackSlice) <= itemIndex {
		tio.resizeSlice(itemIndex*2 + 1)
	}
}

func (tio *ThreeInOneStack) shrinkSlice() {
	maxItemIndex := tio.getMaxItemIndex()
	if len(tio.stackSlice) >= maxItemIndex * 4 {
		tio.resizeSlice(len(tio.stackSlice) / 2)
	}
}

func (tio *ThreeInOneStack) resizeSlice(newSize int) {
	newStack := make([]interface{}, newSize, newSize)
	copy(newStack, tio.stackSlice)
	tio.stackSlice = newStack
}

func (tio *ThreeInOneStack) getMaxItemIndex() (maxItemIndex int) {
	for stackNumber := 0; stackNumber <= 2; stackNumber++ {
		itemIndex := (tio.itemCount[stackNumber]-1)*stackCount + stackNumber
		if maxItemIndex < itemIndex {
			maxItemIndex = itemIndex
		}
	}
	return
}
