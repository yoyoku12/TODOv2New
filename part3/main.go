package main

import "fmt"

func main() {

	m := map[string]string{}

	cycleMenu(m)

}

func cycleMenu(m map[string]string) {
	for {
		choice := getMenu()
		if choice == 1 {
			showAllBookmarks(m)
			fmt.Println("Введите 1, чтобы открыть меню снова")
			fmt.Println("Введите 2, чтобы завершить работу.")
			choice2 := getUserInputInt()
			if choice2 == 1 {
				continue

			} else if choice2 == 2 {
				fmt.Println("Работа закладок завершена")
				break
			} else {
				fmt.Println("Вы ввели неверное значение")
			}

		} else if choice == 2 {
			addNewBookmark(m)

		} else if choice == 3 {
			deleteBookmark(m)
		} else if choice == 4 {
			fmt.Println("Работа закладок завершена")
			break
		} else {
			fmt.Println("Вы ввели неверное значение")
		}

	}

}

func getMenu() int {
	fmt.Println("____Закладки____")
	fmt.Println("Введите 1, чтобы посмотреть все закладки.")
	fmt.Println("Введите 2, чтобы добавить закладку.")
	fmt.Println("Введите 3, чтобы удалить закладку.")
	fmt.Println("Введите 4, чтобы завершить работу.")
	var choice int
	fmt.Scan(&choice)
	return choice

}

func showAllBookmarks(m map[string]string) {
	if len(m) == 0 {
		fmt.Println("Пока нет закладок")
	} else if len(m) > 0 {
		fmt.Println("---------------------")
		fmt.Println("Список всех закладок:")
		for i, v := range m {
			fmt.Println(i, v)
		}
		fmt.Println("---------------------")
	}

}

func addNewBookmark(m map[string]string) *map[string]string {
	fmt.Println("Введите короткое название закладки:")
	var shortname string = getUserInputString()
	fmt.Println("Введите сайт:")
	var site string = getUserInputString()
	m[shortname] = site
	return &m
}

func deleteBookmark(m map[string]string) *map[string]string {
	fmt.Println("Введите короткое название закладки для удаления:")
	var key string = getUserInputString()
	delete(m, key)
	return &m
}

func getUserInputString() string {
	var input string
	fmt.Scan(&input)
	return input
}

func getUserInputInt() int {
	var input int
	fmt.Scan(&input)
	return input
}
