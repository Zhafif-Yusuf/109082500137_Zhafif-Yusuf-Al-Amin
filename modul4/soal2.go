package main
import "fmt"
func hitungSkor(soal *int, skor *int) {
	var waktu int
	*soal = 0
	*skor = 0

	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)

		if waktu <= 300 {
			(*soal)++
			*skor += waktu
		}
	}
}
func main(){
	var nama, pemenang string
	var soal, skor int
	maxSoal := -1
	minWaktu := 999999

	for {
		fmt.Scan(&nama)

		if nama == "Selesai" {
			break
		}

		hitungSkor(&soal, &skor)

		if soal > maxSoal || (soal == maxSoal && skor < minWaktu) {
			maxSoal = soal
			minWaktu = skor
			pemenang = nama
		}
	}

	fmt.Println(pemenang, maxSoal, minWaktu)
}