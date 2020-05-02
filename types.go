package main

import ("fmt")

func tambah(x,y float64) float64{
	return x+y
}


func main(){
	var num1 float64 = 6.6
	var num2 float64 = 6.7
	
	fmt.Println(tambah(num1,num2))
}