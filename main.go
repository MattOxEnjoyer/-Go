package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(10) + 1

	var UserNumber int
	Try := 0
	TryLimit := 5

	fmt.Println("Угадайте число от 1 до 10!")

	for Try < TryLimit {
		fmt.Print("Ваш вариант: ")

		n, err := fmt.Scanf("%d\n", &UserNumber)
		if err != nil || n != 1 {
			fmt.Println("Ошибка! Введите целое число от 1 до 10.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		if UserNumber < 1 || UserNumber > 10 {
			fmt.Println("Ошибка! Число должно быть от 1 до 10!")
			continue
		}

		Try++

		if UserNumber == number {
			fmt.Println("Вы угадали! Это", number)
			return
		} else if UserNumber < number {
			fmt.Println("Загаданное число больше.")
		} else {
			fmt.Println("Загаданное число меньше.")
		}
	}

	fmt.Println("Попытки закончились. Загаданное число было:", number)
}
