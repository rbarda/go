package main

import (
	"fmt"
)

func main(){
	fmt.Println("Go diline hoşgeldin");
	fmt.Println(GetMultiplicationNumbers(6,9))
}

func GetMultiplicationNumbers(x int,y int) int{
	var result int=x*y
	return result
}
