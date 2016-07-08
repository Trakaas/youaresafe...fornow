package main

import (
    "fmt"
    "os"
    "bufio"
//    "math/rand"
    "strconv"
//    "time"
)

func InputPrompt(input *bufio.Reader) int {
    var guess int;

    for {
        fmt.Print("Pick place on the board: ")
        data,err := input.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading from stdin.")
            return InputPrompt(input)
        }
        data = data[:len(data)-2]

        guess,err = strconv.Atoi(data)
        if err == nil && (guess >= 0 && guess <= 8) {
            break
        } else {
            fmt.Println("Your input was invalid.")
            return InputPrompt(input)
        }
    }
        return guess
}

func main() {
    input := bufio.NewReader(os.Stdin)
    a := InputPrompt(input)
    fmt.Println ("Input was: ",a)
}
