// [c] jodua 10042022

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

type Result struct {
	name       string
	guessCount int
}

func saveResult(results *[]Result, name string, guessCount int) {
	*results = append(*results, Result{name, guessCount})
}

func sortAndPrintResults(results []Result) {
	sort.Slice(results, func(i, j int) bool {
		return results[i].guessCount < results[j].guessCount
	})
	for _, v := range results {
		fmt.Printf("%s: %d\n", v.name, v.guessCount)
	}
}

func main() {
	fmt.Println("Teraz będziesz zgadywać liczbę, którą wylosowałem")
	fmt.Println("\"koniec\" - zakończ program")

	rand.Seed(time.Now().UnixNano())
	var randomNumber int = rand.Intn(100)
	var guess string
	var guessAsNumber int
	var guessCount int

	var results []Result
	var name string

gameloop:
	for {
		guessCount++
		fmt.Print("Podaj liczbę: ")
		fmt.Scanln(&guess)
		if guess == "koniec" {
			fmt.Println("Żegnaj")
			break gameloop
		}
		guessAsNumber, _ = strconv.Atoi(guess)
		switch {
		case guessAsNumber == randomNumber:
			fmt.Println("Gratulacje, zgadłeś")
			fmt.Print("Podaj swoje imię: ")
			fmt.Scanln(&name)
			saveResult(&results, name, guessCount)
		questionloop:
			for {
				fmt.Println("Czy gramy jeszcze raz? [T/N]")
				fmt.Scanln(&guess)
				switch guess {
				case "T":
					randomNumber = rand.Intn(100)
					guessCount = 0
					break questionloop
				case "N":
					sortAndPrintResults(results)
					break gameloop
				default:
					fmt.Println("Niepoprawna opcja")
				}
			}
		case guessAsNumber < randomNumber:
			fmt.Println("Za mała")
		case guessAsNumber > randomNumber:
			fmt.Println("Za duża")
		}
	}
}
