
package main

import "fmt"

func main(){
	index:=7
	switch{
	case index==1 :
		fmt.Printf("Ini index ==1\n")
	case index==2 :
		fmt.Printf("Ini index ==2\n")
	case index==3 :
		fmt.Printf("Ini index ==3\n")
	case index==4 :
		fmt.Printf("Ini index ==4\n")
	case index==5 :
		fmt.Printf("Ini index ==5\n")
	case index==6 :
		fmt.Printf("Ini index ==6\n")
	case index==7 :
		fmt.Printf("Ini index ==7\n")
		fallthrough
	case (index<10) && (index>7):
		fmt.Printf("Ini index berada 8 atau 9\n")
	default :{
				fmt.Printf("index tidak berada di angka 1-9\n")
				fmt.Printf("coba lagi yaa..")
			 }
	}
}