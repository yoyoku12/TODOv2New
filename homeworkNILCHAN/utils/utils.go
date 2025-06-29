package utils

import (
	"TODOlist/functional"
	"TODOlist/logging"
	"bufio"
	"database/sql"
	"fmt"
	"os"
)

func PromtUser() string {
	scanner := bufio.NewScanner(os.Stdin)
	if ok := scanner.Scan(); !ok {
		return ("Ошибка ввода!")
	}
	return scanner.Text()
}

func Greeting() {
	fmt.Println("Вас привествует приложение для работы с TODO лист!")
}

func GetMenu() {
	fmt.Println("Cписок доступных команд:")
	fmt.Println("help")
	fmt.Println("add")
	fmt.Println("del")
	fmt.Println("events")
	fmt.Println("mark as done")
	fmt.Println("exit")
}

func ReadFile() {
	data, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Ошибка чтения", err)
		return
	}
	fmt.Println(string(data))
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Запись успешна")

}

func FullMenuList(db *sql.DB) {
	for {
		GetMenu()
		choice := PromtUser()
		switch choice {
		case "help":
			functional.Help()
		case "add":
			logging.AddEventToDb(db, choice)
			fmt.Println("Введите заголовок:")
			title := PromtUser()
			fmt.Println("Введите полное описание:")
			description := PromtUser()
			err := functional.NewItem(db, title, description)
			if err != nil {
				fmt.Println("Ошибка при добавлении задачи:", err)
			}
		case "del":
			logging.AddEventToDb(db, choice)
			fmt.Println("Введите заголовок для удаления задания")
			title := PromtUser()
			msg := functional.DeleteItem(db, title)
			fmt.Println(msg)
		case "list":
			logging.AddEventToDb(db, choice)
			functional.ShowFullList(db)
		case "mark as done":
			logging.AddEventToDb(db, choice)
			fmt.Println("Введите заголовок задания, который хотите отметить как выполненный:")
			title := PromtUser()
			err := functional.MarkAsDoneByTitle(db, title)
			if err != nil {
				fmt.Println("Ошибка при отметке задачи:", err)
			}
		case "events":
			logging.AddEventToDb(db, choice)
			functional.ShowFullListEvents(db)
		case "exit":
			logging.AddEventToDb(db, choice)
			return
		default:
			fmt.Println("Неизвестная команда. Введите 'help' для списка команд.")
		}
	}
}
