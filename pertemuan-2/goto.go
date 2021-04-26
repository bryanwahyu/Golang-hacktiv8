package main

import "fmt"

func main(){
	number:=0

	jumlah:
		fmt.Printf("%d \n",number)	
	label1:
		fmt.Printf("data ke-")
		number++
	if number<10{
		goto jumlah 
	}else if number <20{
		goto label1
	}
}