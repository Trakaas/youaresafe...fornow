package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	fmt.Print("Type change amount (.XX) format: ")
	userChangeS, err := input.ReadString('\n')
	coinTypes := [...]int{25, 10, 5, 1} //quarter,dime,nickel,penny
	var count [4]int

	if err != nil {
		fmt.Println(err)
		return
	}

	userChangeS = userChangeS[1 : len(userChangeS)-2]
	userChangeF, err := strconv.Atoi(userChangeS)

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 4; i++ {
		estimate := userChangeF / coinTypes[i]
		count[i] = int(math.Floor(float64(estimate)))

		if count[i] > 0 {
			userChangeF = userChangeF % (coinTypes[i] * count[i])
		}

	}

	fmt.Println("You will need", count[0], "quarters,", count[1], "dimes,",
		count[2], "nickels, and", count[3], "pennies.")
}
