package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"os"
	"test/files"
	"test/utils"

	"time"
)

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc *Account) OutputPassword(n int) {
	fmt.Println(acc.Login, acc.Password, acc.Url)
}

func (acc *Account) ToBytes() ([]byte, error) {
	file, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (acc *Account) generatePassword(n int) {
	var letterRunes = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890")
	newpassword := []rune{}
	for i := 0; i < n; i++ {
		randomvalue := rand.IntN(len(letterRunes))
		newpassword = append(newpassword, letterRunes[randomvalue])
	}
	acc.Password = string(newpassword)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID URL")
	}

	newAcc := &Account{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Login:     login,
		Password:  password,
		Url:       urlString,
	}

	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}

func CreateAccount() {
	login := utils.PromtUser("Введите логин")
	password := utils.PromtUser("Введите пароль")
	url := utils.PromtUser("Введите URL")

	myAccount, err := NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат ввода данных", err)
		return
	}

	data, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
	}

	var accounts []Account

	json.Unmarshal(data, &accounts)

	accounts = append(accounts, *myAccount)
	file, err := json.MarshalIndent(accounts, "", "  ")
	if err != nil {
		fmt.Println("Ошибка сериализации в JSON:", err)
		return
	}
	files.WriteFile(file, "data.json")

}

func FindAccount(str string) {
	data, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var accounts []Account
	if err := json.Unmarshal(data, &accounts); err != nil {
		fmt.Println("Ошибка чтения данных:", err)
		return
	}

	for _, acc := range accounts {
		if acc.Login == str {
			fmt.Println(acc.Login, acc.Password, acc.Url)
			return
		}
	}

	fmt.Println("Такого аккаунта не существует")
}

func DeleteAccount(str string) {
	data, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
	}
	var accounts []Account

	json.Unmarshal(data, &accounts)

	var newAccounts []Account

	for _, acc := range accounts {
		if acc.Login != str {
			newAccounts = append(newAccounts, acc)
		}
	}

	newdata, err := json.Marshal(newAccounts)
	if err != nil {
		fmt.Println(err)
	}
	files.WriteFile(newdata, "data.json")

	if len(accounts) >= len(newAccounts) {
		fmt.Println("Удаление прошло успешно")
	} else {
		fmt.Println("Аккаунт с таким логином не найден")
	}

}
