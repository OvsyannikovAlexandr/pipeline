package output

import (
	"fmt"
	"log"
)

// PrintBatches — печатает каждый полученный блок чисел с пояснением
func PrintBatches(in <-chan []int) {
	for batch := range in {
		for _, num := range batch {
			log.Printf("[Output] Получены данные: %d", num)
			fmt.Printf("Получены данные: %d\n", num)
		}
	}
	log.Println("[Output] Завершение приёмника")
}
