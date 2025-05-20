package main

import (
	"fmt"
	"math/rand"
	"time"
)

type GameRules struct {
	HiddenNumber int  
	Try          int  
	TryLimit     int  
}

func NewGameSettings(Limit int) *GameRules {
	rand.Seed(time.Now().UnixNano())
	Number := rand.Intn(10) + 1
	return &GameRules{
		HiddenNumber: Number,
		Try:          0,
		TryLimit:     Limit,
	}
}

func (g *GameRules) UserNumberCheck() (int, bool) {
	fmt.Print("Ваш вариант: ")
	var guess int
	n, err := fmt.Scanf("%d\n", &guess)
	if err != nil || n != 1 {
		fmt.Println("Ошибка! Введите целое число от 1 до 10.")
		var discard string
		fmt.Scanln(&discard)
		return 0, false
	}
	if guess < 1 || guess > 10 {
		fmt.Println("Ошибка! Число должно быть от 1 до 10!")
		return 0, false
	}

	return guess, true
}

func (g *GameRules) NumberGuess(guess int) bool {
	if guess == g.HiddenNumber {
		return true
	} else if guess < g.HiddenNumber {
		fmt.Println("Загаданное число больше.")
	} else {
		fmt.Println("Загаданное число меньше.")
	}
	return false
}

func (g *GameRules) Play() {
	fmt.Println("Угадайте число от 1 до 10!")

	for g.Try < g.TryLimit {
		guess, ok := g.UserNumberCheck()
		if !ok {
			continue
		}

		g.Try++
		if g.NumberGuess(guess) {
			fmt.Println("Вы угадали! Это", g.HiddenNumber)
			return
		}

	}

	fmt.Println("Попытки закончились. Загаданное число было:", g.HiddenNumber)
}

func main() {
	game := NewGameSettings(5) 
	game.Play()
}
