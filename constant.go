package main

import "fmt"

func main(){
	const day int = 7

	var i int
	fmt.Println(i)

	const pi float64 = 3.14
	const secondInHours = 3600
	
	duration := 234 
	fmt.Printf("Duration in second : %v", duration*secondInHours)

	x,y := 5, 1
	fmt.Println(x/y)

	const a = 5 
	const b = 9

	// fmt.Printf(a/b)

	const n,m int = 4, 5
	const n1,m1 = 6, 7

	const (
		min1 = -89
		min2 = -90
		min3 = -78
	)
	fmt.Println(min1,min2,min3)
}