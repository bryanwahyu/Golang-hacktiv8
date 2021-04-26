package main

import "fmt"


func main(){
	a:=8
	if a>5{
		if a==10{
			fmt.Printf("Nilai sempurna")
		}else if a>7 && a<10{
			fmt.Printf("Nilai bagus")
		}else{
			fmt.Printf("nilai cukup")
		}
	}else{
		a=7+4
		a=2-9
		fmt.Printf("nilai tidak bagus")
	}
	
}