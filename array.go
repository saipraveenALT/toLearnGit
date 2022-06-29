package main

import (
	"fmt"
)

func ma() {
	var ARR = [...]int{1, 2, 3, 4, 54}
	arrSlice := ARR[1:3]
	arrSlice[0] = 100
	sls := ARR[1:5]
	arrSlice = append(arrSlice, sls...)
	fmt.Println(arrSlice)
}
