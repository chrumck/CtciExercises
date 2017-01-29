package mathAndLogicPuzzles

// FindPrimes finds primes in the set between 2 and n.
// User Arestothenes sieve implemented in concurrent pipeline-like fashion.
func FindPrimes(n int) (primes []int) {
	if n < 2 {
		return
	}

	ints := generate(n)
	for {
		prime, open := <-ints
		if !open {
			break
		}
		primes = append(primes, prime)
		ints = filter(ints, prime)
	}

	return
}

func generate(n int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 2; i <= n; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func filter(in <-chan int, prime int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			if i%prime != 0 {
				out <- i
			}
		}
		close(out)
	}()
	return out
}
