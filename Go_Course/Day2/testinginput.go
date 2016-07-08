package main

import (
    "fmt"
    "os"
    "bufio"
    "math/rand"
    "strconv"
    "time"
) 

func main() {
    input := bufio.NewReader(os.Stdin)

    t := time.Now()
    u := t.UnixNano()
    rand.Seed(u)

    
    answer := rand.Intn(10) + 1
    
    var guess int
    for {
        fmt.Print("Pick a number between 1 and 10: ")
        data,err := input.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading from stdin.")
            return
        }
        data = data[:len(data)-2]

        guess,err = strconv.Atoi(data)
        if err == nil && (guess >= 1 && guess <= 10) {
            break
        } else {
            fmt.Println("Your input was invalid.")
        }
    }
        if answer == guess {
            fmt.Println("You guessed correctly!")
        } else {
            fmt.Print("Nope, the answer was ", answer)
            fmt.Println(".")
        }

}