package main 

import "fmt"

// baris komentar

/*
	ini kalian komentar 
	komentar 
*/
func main(){
	 nama,nama_kedua:= "Bryan","Wahyu"
	 const Pi float64 =3.14
	 var angka uint16=999
	 var negatif int = -999
     var desimal =2.459
	 isMonday :=true 
	 
	fmt.Printf("Halo %s dan %s !\n",nama,nama_kedua)
	fmt.Printf("ini angka Positif %d dan Negatifnya %d \n",angka,negatif)
	fmt.Printf("ini desimal %.2f\n",desimal)
	fmt.Printf("ini hari Senin? %t\n",isMonday)
	fmt.Printf("--------------------------------\n")
	
	angka1:=2
	angka2:=4
	hasil:=angka1+angka2
	fmt.Printf("ini penjumlahan angka1 dan angka2 %d \n",hasil)
	hasil= angka1-angka2
	fmt.Printf("ini pengurangan angka1 dan angka2 %d\n",hasil)
	hasil= angka1*angka2
	fmt.Printf("ini perkalian angka1 dan angka2 %d\n",hasil)
	hasil= angka1/angka2
	fmt.Printf("ini Pembagian angka1 dan angka2 %d\n",hasil)
	hasil= angka1%angka2
	fmt.Printf("ini sisa bagi dari angka1 dan angka2 %d\n",hasil)
	var istrue bool
	istrue=(angka1==angka2)
	fmt.Printf("coba == (%d == %d) %t\n",angka1,angka2,istrue)
	istrue=(angka1!=angka2)
	fmt.Printf("coba != (%d != %d) %t\n",angka1,angka2,istrue)

	bool1:=true
	bool2:=false
	var hasil2 bool

	hasil2=bool1&&bool2
	fmt.Printf("&& (%t && %t) %t\n",bool1,bool2,hasil2)
	hasil2=bool1||bool2
	fmt.Printf("|| (%t || %t) %t\n",bool1,bool2,hasil2)	
	hasil2=!bool1
	fmt.Printf("! (%t) %t\n",bool1,hasil2)
}
