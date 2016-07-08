package main

import ("fmt")

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

func find(a []int, val int) int {
    for i:=0; i<len(a); i++ {
        if a[i]==val {
            return val
        }
    }
    return -1
}
   

func binarySearch(a []int, val int) int{
    mid := (len(a)-1)/2
    if mid < 0 {
        return -1
    }
    
    if val > a[mid] {
        return binarySearch(a[mid+1:len(a)-1], val)
    } else if val < a[mid] {
        return binarySearch(a[0:mid-1], val)
    } else {
        return mid
    }
    return -1
}

func bubbleSort(a[]int) {
    for pass := 0; pass < len(a)-1; pass++ {
        for i := 1; i < len(a)-pass; i++ {
            if a[i-1] > a[i] {
                a[i-1],a[i] = a[i], a[i-1]
            }
        }
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
    var full_slice = array[0:1000]
    fmt.Println("Print slice - array[2:8]:")
    printAll(my_slice)
    fmt.Println("Print pairs - array[2:8]:")
    printPairs(my_slice)
    fmt.Println("Find 235 linear search:")
    fmt.Print("\t")
    fmt.Println(find(full_slice,235))
    fmt.Println("Find 235 binary search:")
    fmt.Print("\t")
    fmt.Println(binarySearch(full_slice,235)) 
    random_array := [...]int{3,5,2,9,5,3,7,5,0}
    random_slice := random_array[:]
    bubbleSort(random_slice)
    fmt.Println(random_slice)


}