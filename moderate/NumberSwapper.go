package moderate

// SwapNumber swaps number without a temporary variable.
func SwapNumber(one *int, two *int){
   *one = *one + *two
   *two = *one - *two
   *one = *one - *two
}
