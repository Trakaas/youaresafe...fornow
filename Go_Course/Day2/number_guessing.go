package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	input := bufio.NewReader(os.Stdin)

	t := time.Now()   // returns a struct representing a datetime
	u := t.UnixNano() // the datetime represented as the number of nanoseconds since midnight, Jan 1st, 1970
	rand.Seed(u)      // seed the defaut random source; sets a global variable that is used by rand.Intn

	// pick a random number between 1 and 10 (inclusive)
	answer := rand.Intn(10) + 1 // Intn return a number 0 up to (but not including) its arg, so in this case 0 to 9

	var guess int // this has to be declared outside the loop because we want to use it outside the loop
	for {
		fmt.Println("Pick a number between 1 and 10:")
		data, err := input.ReadString('\n') // return input up to (and including) LF
		if err != nil {
			fmt.Println("Error reading from standard input.")
			return
		}
		data = data[:len(data)-2] // slice string to get substring (hack off last two characters: CR LF)
		// Atoi = Ascii to Integer
		guess, err = strconv.Atoi(data) // do not use := because guess and err are already declared
		if err == nil && (guess >= 1 && guess <= 10) {
			break
		} else {
			fmt.Println("Your input was invalid.")
		}
	}
	if answer == guess {
		fmt.Println("You guessed correctly!")
	} else {
		fmt.Println("No, dummy! The answer was", answer)
	}
}
