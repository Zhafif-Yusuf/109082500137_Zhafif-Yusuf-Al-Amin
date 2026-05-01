package main
import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("masukkan jumlah data: ")
	fmt.Scan(&n)

	var data [100]int 
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	fmt.Print("isi array: ")
	for i := 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	fmt.Print("indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	fmt.Print("indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Print("masukkan x: ")
	fmt.Scan(&x)

	fmt.Print("kelipatan indeks ", x, ": ")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(data[i], " ")
		}
	}
	fmt.Println()

	var hapus int
	fmt.Print("hapus indeks ke: ")
	fmt.Scan(&hapus)

	for i := hapus; i < n-1; i++ {
		data[i] = data[i+1]
	}
	n = n - 1

	fmt.Print("setelah dihapus: ")
	for i := 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	jumlah := 0
	for i := 0; i < n; i++ {
		jumlah += data[i]
	}
	rata := float64(jumlah) / float64(n)
	fmt.Println("rata-rata:", rata)

	total := 0.0
	for i := 0; i < n; i++ {
		selisih := float64(data[i]) - rata
		total += selisih * selisih
	}
	std := math.Sqrt(total / float64(n))
	fmt.Println("standar deviasi:", std)

	var cari int
	fmt.Print("cari angka: ")
	fmt.Scan(&cari)

	jml := 0
	for i := 0; i < n; i++ {
		if data[i] == cari {
			jml++
		}
	}
	fmt.Println("frekuensi:", jml)
}