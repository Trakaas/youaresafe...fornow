package main

import (
	"fmt"
	"strconv"
)

func setMsbitTo1(a byte) byte {
	var sig byte = 128
	return a | sig
}

func setMsbitTo0(a byte) byte {
	var sig byte = 127
	return a & sig
}

func flipHighBit(a byte) byte {
	var sig byte = 129
	return a ^ sig
}

func main() {
	var a byte = 255 //0000_1100
	var max byte = 128
	//a = setMsbitTo1(a)
	fmt.Println(a)
	fmt.Print("'", strconv.FormatInt(int64(a), 2), "'\n")

	a = setMsbitTo0(a)
	/*var b byte = 8  //0000_1000
	//a = ^a       //logical not
	a = a | b //0000_1100
	a = a & b //0000_1000
	a = a ^ b // xor 0000_0100*/
	fmt.Println(a)
	fmt.Print("'", strconv.FormatInt(int64(a), 2), "'\n")
	fmt.Println(max)
	fmt.Print("'", strconv.FormatInt(int64(max), 2), "'\n")

}
