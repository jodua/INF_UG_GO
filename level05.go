// [c] jodua 10042022

package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Result struct {
	name       string
	guessCount int
	date       string
}

func saveResult(results *map[string]Result, name string, guessCount int) {
	// Funkcja która aktualizuje wyniki w trakcie działania programu
	// Przekazujemy do niej wskażnik na mape wyników, nazwę gracza
	// oraz ilość prób danego gracza

	playerResult, existence := (*results)[name]
	// Sprawdzamy czy gracz wcześniej posiadał rekord
	var globalBest int = getGlobalBest(*results)
	// Pobieramy najlepszy wynik globalny

	if existence {
		// Jeżeli gracz wcześniej posiadał rekord oraz nowy rekord
		// jest mniejszy od poprzedniego rekordu, aktualizujemy
		// rekord danego gracza
		if playerResult.guessCount > guessCount {
			(*results)[name] = Result{name, guessCount, time.Now().String()}
			fmt.Println("Gratulacje, pobiłeś swój rekord")
		}
	} else {
		// W przeciwnym przypadku aktualizujemy rekord bez sprawdzania
		(*results)[name] = Result{name, guessCount, time.Now().String()}
		fmt.Println("Gratulacje, pobiłeś swój rekord")

	}

	if guessCount < globalBest {
		fmt.Println("Gratulacje, pobiłeś globalny rekord")
	}
}

func updateResultsFile(results map[string]Result) {
	// Funkcja które aktualizuje wyniki w pliku "leaderboard.csv"
	// Jako argument przyjmuje mape wyników

	f, err := os.Create("leaderboard.csv")
	if err != nil {
		panic(err)
	}
	// Tworzymy plik "leaderboard.csv"
	var line string
	for _, v := range results {
		// Dla każdego rekordu, tworzymy łancuch znaków z wszystkimi
		// informacjami, oraz zapisujemy go do pliku
		line = fmt.Sprintf("%s,%d,%s\n", v.name, v.guessCount, v.date)
		f.WriteString(line)
	}
}

func loadResults(results *map[string]Result) {
	// Funkcja która ładuje wyniki z pliku "leaderboard.csv"
	// do mapy wyników w naszym programie przekazanej jako wskaźnik

	f, err := os.ReadFile("leaderboard.csv")
	if err != nil {
		fmt.Println("Błąd przy otwieraniu pliku leaderboard.csv")
		return
	}
	// Otwieramy plik "leaderboard.csv"

	var lines []string = strings.Split(string(f), "\n")
	// Plik dzielimy na łancuchy znaków z rekordami

	var splittedLine []string
	var guessCount int
	for _, v := range lines {
		// Dla każdej niepustej linii w pliku, tworzymy nowy wynik
		// oraz zapisujemy go w mapie rekordów
		if v != "" {
			splittedLine = strings.Split(v, ",")
			guessCount, _ = strconv.Atoi(splittedLine[1])
			(*results)[splittedLine[0]] = Result{splittedLine[0], guessCount, splittedLine[2]}
		}
	}

}

func getSortedResults(results map[string]Result) []Result {
	// Funkcja która zwraca posortowany slice wyników
	// w kolejności od najmniejszej liczby prób
	// Jako argument przyjmuje mape wyników

	var resultsSlice []Result

	for _, v := range results {
		// Z mapy wyników tworzymy listę wyników
		resultsSlice = append(resultsSlice, v)
	}

	sort.Slice(resultsSlice, func(i, j int) bool {
		// Listę wyników sortujemy po ilośsci prób
		return resultsSlice[i].guessCount < resultsSlice[j].guessCount
	})

	return resultsSlice

}

func printResults(results map[string]Result) {
	// Funkcja wypisująca wyniki gry
	// Jako argument przyjmuje mape wyników

	var resultsSlice []Result = getSortedResults(results)

	for _, v := range resultsSlice {
		fmt.Printf("%s: %d - %s\n", v.name, v.guessCount, v.date)
	}
}

func getGlobalBest(results map[string]Result) int {
	// Funkcja zwracająca najlepszy wynik globalny
	// Jako argument przyjmuje mape wyników

	var resultsSlice []Result = getSortedResults(results)
	if len(resultsSlice) > 0 {
		return resultsSlice[0].guessCount
	}
	return 100
}

func main() {
	fmt.Println("Teraz będziesz zgadywać liczbę, którą wylosowałem")
	fmt.Println("\"koniec\" - zakończ program")

	// Ustawiamy losowy seed oraz losujemy liczbę do zgadnięcia
	rand.Seed(time.Now().UnixNano())
	var randomNumber int = rand.Intn(100)

	// Deklarujemy zmienne gry
	var guess string
	var guessAsNumber int
	var guessCount int
	var name string

	// Ładujemy wyniki z pliku oraz wypisujemy je
	var results map[string]Result = make(map[string]Result)
	loadResults(&results)
	printResults(results)

gameloop:
	for {
		guessCount++
		fmt.Print("Podaj liczbę: ")
		fmt.Scanln(&guess)

		if guess == "koniec" {
			fmt.Println("Żegnaj")
			updateResultsFile(results)
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
					printResults(results)
					updateResultsFile(results)
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
