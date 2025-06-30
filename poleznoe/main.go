package main

// // Reverse slice
// func main() {
// 	a := []int{1, 2, 3, 4}
// 	reverse(a)
// 	fmt.Println(a)

// }

// func reverse(arr []int) {
// 	// Создаём копию исходного слайса
// 	copyArr := make([]int, len(arr))
// 	copy(copyArr, arr)

// 	// Читаем значения из копии, записываем в оригинал
// 	for index, value := range copyArr {
// 		arr[len(arr)-1-index] = value
// 	}
// }

// // Reverse massive
// func main() {
// 	a := [4]int{1, 2, 3, 4}
// 	reverse(&a)
// 	fmt.Println(a)

// }

// func reverse(arr *[4]int) {
// 	for index, value := range *arr {
// 		(*arr)[len(arr)-1-index] = value
// 	}
// }
