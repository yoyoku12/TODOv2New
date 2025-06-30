package utils

import "fmt"

func PromtDataString() string {
	var res string
	fmt.Scanln(&res)
	return res
}

func PromtDataInt() int {
	var res int
	fmt.Scanln(&res)
	return res
}

func TextLogin() {
	fmt.Println("Введите логин")
}

func TextPassword() {
	fmt.Println("Введите пароль")
}

func TextUrl() {
	fmt.Println("Введите URL")

}

func GetMenu() {
	fmt.Println("Пожалуйста выберите один из вариантов")
	fmt.Println("1 - Найти закладку по логину")
	fmt.Println("2 - Найти закладку по паролю")
	fmt.Println("3 - Найти закладку по сайту")
	fmt.Println("4 - Добавить новую закладку")
	fmt.Println("5 - Выйти из приложения")
}

func Greetings() {
	fmt.Println("Вас приветствует приложение по добавлению закладок")
}
