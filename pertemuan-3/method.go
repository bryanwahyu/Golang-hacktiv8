package main

import "fmt"

type murid struct{
	nama string 
	kelas int8
}
 func (siswa murid) getnama() {
	fmt.Println("halo ini murid bernama",siswa.nama)

 }
 func (siswa murid) getkelas(){
	 fmt.Println("kelas :",siswa.kelas)
 }
 func (siswa *murid)gantinama(nama string){
	 fmt.Println("Nama telah diganti")
	 siswa.nama=nama
 }
 func (siswa *murid)setKelas(kls int8){
	 fmt.Println("kelas sudah diset")
		siswa.kelas=kls
	}
	
 func main(){
	 var s1=murid{nama: "Bryan",kelas:5}
	 s1.getkelas()
	 s1.getnama()
	
	 s1.gantinama("Ade")
	 s1.setKelas(7)
	s1.getnama()
	s1.getkelas()













}

