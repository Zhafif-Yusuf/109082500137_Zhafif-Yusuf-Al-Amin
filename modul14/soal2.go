package main

import "fmt"

const MAX = 1000

type DataAngka [MAX]int

func urutNaik(data *DataAngka, jumlah int) {
	var kiri, kanan, posisiKecil, tukar int

	for kiri = 0; kiri < jumlah-1; kiri++ {
		posisiKecil = kiri

		for kanan = kiri + 1; kanan < jumlah; kanan++ {
			if data[kanan] < data[posisiKecil] {
				posisiKecil = kanan
			}
		}

		tukar = data[kiri]
		data[kiri] = data[posisiKecil]
		data[posisiKecil] = tukar
	}
}

func main() {
	var banyakDaerah, banyakRumah int
	var daerah, rumah, angka int

	var ganjil, genap DataAngka
	var jumlahGanjil, jumlahGenap int

	fmt.Scan(&banyakDaerah)

	for daerah = 0; daerah < banyakDaerah; daerah++ {
		fmt.Scan(&banyakRumah)

		jumlahGanjil = 0
		jumlahGenap = 0

		for rumah = 0; rumah < banyakRumah; rumah++ {
			fmt.Scan(&angka)

			if angka%2 == 1 {
				ganjil[jumlahGanjil] = angka
				jumlahGanjil++
			} else {
				genap[jumlahGenap] = angka
				jumlahGenap++
			}
		}

		urutNaik(&ganjil, jumlahGanjil)
		urutNaik(&genap, jumlahGenap)

		for rumah = 0; rumah < jumlahGanjil; rumah++ {
			fmt.Print(ganjil[rumah])

			if rumah < jumlahGanjil-1 || jumlahGenap > 0 {
				fmt.Print(" ")
			}
		}

		for rumah = jumlahGenap - 1; rumah >= 0; rumah-- {
			fmt.Print(genap[rumah])

			if rumah > 0 {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}