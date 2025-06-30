package main

import (
	"fmt"
	"test/account"
	"test/utils"
)

func main() {

	utils.Greetings()

	for {

		utils.GetMenu()
		choice := utils.PromtDataInt()
		switch choice {
		case 1:
			fmt.Println("Введите логин для поиска")
			result, isLoginExists := account.FindExistingAccountByLogin(utils.PromtDataString())
			if isLoginExists {
				fmt.Println(result)
			} else {
				fmt.Println("Такого логина не существует")
			}
		case 2:
			fmt.Println("Введите пароль для поиска")
			result, isPasswordExists := account.FindExistingAccountByPassword(utils.PromtDataString())
			if isPasswordExists {
				fmt.Println(result)
			} else {
				fmt.Println("Такого пароля не существует")
			}
		case 3:
			fmt.Println("Введите URL для поиска")
			result, isUrlExists := account.FindExistingAccountByURL(utils.PromtDataString())
			if isUrlExists {
				fmt.Println(result)
			} else {
				fmt.Println("Такого URL не существует")
			}
		case 4:
			account.CreateAccount()
		case 5:
			fmt.Println("Выход из программы...")
			return
		}
	}
}
