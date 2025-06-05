package input

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadFromConsole — читает числа с консоли, фильтрует нечисловой ввод
func ReadFromConsole(out chan<- int) {
	defer close(out)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите целые числа по одному на строку. Для выхода введите 'exit':")

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		log.Printf("[Input] Ввод: %s", line)

		if line == "exit" {
			log.Println("[Input] Завершение ввода по команде 'exit'")
			break
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			log.Printf("[Input] !!! Нецелочисленный ввод: %s", line)
			fmt.Println("!!! Введите целое число или 'exit' для выхода.")
			continue
		}

		log.Printf("[Input] Отправка числа в конвейер: %d", num)
		out <- num
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения:", err)
	}
}
