	package main

	import "fmt"

	const MAX = 1000

	type DataRumah [MAX]int

	func urutNaik(daftar *DataRumah, jumlah int) {
		var kiri, kanan, posisiKecil, tukar int

		for kiri = 0; kiri < jumlah-1; kiri++ {
			posisiKecil = kiri

			for kanan = kiri + 1; kanan < jumlah; kanan++ {
				if daftar[kanan] < daftar[posisiKecil] {
					posisiKecil = kanan
				}
			}

			tukar = daftar[kiri]
			daftar[kiri] = daftar[posisiKecil]
			daftar[posisiKecil] = tukar
		}
	}

	func main() {
		var banyakDaerah, banyakRumah int
		var daerah, rumah int
		var nomorRumah DataRumah

		fmt.Scan(&banyakDaerah)

		for daerah = 0; daerah < banyakDaerah; daerah++ {
			fmt.Scan(&banyakRumah)

			for rumah = 0; rumah < banyakRumah; rumah++ {
				fmt.Scan(&nomorRumah[rumah])
			}

			urutNaik(&nomorRumah, banyakRumah)

			for rumah = 0; rumah < banyakRumah; rumah++ {
				fmt.Print(nomorRumah[rumah])

				if rumah < banyakRumah-1 {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	}