package main

import "fmt"

func main(){
	var a int =5
	var pointer *int =&a
	 
	fmt.Println("a =",a)
	fmt.Println("alamat a=",&a)
	
	fmt.Println("pointer =",pointer)
	fmt.Println("nilai pointer= ",*pointer)

	fmt.Println("==========================")
	fmt.Println("a diganti 7")
	a=7
	
	fmt.Println("pointer =",pointer)
	fmt.Println("nilai pointer= ",*pointer)
}