package main

import "fmt"

func main(){
	var array = [...]int{1,2,6,4}
	fmt.Println(len(array))
	fmt.Println(array)
	var array2d=[4][4]int{{1,3,4,6},{2,6,8,9},{3,8,8,8},{9,9,9,0}}
	fmt.Println(array2d)

	for i:=0;i<len(array);i++{
		fmt.Printf("elemen ke-%d:%d\n",i,array[i])
	}
	for _,elemen:=range array{
		fmt.Printf("elemen:%d\n",elemen)
	} 
//	buah:=make([]string,1)
//	buah[0]="Apel"
//	buah=append(buah, "Mangga")
	//fmt.Println(buah)
	
	var	slicebuah=[]string{"mangga","apel","pisang","rambutan","melon"}
		fmt.Println(slicebuah)
		buaha:=slicebuah[1:]
		buahb:=slicebuah[1:4]

		buahbb:=buahb[0:1]
		buahaa:=buaha[0:1]

		fmt.Println("lenght panjang sebuah slice /array ",len(buahb))
		fmt.Println("Caps panjang maksimal sebuah slice /array ",cap(buahb))
		
		fmt.Println(buaha)
		fmt.Println(buahb)
		fmt.Println(buahaa)
		fmt.Println(buahbb)
		fmt.Println("-----------------------")		
		buahaa[0]="nanas"

		fmt.Println(slicebuah)
		fmt.Println(buaha)
		fmt.Println(buahb)
		fmt.Println(buahaa)
		fmt.Println(buahbb)
		fmt.Println("-------------------------")
		 
		buahbaru:=append(buaha,"pepaya")
		fmt.Println(buaha)
		fmt.Println(buahbaru)
		fmt.Println("--------------------------")

		copybuaha:=make([]string,2)
		n:=copy(copybuaha,buaha)

		fmt.Println(copybuaha)
		fmt.Println(n)
		
		fmt.Println("----------------------------")

		buahslicing3index:=slicebuah[0:2:2]
		buahslicing2index:=slicebuah[0:2]

		fmt.Println("len=",len(buahslicing3index))
		fmt.Println("cap=",cap(buahslicing3index))

		fmt.Println("--------------------------------")
		fmt.Println("len=",len(buahslicing2index))
		fmt.Println("cap=",cap(buahslicing2index))

		fmt.Println("===============================")
	
		  map1:= map[string]int{}
		 
		 
		  map1["januari"]=200
		  map1["Febuari"]=450
		
		   
		  fmt.Println(map1["januari"])

		  for k,v:=range map1{
			  fmt.Println(k,":",v)
		  }
		  delete(map1,"januari")
		  fmt.Println("=============================")
		  for k,v:=range map1{
			fmt.Println(k,":",v)
		}
		fmt.Println("===================")
		var value, isexits =map1["Febuari"]
		if isexits{
			fmt.Println(value)
		}else{
			fmt.Println("maaf data tidak ada")
		}
		fmt.Println("====================================")

		map3:=[]map[string]string{
			{
			"bulan":"Januari",
			"harga":"400",
			},
			{
			"bulan":"Febuari",
			"harga":"900",
			},
		}
		
		for _,val:=range map3{
			fmt.Println(val["bulan"],val["harga"])
		}

}