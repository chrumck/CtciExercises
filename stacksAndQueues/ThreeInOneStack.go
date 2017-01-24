package stacksAndQueues

import (
	"fmt"
)

const threeInOneStackCount = 3

// ThreeInOneStack implement three stacks held in one array.
type ThreeInOneStack struct {
	stackSlice []interface{}
	itemCount  [3]int
}

// Push adds an item to the given stack.
func (tio *ThreeInOneStack) Push(stackNumber int, value interface{}) {
	checkStackNumber(stackNumber)

	itemIndex := tio.itemCount[stackNumber]*threeInOneStackCount + stackNumber
	tio.growSlice(itemIndex)

	tio.stackSlice[itemIndex] = value
	tio.itemCount[stackNumber]++
}

// Pop removes the item from the stack and returns its value.
func (tio *ThreeInOneStack) Pop(stackNumber int) (value interface{}) {
	checkStackNumber(stackNumber)
	tio.checkIfStackEmpty(stackNumber)

	tio.itemCount[stackNumber]--
	itemIndex := tio.itemCount[stackNumber]*threeInOneStackCount + stackNumber

	value = tio.stackSlice[itemIndex]
	tio.shrinkSlice()
	return
}

// TODO: implement IsEmpty(stackNumber int)

func checkStackNumber(stackNumber int) {
	if stackNumber >= threeInOneStackCount || stackNumber < 0 {
		panic("The stack number has to be between 0 and 2")
	}
}

func (tio *ThreeInOneStack) growSlice(itemIndex int) {
	if len(tio.stackSlice) > itemIndex {
		return
	}
	newSize := itemIndex*2 + 1
	newStack := make([]interface{}, newSize, newSize)
	copy(newStack, tio.stackSlice)
	tio.stackSlice = newStack
}

func (tio *ThreeInOneStack) checkIfStackEmpty(stackNumber int) {
	if tio.itemCount[stackNumber] <= 0 {
		panic(fmt.Sprintf("Stack number %v is empty", stackNumber))
	}
}

func (tio *ThreeInOneStack) shrinkSlice() {
	// TODO: finish implementing. Need to check all three stack item Counts before shrinking.
	// index0 := (tio.itemCount[0]-1)*threeInOneStackCount + 0
	// index1 := (tio.itemCount[1]-1)*threeInOneStackCount + 1
	// index2 := (tio.itemCount[2]-1)*threeInOneStackCount + 2
}
