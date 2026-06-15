package main

import "fmt"

const MAX = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [MAX]Buku

func DaftarkanBuku(dataBuku *DaftarBuku, jumlah int) {
	var i int

	for i = 0; i < jumlah; i++ {
		fmt.Scan(
			&dataBuku[i].id,
			&dataBuku[i].judul,
			&dataBuku[i].penulis,
			&dataBuku[i].penerbit,
			&dataBuku[i].eksemplar,
			&dataBuku[i].tahun,
			&dataBuku[i].rating,
		)
	}
}

func CetakTerfavorit(dataBuku DaftarBuku, jumlah int) {
	var posisiTerbaik int

	for i := 1; i < jumlah; i++ {
		if dataBuku[i].rating > dataBuku[posisiTerbaik].rating {
			posisiTerbaik = i
		}
	}

	fmt.Println("=== BUKU TERFAVORIT ===")
	fmt.Println("Judul     :", dataBuku[posisiTerbaik].judul)
	fmt.Println("Penulis   :", dataBuku[posisiTerbaik].penulis)
	fmt.Println("Penerbit  :", dataBuku[posisiTerbaik].penerbit)
	fmt.Println("Tahun     :", dataBuku[posisiTerbaik].tahun)
	fmt.Println("Rating    :", dataBuku[posisiTerbaik].rating)
	fmt.Println()
}

func UrutBuku(dataBuku *DaftarBuku, jumlah int) {
	var i, j int
	var sementara Buku

	for i = 1; i < jumlah; i++ {
		sementara = dataBuku[i]
		j = i

		for j > 0 && sementara.rating > dataBuku[j-1].rating {
			dataBuku[j] = dataBuku[j-1]
			j--
		}

		dataBuku[j] = sementara
	}
}

func Cetak5Terbaru(dataBuku DaftarBuku, jumlah int) {
	fmt.Println("=== 5 BUKU DENGAN RATING TERTINGGI ===")

	batas := 5
	if jumlah < 5 {
		batas = jumlah
	}

	for i := 0; i < batas; i++ {
		fmt.Printf("%d. %s (Rating %d)\n",
			i+1,
			dataBuku[i].judul,
			dataBuku[i].rating)
	}

	fmt.Println()
}

func CariBuku(dataBuku DaftarBuku, jumlah int, ratingCari int) {
	var kiri, kanan, tengah int
	var ketemu bool

	kiri = 0
	kanan = jumlah - 1

	for kiri <= kanan && !ketemu {
		tengah = (kiri + kanan) / 2

		if dataBuku[tengah].rating == ratingCari {
			ketemu = true

			fmt.Println("=== HASIL PENCARIAN ===")
			fmt.Println("Judul      :", dataBuku[tengah].judul)
			fmt.Println("Penulis    :", dataBuku[tengah].penulis)
			fmt.Println("Penerbit   :", dataBuku[tengah].penerbit)
			fmt.Println("Tahun      :", dataBuku[tengah].tahun)
			fmt.Println("Eksemplar  :", dataBuku[tengah].eksemplar)
			fmt.Println("Rating     :", dataBuku[tengah].rating)

		} else if ratingCari > dataBuku[tengah].rating {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if !ketemu {
		fmt.Println("=== HASIL PENCARIAN ===")
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var dataBuku DaftarBuku
	var jumlahBuku int
	var ratingCari int

	fmt.Scan(&jumlahBuku)

	DaftarkanBuku(&dataBuku, jumlahBuku)

	CetakTerfavorit(dataBuku, jumlahBuku)

	UrutBuku(&dataBuku, jumlahBuku)

	Cetak5Terbaru(dataBuku, jumlahBuku)

	fmt.Scan(&ratingCari)

	CariBuku(dataBuku, jumlahBuku, ratingCari)
}