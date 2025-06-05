package main

import (
	"pipeline/pkg/buffer"
	"pipeline/pkg/filters"
	"pipeline/pkg/input"
	"pipeline/pkg/output"
	"time"
)

func main() {
	inputChan := make(chan int)
	done := make(chan struct{})

	// Старт источника данных из консоли
	go input.ReadFromConsole(inputChan)

	// Стадии обработки
	stage1 := filters.FilterNegative(inputChan)
	stage2 := filters.FilterDivisibleBy3(stage1)
	stage3 := buffer.BufferInts(stage2, 5*time.Second, 5, done)

	// Потребитель выводит в консоль
	output.PrintBatches(stage3)
}
