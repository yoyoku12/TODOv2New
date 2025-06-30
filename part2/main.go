package main

import "fmt"

// В цикле спрашиваем ввод транзакции: -10, 10...
// Добавлять каждую транзакцию в массив
// Вывести массив

func main() {
	transactions := []float64{}
	for {
		transaction := getUserInput()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
	}

	fmt.Printf("Ваш баланс: %.2f", calculateBalance(transactions))

}

func getUserInput() float64 {
	var transaction float64
	fmt.Print("Введите транзакцию (n для выхода): ")
	fmt.Scan(&transaction)
	return transaction
}

func calculateBalance(transactions []float64) float64 {
	sum := 0.0
	for _, v := range transactions {
		sum = sum + v
	}

	return sum
}
