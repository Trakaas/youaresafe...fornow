package main

import ( 
    "fmt"
    "os"
    "bufio"
    "math/rand"
    "strconv"
    "time" )

func Drawboard(board [9]int) {
    var printQueue[9] string
    for i :=0 ; i < 9; i++ {
        if board[i] == 5 {
            printQueue[i] = "X"
        } else if board[i] == 3 {
            printQueue[i] = "O"
        } else {
            printQueue[i] = strconv.Itoa(i)
        }
    }

    fmt.Println()
    fmt.Println("\t", printQueue[0] ,"|", printQueue[1] ,"|", printQueue[2])
    fmt.Println("\t","- - - - -")
    fmt.Println("\t", printQueue[3] ,"|", printQueue[4] ,"|", printQueue[5])
    fmt.Println("\t","- - - - -")
    fmt.Println("\t", printQueue[6] ,"|", printQueue[7] ,"|", printQueue[8])
    fmt.Println("\n\t","- - - - -")
}

func Moveprompt(input *bufio.Reader) int {
    var guess int
    var err error
    var data string

    for {
        fmt.Print("Pick place on the board (-1 to exit game): ")
        data,err = input.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading from stdin.")
            return Moveprompt(input)
        }
        data = data[:len(data)-2]

        guess,err = strconv.Atoi(data)
        if err == nil && (guess >= -1 && guess <= 8) {
            break
        } else {
            fmt.Println("Your input was invalid.")
            return Moveprompt(input)
        }
    }
        return guess
}

func Exitprompt(input *bufio.Reader, currentState int) bool {
    var err error
    var data string
    if currentState == 8 {
            fmt.Println("Game is a time.")
            fmt.Print("Play again(yn)? ")
            data,err = input.ReadString('\n')
            if err != nil {
                fmt.Println("Error reading from stdin.")
            }
            data = data[:len(data)-2]

            if data == "n" || data == "N" || data == "no" || data == "No" {
                return false
            }
    } else if currentState == 5 {
        fmt.Println("Computer wins.")
        fmt.Print("Play again(yn)? ")
            data,err = input.ReadString('\n')
            if err != nil {
                fmt.Println("Error reading from stdin.")
            }
            data = data[:len(data)-2]

            if data == "n" || data == "N" || data == "no" || data == "No" {
                return false
            }
    } else if currentState == 3 {
        fmt.Println("Player wins.")
        fmt.Print("Play again(yn)? ")
            data,err = input.ReadString('\n')
            if err != nil {
                fmt.Println("Error reading from stdin.")
            }
            data = data[:len(data)-2]

            if data == "n" || data == "N" || data == "no" || data == "No" {
                return false
            }
    }
    return true
}

func Checkgrid(grid [9]int) int {
    // check horizontal
    for i := 0; i <= 6; i=i+3 {
        if grid[i]+grid[i+1]+grid[i+2] == 9 {
            return 3
        } else if grid[i]+grid[i+1]+grid[i+2] == 15 {
            return 5
        }
    }
    // check vertical
    for i := 0; i <= 2; i++ {
        if grid[i]+grid[i+3]+grid[i+6] == 9 {
            return 3
        } else if grid[i]+grid[i+3]+grid[i+6] == 15 {
            return 5
        }
    }
    // check diagonals
    if grid[0] + grid[4] + grid[8] == 9 {
        return 3
    } else if grid[0] + grid[4] + grid[8] == 15 {
        return 5
    }

    if grid[2] + grid[4] + grid[6] == 9 {
        return 3
    } else if grid[2] + grid[4] + grid[6] == 15 {
        return 5
    }
    // check tie
    spaceLeft := false
    for i := 0; i < 9; i++ {
        if grid[i] == -100 {
            spaceLeft =true
            break
        }
    }
    // space is left, no tie, return -100
    if spaceLeft {
        return -100
    }
    //tie condition 5+3=8
    return 8
}

func main(){
    // initialize grid
    board := [9]int{-100,-100,-100,-100,-100,-100,-100,-100,-100}
    var compChoice int
    var playerChoice int
    var currentState int
    gameOn := true
    // initialize io and seed computer generator
    input := bufio.NewReader(os.Stdin)
    t := time.Now()
    u := t.UnixNano()
    rand.Seed(u)

    for gameOn {
        compChoice = rand.Intn(9)

        // take turns
        for board[compChoice] != -100 {
            compChoice = rand.Intn(9)
        }

        board[compChoice] = 5

        // show computer move
        Drawboard(board)

        // check if computer won
        currentState = Checkgrid(board)
        if currentState != -100 {
            gameOn = Exitprompt(input, currentState)
            if gameOn == false {
                fmt.Println("\nThanks for playing!")
                break
            } else {
                board = [9]int{-100,-100,-100,-100,-100,-100,-100,-100,-100}
                continue
            }
        }
        
        playerChoice = Moveprompt(input)
        
        if playerChoice == -1 {
            gameOn = false
            fmt.Println("\nThanks for playing!")
            break
        }

        for board[playerChoice] != -100 {
            fmt.Println("Place already taken.")
            playerChoice = Moveprompt(input)
        }

        board[playerChoice] = 3
    
        // show player move
        Drawboard(board)

        // check if player won
        currentState = Checkgrid(board)
        gameOn = Exitprompt(input, currentState)
        if currentState != -100 {
            gameOn = Exitprompt(input, currentState)
            if gameOn == false {
                fmt.Println("\nThanks for playing!")
                break
            } else if gameOn == true {
                board = [9]int{-100,-100,-100,-100,-100,-100,-100,-100,-100}
                continue
            }
        }

    }
    
}