package buffer

import (
	"time"
)

func BufferInts(in <-chan int, flushInterval time.Duration, bufferSize int, done <-chan struct{}) <-chan []int {
	out := make(chan []int)
	buffer := make([]int, 0, bufferSize)
	ticker := time.NewTicker(flushInterval)

	go func() {
		defer close(out)
		defer ticker.Stop()

		for {
			select {
			case val, ok := <-in:
				if !ok {
					if len(buffer) > 0 {
						out <- buffer
					}
					return
				}
				buffer = append(buffer, val)
				if len(buffer) >= bufferSize {
					out <- buffer
					buffer = make([]int, 0, bufferSize)
				}
			case <-ticker.C:
				if len(buffer) > 0 {
					out <- buffer
					buffer = make([]int, 0, bufferSize)
				}
			case <-done:
				return
			}
		}
	}()

	return out
}
