package buffer

import (
	"log"
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
						log.Printf("[Buffer] Конец входного потока. Отправка оставшихся %d элементов", len(buffer))
						out <- buffer
					}
					log.Printf("[Buffer] Завершение буфера")
					return
				}
				buffer = append(buffer, val)
				log.Printf("[Buffer] Добавлено в буфер: %d (размер: %d)", val, len(buffer))
				if len(buffer) >= bufferSize {
					log.Printf("[Buffer] Буфер заполнен. Отправка: %v", buffer)
					out <- buffer
					buffer = make([]int, 0, bufferSize)
				}
			case <-ticker.C:
				if len(buffer) > 0 {
					log.Printf("[Buffer] Таймер. Отправка буфера: %v", buffer)
					out <- buffer
					buffer = make([]int, 0, bufferSize)
				} else {
					log.Println("[Buffer] Таймер. Буфер пуст, ничего не отправлено")
				}
			case <-done:
				log.Println("[Buffer] Получен сигнал завершения")
				return
			}
		}
	}()

	return out
}
