package files

import (
	"fmt"
	"os"
)

func ReadFile() {
	data, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
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
