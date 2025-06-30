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

func CreateAccount() {
	utils.TextLogin()
	login := utils.PromtDataString()
	utils.TextPassword()
	password := utils.PromtDataString()
	utils.TextUrl()
	url := utils.PromtDataString()

	myAccount, err := NewAccount(login, password, url)
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

func FindExistingAccountByLogin(login string) (*Account, bool) {
	data, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла", err)
		return nil, false
	}
	var acc Account
	err = json.Unmarshal(data, &acc)
	if err != nil {
		fmt.Println("Ошибка разбора JSON", err)
		return nil, false
	}
	return &acc, acc.Login == login
}

func FindExistingAccountByPassword(password string) (*Account, bool) {
	data, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла", err)
		return nil, false
	}
	var acc Account
	err = json.Unmarshal(data, &acc)
	if err != nil {
		fmt.Println("Ошибка разбора JSON", err)
		return nil, false
	}
	return &acc, acc.Password == password
}

func FindExistingAccountByURL(Url string) (*Account, bool) {
	data, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла", err)
		return nil, false
	}
	var acc Account
	err = json.Unmarshal(data, &acc)
	if err != nil {
		fmt.Println("Ошибка разбора JSON", err)
		return nil, false
	}
	return &acc, acc.Url == Url
}

func OutPutData(acc Account) {
	fmt.Println(acc.Login, acc.Password, acc.Url)
}
