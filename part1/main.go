package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	for {
		readyOutputIMT()
		isRepeatCalculation := chechRepeatCalculation()
		if !isRepeatCalculation {
			break
		}
	}
}

func chechRepeatCalculation() bool {
	var userChoice string
	fmt.Print("Вы хотите сделать еще расчет? (y/n)?: ")
	fmt.Scan(&userChoice)
	if userChoice == "y" || userChoice == "Y" {
		return true
	}
	return false
}

func readyOutputIMT() {
	fmt.Println("__ Калькулятор индекса массы тела __")
	userHeight, userKg := getUserInputForIMT()
	IMT, err := calculateIMT(userKg, userHeight)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	outputResult(IMT)
	fmt.Println(imtTable(IMT))
}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", IMT)
	fmt.Println(result)

}

func calculateIMT(userKg float64, userHeight float64) (float64, error) {
	if userKg <= 0 || userHeight <= 0 {
		return 0, errors.New("NO_PARAMS_ERROR")
	}
	const IMTPower = 2
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT, nil
}

func getUserInputForIMT() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	return userHeight, userKg
}

func imtTable(imt float64) string {
	switch {
	case imt < 16:
		return "Сильный дефицит массы тела"
	case imt < 18.5:
		return "Дефицит массы тела"
	case imt <= 25:
		return "Норма"
	case imt <= 30:
		return "Избыточная масса"
	case imt <= 35:
		return "1-я степень ожирения"
	case imt <= 40:
		return "2-я степень ожирения"
	default:
		return "3-я степень ожирения"
	}
}
