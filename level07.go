// [c] jodua 10042022

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func play(numbersRange int) {
	level01 := exec.Command("go", "run", "./level01.go")
	stdin, _ := level01.StdinPipe()
	stdout, _ := level01.StdoutPipe()

	reader := bufio.NewReader(stdout)

	go func(reader io.Reader) {
		scanner := bufio.NewScanner(reader)
		var low, mid, high int = 0, numbersRange / 2, numbersRange
		reply := fmt.Sprintf("%d\n", mid)
		fmt.Printf("[level07] %s", reply)
		stdin.Write([]byte(reply))

	gameloop:
		for scanner.Scan() {
			text := scanner.Text()
			switch {
			case strings.Contains(text, "Gratulacje, zgadłeś"):
				fmt.Println("[level01]", text)
				break gameloop
			case strings.Contains(text, "Za duża"):
				high = mid
				mid = (low + high) / 2
				fmt.Printf("[level01] %s\n", text)
				reply := fmt.Sprintf("%d\n", mid)
				fmt.Printf("[level07] %s", reply)
				stdin.Write([]byte(reply))
			case strings.Contains(text, "Za mała"):
				low = mid
				mid = (low + high) / 2
				fmt.Printf("[level01] %s\n", text)
				reply := fmt.Sprintf("%d\n", mid)
				fmt.Printf("[level07] %s", reply)
				stdin.Write([]byte(reply))
			}
		}
	}(reader)
	level01.Start()
	level01.Wait()
}

func main() {
	var numbersRange = flag.Int("r", 100, "Specify numbers range")
	flag.Parse()
	play(*numbersRange)
}
