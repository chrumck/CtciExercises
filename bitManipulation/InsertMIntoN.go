package bitManipulation

import "fmt"

// InsertMIntoN binary inserts M into N at position j through i, where M and N are int32.
func InsertMIntoN(M, N int32, i, j int) int32 {

   // printouts for debugging purposes
	fmt.Printf("M before:  %b\n", M)
	fmt.Printf("N before:  %b\n", N)

	bitClearMask := getBitClearMask(i, j)
   N = N & int32(bitClearMask)
	M = M << uint(i)

   // printouts for debugging purposes
   fmt.Printf("Mask:      %b\n", ^bitClearMask)
	fmt.Printf("M after:   %b\n", M)
	fmt.Printf("N after:   %b\n", N)
	fmt.Printf("M|N after: %b\n", M|N)
	fmt.Printf("M|N dec:   %v\n", M|N)

	return M | N
}

func getBitClearMask(i, j int) (clearMask int32) {
   clearMask = ^0
	for index := 0; index <= j; index++ {
		clearMask = clearMask << 1
		if index > i {
			clearMask++
		}
	}
	return
}
