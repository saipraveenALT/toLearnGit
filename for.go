package main

import (
	"fmt"
)

func man() {
	var arr = [...]int{1,2,3,5,5}
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	for idx,value := range arr{
		fmt.Println(idx,value)
	}
}
