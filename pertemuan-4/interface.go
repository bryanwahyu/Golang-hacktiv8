package main

import (
	"fmt"
	"math"
)

type hitung2d interface{
	luas() float64
	keliling() float64
}
type hitung3d interface{
	volume() float64	
}
type hitung interface{
	hitung2d 
	hitung3d
}

type circle struct{
	diameter float64
} 
func (c *circle) setdiameter( d float64){
	c.diameter=d
}
func (c circle) carijari2()float64{
	return c.diameter/2
}
func (c circle) luas() float64{
	return math.Pi * math.Pow(c.carijari2(),2)
}
func (c circle) keliling() float64{
	return math.Pi *c.diameter
}
type persegi struct{
	sisi float64
} 
func (p *persegi) setpersegi(s float64){
	p.sisi=s
}
func (p persegi) luas() float64{
	return math.Pow(p.sisi,2)
}
func (p persegi) keliling() float64{
	return p.sisi *4
}
type kubus struct{
	sisi float64
}
func (k *kubus) volume() float64{
	return math.Pow(k.sisi,3)
}
func (k *kubus) luas() float64{
	return math.Pow(k.sisi,2)*6
}
func (k *kubus) keliling() float64{
	return k.sisi *12
}
func main(){
	 var bangunandatar hitung2d
	 var bangunaruang hitung
	 bangunandatar=persegi{sisi:2.0}
	 fmt.Println("luas persegi",bangunandatar.luas())
	 fmt.Println("keliling persegi",bangunandatar.keliling())

	 bangunandatar=circle{diameter: 14.0}
	 fmt.Println("luas lingkaran",bangunandatar.luas())
	 fmt.Println("keliling lingkaran",bangunandatar.keliling())
	 fmt.Println("===============================================")
	 bangunaruang=&kubus{sisi:4.0}
	 fmt.Println("volume kubus",bangunaruang.volume())
	 fmt.Println("luas Kubus :",bangunandatar.luas())
	 fmt.Println("keliling kubus:",bangunaruang.keliling())
	 fmt.Println("========================================")
	var kosong interface{}
	kosong="Bryan"
	fmt.Println("isi interface kosong",kosong)
	kosong=[]string{"namaku","bryan"} 
	fmt.Println(kosong)
 }	

