package main

import (
	"fmt"	
	"strings"
)
type fungsipanggil func(string) bool
func main(){
	getumur:=func(tahun int)int{
		return 2021-tahun
	}
	umur:=getumur(1994)

	fmt.Printf("umur saya %v\n",float64(umur))

	fmt.Println("==================================")

	//filter data minimal
	angka:=[]int{3,5,6,7,8,1,2,3,4,5,5}
	filtermin:=func(min int) []int{
		 var r []int
		 for _,element:=range angka{
			 if element<min{
				 continue
			 }
			 r = append(r, element)
		 }

		 return r
	}(4)
	fmt.Printf("bentuk bilangan awal = %v\n",angka)
	fmt.Printf("bentuk bilangan yang terfilter = %v\n", filtermin)
	fmt.Println("==================================================")
	max,getData:=findMax(angka,4)
	data1:=getData()
	fmt.Printf("panjang yang difilter %v dan bilangan yang terfilter x2 %v\n",max,data1)
	fmt.Println("============================================================")
	 var datahuruf=[]string{"Bryan","Adi","Hanif","Ado","Bambang"}
	 dataisin:=filterhuruf(datahuruf,func(s string) bool {
		 return strings.Contains(s,"a")
	 })
	 datapanjanghuruf5:=filterhuruf(datahuruf,func(s string) bool{
		 return len(s)==3
	 })
	 fmt.Printf("Ini data awal %v\n",datahuruf)
	 fmt.Printf("Ini data yang terfilter dengan huruf a %v\n",dataisin)
	 fmt.Printf("Ini data yang terfilter dengan panjang huruf 3  %v\n",datapanjanghuruf5)
}
func findMax(number []int,max int)(int,func() []int){
	 var res []int
	 for _,el:=range number{
		 if el<=max{
			res=append(res, el)
		 }
	}
	return len(res),func() []int {
		var resbaru []int
		for _,element:= range res{
			resbaru = append(resbaru, element*2)
		}
		return resbaru
	}

}

func filterhuruf(data []string,	callback fungsipanggil )[]string{
	var res []string
	for _,el:=range data{
		if filtered:=callback(el);filtered{
			res = append(res, el)
		} 
	}
	return res
}
