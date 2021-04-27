package main

import "fmt"
import "runtime"
func print(jumlah int,data string){
	for i:=0;i<=jumlah;i++{
		fmt.Println(i+1,data)
	}
}
func main(){
	runtime.GOMAXPROCS(3)
	go	print(5,"halo qaqa")
	 print(5,"apa kabar bro?")
	var input string
	fmt.Scanln(&input)

}