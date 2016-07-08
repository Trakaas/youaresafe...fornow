package main

import (
    "fmt"
    "errors"
)

func findSubstring(s string, sub string) int {

    for i :=0; i<len(s)-len(sub); i++ {
        if s[i:i+len(sub)]==sub {
            return i
        }
    }
    return -1
}

/*func Splitstring(s string, sep string) []string {
    firstIdx  := 0
    var output []string

    for i := 0; i < len(s)-len(sep)+1; i++ {
        if s[i:i+len(sep)] == sep {
            output.append(output,s[firstIdx:i)
            firstIdx = i
        }
    }
    return output
}*/

func Trim(s string, cutset string) string {
    start :=0
    end :=len(s)
    stringChar := []rune(s)
    cutChar := []rune(cutset)
    for i := 0 ; i < len(stringChar) ; i++ {
        match := false
        for j := 0; j < len(cutChar) ; j++ {
            if stringChar[i] == cutChar[j] {
                match = true
                start++
            }
        }
        if !match {
            break
        }
    }

    for i := len(s)-1 ; i >= start+1 ; i-- {
        match := false
        for j := 0; j < len(cutChar) ; j++ {
            if stringChar[i] == cutChar[j] {
                match = true
                end--
            }
        }
        if !match {
            break
        }
    }
    if start == end {
        return s
    }
    fmt.Println(start,end)
    return s[start:end]
}

/*func parseInt(s string) (int, error) {
    neg := false
    digits := make([]byte,len(s))
    pow := 1
    val := 0
    if s[0] == '-' {
        neg=true
        s=s[1:]
    }
    for i := 0; i < len(s) ; i++ {
        if s[i] < '0' || s[i] > '9' {
            return 0 , errors.New("String contains non-numeral.")
        }
        digits[i] = s[i] - '0' 
    }

    for i := len(digits)-1; i <=0 ; i-- {
        val += int(digits[i])*pow
        pow *= 10
    }

    if neg {
        val = -val
    }
    return val,nil
} */

func main() {
    s := "#&^#@#hello##&&&#@*####"
    //sub := "el"

    a := Trim(s,"#*@^&")
    fmt.Println(a)

    fmt.Println(parseInt("45"))
}