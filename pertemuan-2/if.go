package main

import "fmt"
//import "math/rand"

func main(){
	if number:=9; number<10{
		a:=123
		fmt.Printf("jika bilangan random dibawah 10 bilangan random=%d\n",number)
		fmt.Printf(" coba block scope %d",a)
	}else if number==10{
		fmt.Printf("jika bilangan random sama persis 10 bilangan random=%d\n",number)
	}else{
		fmt.Printf("jika bilangan random diatas 10 bilangan random=%d\n",number)
	}
}