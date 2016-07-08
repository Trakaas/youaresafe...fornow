package main

import (
"fmt" )


func printAll( a []int) {
    for i := 0; i < len(a); i++ {
        fmt.Print("\t")
        fmt.Println(a[i])
    }
}

func printPairs( a []int) {
    for i:=1; i<len(a); i++ {
    second,first := a[i],a[i-1]
    fmt.Print("\t")
    fmt.Println(first+second)
    }
}

func main() {
    i := 400
    var array[1000] int

    for j := 0; j < 1000; j++ {
        array[j]=i
        i--
//        fmt.Println(array[j])
    }
    var my_slice = array[2:8]
    fmt.Println("Print slice - array[2:8]:")
    printAll(my_slice)
    fmt.Println("Print pairs - array[2:8]:")
    printPairs(my_slice)
}