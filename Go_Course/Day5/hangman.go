package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var words = [...]string{"food", "rubber", "elevator", "wombat", "rotary", "heuristic", "whatever"}
	const maxMisses = 8
	rand.Seed(time.Now().UnixNano())
	word := words[rand.Intn(len(words))]
	input := bufio.NewReader(os.Stdin)
	misses := 0

	blank := "___________________________________________"
	left := blank[:len(word)]

	for {
		fmt.Print("Type a letter: ")
		guess, err := input.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from stdin.")
			continue
		}
		guess = guess[:len(guess)-2]
		match := false
		for i := 0; i < len(word); i++ {
			if guess == string(word[i]) {
				match = true
				var end string
				if i != len(word)-1 {
					end = left[i+1:]
				}
				left = left[:i] + guess + end
			}
		}

		if !match {
			misses++
		}

		notDone := true
		for i := 0; i < len(left); i++ {
			if left[i] == '_' {
				notDone = false
				break
			}
		}
		if notDone {
			fmt.Println("You win.")
			return
		}

		if misses == maxMisses {
			fmt.Println("You lose. Word was:", word)
			return
		}
		fmt.Println("Word:", left, "\tGuesses left: ", misses, "of", maxMisses)
	}
}
