package filters

func FilterNegative(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			if v >= 0 {
				out <- v
			}
		}
	}()
	return out
}

func FilterDivisibleBy3(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			if v != 0 && v%3 == 0 {
				out <- v
			}
		}
	}()
	return out
}
