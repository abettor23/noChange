package main

import "fmt"

func main() {
	fmt.Println("В кассе отсутствует сдача, а у вас всего три купюры.")
	fmt.Println("Сможете ли вы расплатиться без сдачи?")

	fmt.Println("Введите стоимость товара:")
	var productPrice int
	fmt.Scan(&productPrice)

	fmt.Println("Введите номинал первой купюры:")
	var banknoteOne int
	fmt.Scan(&banknoteOne)

	fmt.Println("Введите номинал второй купюры:")
	var banknoteTwo int
	fmt.Scan(&banknoteTwo)

	fmt.Println("Введите номинал третьей купюры:")
	var banknoteThree int
	fmt.Scan(&banknoteThree)

	if banknoteOne == productPrice || banknoteTwo == productPrice || banknoteThree == productPrice {
		fmt.Println("Вы можете заплатить за товар без сдачи одной купюрой.")
	} else if banknoteOne+banknoteTwo == productPrice || banknoteOne+banknoteThree == productPrice || banknoteTwo+banknoteThree == productPrice {
		fmt.Println("Вы можете заплатить за товар без сдачи двумя купюрами.")
	} else if banknoteOne+banknoteTwo+banknoteThree == productPrice {
		fmt.Println("Вы можете заплатить за товар без сдачи всеми купюрами.")
	} else {
		fmt.Println("Вы не можете заплатить за товар без сдачи.")
	}
}
