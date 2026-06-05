package main
import "fmt"

const JUMLAHCALON = 20
type dataCalon [JUMLAHCALON]int
func cariCalon(A dataCalon, n, nomor int) int {
	var posisi int = -1
	var i int = 0

	for i < n && posisi == -1 {
		if A[i] == nomor {
			posisi = i
		}
		i++
	}

	return posisi
}

func main() {
	var daftarNomor dataCalon
	var hasilSuara [JUMLAHCALON]int

	var pilihan int
	var totalMasuk int
	var totalValid int

	for i := 0; i < JUMLAHCALON; i++ {
		daftarNomor[i] = i + 1
	}

	for {
		fmt.Scan(&pilihan)

		if pilihan == 0 {
			break
		}

		totalMasuk++

		indeks := cariCalon(daftarNomor, JUMLAHCALON, pilihan)

		if indeks != -1 {
			totalValid++
			hasilSuara[indeks]++
		}
	}

	ketua := 0
	wakil := 1

	for i := 1; i < JUMLAHCALON; i++ {
		if hasilSuara[i] > hasilSuara[ketua] {
			wakil = ketua
			ketua = i
		} else if i != ketua && hasilSuara[i] > hasilSuara[wakil] {
			wakil = i
		}
	}

	if hasilSuara[wakil] > hasilSuara[ketua] {
		ketua, wakil = wakil, ketua
	}

	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", totalValid)
	fmt.Println("Ketua RT:", daftarNomor[ketua])
	fmt.Println("Wakil ketua:", daftarNomor[wakil])
}