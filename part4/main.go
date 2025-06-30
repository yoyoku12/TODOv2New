package main

import (
	"fmt"

	"demo/account"
	"demo/files"
)

func main() {
	createAccount()

}

func createAccount() {
	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	url := promtData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат ввода данных", err)
		return
	}
	file, err := myAccount.ToBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать данные в JSON")
		return
	}
	files.WriteFile(file, "data.json")

}

func promtData(promt string) string {
	fmt.Println(promt)
	var res string
	fmt.Scanln(&res)
	return res
}
