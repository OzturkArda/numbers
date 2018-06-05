package main

import (
	"fmt"
)

type number int

func main() {

	numbers := []number{0,1,2,3,4,5,6,7,8,9,10}

	for _, value := range numbers{
		fmt.Println(value, value.isEvenOrOdd())
	}

}

func (n number) isEvenOrOdd () string{

	if n%2 == 0{
		return " even"
	}else{
		return " odd"
	}

}


