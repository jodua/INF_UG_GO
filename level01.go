// [c] jodua 10042022

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Teraz będziesz zgadywać liczbę, którą wylosowałem")
	rand.Seed(time.Now().UnixNano())
	var randomNumber int = rand.Intn(100)
	var guess string
	var guessAsNumber int
gameloop:
	for {
		fmt.Println("Podaj liczbę: ")
		fmt.Scanln(&guess)
		guessAsNumber, _ = strconv.Atoi(guess)
		switch {
		case guessAsNumber == randomNumber:
			fmt.Println("Gratulacje, zgadłeś")
			break gameloop
		case guessAsNumber < randomNumber:
			fmt.Println("Za mała")
		case guessAsNumber > randomNumber:
			fmt.Println("Za duża")
		}
	}
}
