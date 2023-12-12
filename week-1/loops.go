package main

import "fmt"


func loops()  {

	// var i int = 1;

	// for i <= 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	for n:= 0 ; n < 5 ; n++ {
		fmt.Println(n)
	}
}

func fuzzbuzz(num int){
	for i := 1 ; i< num ;i++{
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("FuzzBuzz " , i )
		}else if i % 3 == 0 {
			fmt.Println("Fuzz " , i)
		}else if i % 5 ==0 {
			fmt.Println("Buzz " , i)
		}
	}
}



func printEven(num int) {

	for i:= 0 ;i < num ; i++{
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}
	
}



