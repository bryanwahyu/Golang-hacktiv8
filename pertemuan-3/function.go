package main

import (
	"fmt"	
	"strings"
)


func main(){
	getnama("Bryan")
	getnama("Alfin")
	nama:="Budi"
	getnama(nama)
	umur:=getumur(1994)
	fmt.Printf("Umur saya adalah %d\n",umur)
	tambah,kurang:=tambahkurang(5,1)
	fmt.Printf("Nilai tambah=%d,Nilai kurang=%d\n",tambah,kurang)
	bagi(15,0)
	bagi(15,3)
	fmt.Println("================================================")
	rata:=hitungrata(3,4,5,6,7,8,9,3,9,1)
	fmt.Printf("Hitung rata-rata menjadi =%.2f\n",rata)
	
	angka:=[]int{2,3,4,5,5,5,8,7,8,9}
	rerata:=hitungrata(angka...)
	msg:=fmt.Sprintf("hitung rata-rata menjadi %2.f",rerata)
	fmt.Println(msg)
	fmt.Println("==================================")
	hobby("Bryan","makan","tidur","koding","ngajar")
}
func hobby(nama string,hoby ...string){
	hobb:=strings.Join(hoby,",")
	fmt.Println("halo nama saya",nama)
	fmt.Println("dan hobby saya",hobb)

}
func hitungrata(angkas ...int) float64{
	total:=0
	
	for _,elemen:=range angkas{
		total +=elemen
	}
	jumlah_bil:=len(angkas)

	avg:=float64(total) / float64(jumlah_bil)

	return avg
}
func getnama(nama string){
	fmt.Printf("Halo nama saya %s\n",nama)
}
func getumur(tahun int) int{
	return 2021-tahun 
}
func tambahkurang(a int,b int) (jumlah_tambah int,jumlah_kurang int){
	jumlah_tambah=a+b
	jumlah_kurang=a-b
	
	return jumlah_tambah,jumlah_kurang	

}
func bagi(a int,b int){
	if b==0{
		fmt.Printf("Maaf paramater salah\n")
		return 
	}
	fmt.Printf("hasil pembagian %d\n",a/b)
}