package main

import (
	"TODOlist/database"
	"TODOlist/utils"
	"fmt"
)

func main() {
	fmt.Println("Запуск программы..")
	db := database.Connect()
	defer db.Close()
	utils.Greeting()
	utils.FullMenuList(db)

}
