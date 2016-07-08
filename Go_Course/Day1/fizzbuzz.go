package main

import (
"fmt"
)

func main() {
    var i int
    var by3 bool
    var by5 bool
    
    i=1
    for i<100 {
        by3 = i%3 == 0
        by5 = i%5 == 0

        if (by3 && by5 == true) {
            fmt.Println("FizzBuzz")
        } else if (by3 == true) {
            fmt.Println("Fizz")
        } else if (by5 == true) {
            fmt.Println("Buzz")
        } else { fmt.Println(i) }

        i++
    }
}