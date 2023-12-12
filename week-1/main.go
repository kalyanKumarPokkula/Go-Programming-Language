package main

import "fmt"

func main () {

	// Variable 

	var i, j, k int = 1,2,3

	const fruit = "apple"
	
	fmt.Println(i,j,k)

	// Data types 

	// bool

	var b bool = true

	// string

	// var str string = "My name is kalyan"

	// int  int8  int16  int32  int64
	var num int = 23;
	// uint uint8 uint16 uint32 uint64 uintptr

	// byte // alias for uint8

	// rune // alias for int32
	// 	// represents a Unicode code point

	// float32 float64

	var flo float32 = 2.00

	// complex64 complex128

	

	

	fmt.Println( num, flo , b)

	// printEven(10)
	// fmt.Println()

	// fuzzbuzz(30)
	// fmt.Println()

	var str string= "Hello World"

    for i , c := range str {
		fmt.Printf("%d %c", i , c)
		fmt.Println()
	}

	fmt.Println(str)



}

