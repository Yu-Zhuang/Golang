package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func plyer() int {
	return rand.Intn(100) + 1
}

func getInput(num int) int {
	// fake input for test
	return num
	// real input
	rd := bufio.NewReader(os.Stdin)

	input, err := rd.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	ret, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	return ret
}

func guessGame(max int) (bool, int) {
	rand.Seed(time.Now().UnixNano())
	ans := rand.Intn(100) + 1
	input := 0
	fmt.Printf("\t@Welcome to the guess number Game@\nEnter a num(1~100): ")

	// loop
	for count := 1; ; count++ {
		// input
		input = getInput(plyer())

		if input < ans {
			fmt.Println(input, " is too low")
		} else if input > ans {
			fmt.Println(input, "is too high")
		} else {
			fmt.Println(input, "is correct")
			fmt.Println("Congratulation, try times: ", count)
			return true, count
		}

		if count == max {
			fmt.Println("Game Over, try times: ", count)
			return false, count
		}
		fmt.Printf("Try again(1~100): ")
	}
}

func testGame(times int) (win int, lose int, avg float64) {
	total := 0
	for i := 0; i < times; i++ {
		rst, cnt := guessGame(10)
		if rst == true {
			win++
		} else {
			lose++
		}
		total += cnt
	}
	avg = float64(total) / float64(times)
	return
}

func main() {
	win, lose, avg := testGame(100)
	fmt.Println("win: ", win, "\nlose: ", lose, "\navg try times: ", avg)
}

