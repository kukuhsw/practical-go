package main

import "fmt"

func main(){

	a,b := 8,9
	fmt.Println("Penjumlahan",a+b, "Rataan", (a+b)/2)
		
	fmt.Printf("Usia anda %d\n", 34)

	fmt.Printf("masukkan x: %d, masukkan y: %f\n", 4, 6.9)

	fmt.Printf("Cara menulis petik \"Hallo\"\n")

	figure := "Lingkaran"
	radius := 5
	pi := 3.14
	_, _ = figure,pi
	fmt.Printf("Figure:%s\t Radius :%d\t Pi:%f \n", figure, radius, float64(radius)*2*pi)

	// %q cetak string
	fmt.Printf("Tes string : %q\n", figure)

	// %v bisa diganti dengan tipe apapun
	fmt.Printf("Figure:%v\t Radius :%v\t Pi:%v \n", figure, radius, float64(radius)*2*pi)

	// %T -> mengetahui tipe data
	fmt.Printf("Figure tipenya\t :%T\n",figure)
	fmt.Printf("Radius tipenya\t :%T\n",radius)
	fmt.Printf("Pi tipenya\t : %T\n",pi)

	// %t -> boolean
	closed := true
	fmt.Printf("File ditutup %t \n", closed)

	// %b -> base 2
	fmt.Printf("Basis dua dari 88 : %b\n", 88)

	x := 6.7
	y := 9.9
	fmt.Printf("Perkalian x * y : %0.3f", x*y)
}