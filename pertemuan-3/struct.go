package main

import "fmt"
type pengarang struct{
	nama string
} 
type buku struct {
	id      uint
	nama    string
	halaman uint8
	penulis pengarang
}

//type pengarang=penulis

func main() {
	
	var s1 = []struct {
		penulis pengarang
	}{	{
			penulis:pengarang{nama:"Agus"},
		},
		{
			penulis:pengarang{nama:"Maudy"},
		},
	}
	for _,el:= range s1 {
		fmt.Println("penulis",el.penulis.nama)
	}

	fmt.Println("=============================")
	penulis2:=pengarang{nama:"Agus"}
	buku1:=buku{}
	buku1.id = 2
	buku1.penulis.nama = "Bryan Wahyu"
	buku1.nama="buku golang untuk pemula"
	buku1.halaman = 20
	buku2:=buku{nama:"Ayat-ayat Cinta"}
	buku2.penulis=penulis2
	fmt.Println("ini buku ke-",buku1.id)
	fmt.Println("ini nama buku",buku1.nama)
	fmt.Println("ini penulis bernama",buku1.penulis.nama)
	fmt.Println("jumlah halaman buku ini adalah",buku1.halaman)
	fmt.Println("buku ke-3 yang berjudul",buku2.nama,"yang berhalaman",buku2.halaman,"pengarang dibuat oleh ",buku2.penulis.nama)

	fmt.Println("=====================================================")
	var alamatbuku1 *buku =&buku1
	alamatbuku1.nama="Buku baru"
	fmt.Println("ini nama buku",buku1.nama)
	fmt.Println("ini nama buku dengan referensi buku1",alamatbuku1.nama)

	fmt.Println("=======================================================")
	kumpulanbuku:=[]buku{{nama:"Buku1",halaman:99},{nama:"Buku2",halaman:95}}
	for _,el:=range kumpulanbuku{
		fmt.Println("judul buku",el.nama)
		fmt.Println("halaman sebuah buku",el.halaman)
	}
	fmt.Println("=========================================================")
		var student struct{
			nama string `nama_murid`
			kelas int8
		} 
		student.nama="Bryan"
		student.kelas=13

		fmt.Println("nama:",student.nama)
		fmt.Println("kelas:",student.kelas)
}