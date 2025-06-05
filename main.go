package main

import (
	"log"
	"os"
	"pipeline/pkg/buffer"
	"pipeline/pkg/filters"
	"pipeline/pkg/input"
	"pipeline/pkg/output"
	"time"
)

func main() {
	logFile, err := os.OpenFile("pipeline.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Не удалось открыть лог-файл: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Println("[Main] Запуск пайплайна")

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

	close(done)
	log.Println("[Main] Пайплайн завершён")
}
