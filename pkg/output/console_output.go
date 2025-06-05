package output

import "fmt"

// PrintBatches — печатает каждый полученный блок чисел с пояснением
func PrintBatches(in <-chan []int) {
	for batch := range in {
		for _, num := range batch {
			fmt.Printf("Получены данные: %d\n", num)
		}
	}
}
