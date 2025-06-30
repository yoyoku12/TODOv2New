package ui

import (
	"fmt"
	"test/account"
	"test/utils"
)

func GetMenu() {
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выйти")
	choice := utils.PromtUser("Введите один из вариантов: ")
	switch choice {
	case "1":
		account.CreateAccount()
	case "2":
		name := utils.PromtUser("Введите логин для поиска аккаунта")
		account.FindAccount(name)
	case "3":
		name := utils.PromtUser("Введите логин для удаления аккаунта")
		account.DeleteAccount(name)
	case "4":
		return
	}

}

func Greeting() {
	fmt.Println("-------------------------------------------------")
	fmt.Println("Вас приветсвует программа по сохранению паролей")
	fmt.Println("-------------------------------------------------")
}
