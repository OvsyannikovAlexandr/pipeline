package input

import (
	"bufio"
	"fmt"
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
		if line == "exit" {
			break
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("⚠️ Введите целое число или 'exit' для выхода.")
			continue
		}

		out <- num
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения:", err)
	}
}
