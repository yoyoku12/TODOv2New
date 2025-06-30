package utils

import (
	"bufio"
	"fmt"

	"os"
)

func PromtUser(prompt string) string {
	fmt.Print(prompt + " ")
	scanner := bufio.NewScanner(os.Stdin)
	if ok := scanner.Scan(); !ok {
		return "Ошибка ввода!"
	}
	return scanner.Text()
}
