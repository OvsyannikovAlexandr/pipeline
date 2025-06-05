package generator

import (
	"time"
)

func Generate(out chan<- int, data []int) {
	defer close(out)
	for _, val := range data {
		out <- val
		time.Sleep(500 * time.Millisecond)
	}
}
