package main

import "fmt"

func main(){
	sum:=0

	for index:=0;index<10;index++{
		if index==5{
			continue
		}
		sum += index
		fmt.Printf(" nilai index=%d dan sum=%d\n",index,sum)
	}
	for sum<100{
		fmt.Printf("nilai sum_sebelum=%d ",sum)

		sum+=sum
		fmt.Printf("nilai sesudah sum=%d\n",sum)
	}

}