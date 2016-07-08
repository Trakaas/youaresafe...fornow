package main

import (
    "fmt"
    "os"
    "bufio"
    "math/rand"
    "time"
) 

func find(a [3]string, val string) int {
    for i:=0; i<len(a); i++ {
        if a[i]==val {
            return i
        }
    }
    return -1
}

func main() {
    input := bufio.NewReader(os.Stdin)

    t := time.Now()
    u := t.UnixNano()
    rand.Seed(u)
    
    possible_choices := [3]string{"paper","rock","scissors"}
    comp_choice_ind := rand.Intn(3)
    comp_choice := possible_choices[comp_choice_ind]
    var player_choice string = "rock"
    var player_choice_ind int
    var err error

    for {
        fmt.Print("Rock, paper or scissors? ")
        player_choice, err = input.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading from stdin.")
            continue
        }
        
        player_choice = player_choice[:len(player_choice)-2]

        player_choice_ind = find(possible_choices,player_choice)

        if player_choice_ind != -1 {
            break
        } else {
            fmt.Println("Your input was invalid.")
        }
    }

    if player_choice_ind == comp_choice_ind {
        fmt.Println("Tie!")
    } else if player_choice_ind < comp_choice_ind {
        if player_choice_ind + 1 == comp_choice_ind {
            fmt.Println("You win.","Computer chose:",comp_choice)
        } else {
            fmt.Println("You lose.","Computer chose:",comp_choice)
        }
    } else if player_choice_ind > comp_choice_ind {
        if player_choice_ind - 1 == comp_choice_ind {
            fmt.Println("You lose.","Computer chose:",comp_choice)
        } else {
            fmt.Println("You win.","Computer chose:",comp_choice)
        }
    }
}