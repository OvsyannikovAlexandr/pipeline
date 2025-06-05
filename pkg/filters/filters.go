package filters

import "log"

func FilterNegative(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			if v >= 0 {
				log.Printf("[FilterNegative] Пропущено: %d", v)
				out <- v
			} else {
				log.Printf("[FilterNegative] Отброшено отрицательное: %d", v)
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
				log.Printf("[FilterDivisibleBy3] Пропущено: %d", v)
				out <- v
			} else {
				log.Printf("[FilterDivisibleBy3] Отброшено (некратно 3 или 0): %d", v)
			}
		}
	}()
	return out
}
