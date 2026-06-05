# <h1 align="center">Laporan Praktikum Modul 12 - ... </h1>

<p align="center">Zhafif Yusuf Al Amin - 109082500137</p>

## Unguided

### 1. [soal1]

#### soal1.go

```go
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

	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", totalValid)

	for i := 0; i < JUMLAHCALON; i++ {
		if hasilSuara[i] > 0 {
			fmt.Printf("%d: %d\n", daftarNomor[i], hasilSuara[i])
		}
	}
}
```

### Output Unguided :

##### Output

![Screenshot Output soal1](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul12/output/outputSoal1.png)
program untuk membaca dan menghitung hasil pemilihan ketua rt dengan 20 calon. pertama, program membuat daftar nomor calon dari 1 sampai 20, kemudian membaca setiap suara yang dimasukkan pengguna hingga ditemukan angka 0 sebagai tanda akhir input. setiap suara yang masuk dihitung sebagai suara masuk, lalu divalidasi menggunakan fungsi sequential search (caricalon) untuk mengecek apakah nomor yang dipilih termasuk dalam daftar calon yang valid. jika ditemukan, suara tersebut dihitung sebagai suara sah dan jumlah suara calon terkait ditambahkan ke dalam array hasilsuara. setelah semua data selesai dibaca, program menampilkan jumlah suara masuk, jumlah suara sah, serta daftar calon yang memperoleh suara beserta banyaknya suara yang didapatkan masing-masing. dengan kata lain, program ini berfungsi untuk memvalidasi suara dan merekap hasil perolehan suara setiap calon dalam pemilihan ketua rt.


### 2. [soal2]

#### soal2.go

```go
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
```

### Output Unguided :

##### Output

![Screenshot Output soal1](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul12/output/outputSoal2.png)
program  untuk membaca dan memvalidasi data suara dalam pemilihan ketua rt yang terdiri dari 20 calon. setiap suara yang dimasukkan akan diperiksa menggunakan algoritma sequential search melalui fungsi caricalon() untuk memastikan bahwa nomor yang dipilih berada dalam rentang calon yang valid, yaitu 1 sampai 20. seluruh suara yang terbaca dihitung sebagai suara masuk, sedangkan suara yang valid dihitung sebagai suara sah dan direkap pada array hasilsuara. setelah proses pembacaan selesai saat ditemukan angka 0, program mencari calon dengan jumlah suara terbanyak sebagai ketua rt dan calon dengan jumlah suara terbanyak berikutnya sebagai wakil ketua rt. terakhir, program menampilkan jumlah suara masuk, jumlah suara sah, serta nomor calon yang terpilih sebagai ketua dan wakil ketua rt berdasarkan hasil perolehan suara.



### 3. [soal3.go]

#### soal3.go

```go
package main
import "fmt"

const NMAX = 1000000
var data [NMAX]int
func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	var kiri int = 0
	var kanan int = n - 1
	var tengah int

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2

		if data[tengah] == k {
			return tengah
		} else if k < data[tengah] {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	return -1
}

func main() {
	var n, k int

	fmt.Scan(&n, &k)

	isiArray(n)

	hasil := posisi(n, k)

	if hasil == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(hasil)
	}
}
```

### Output Unguided :

##### Output

![Screenshot Output soal2](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul12/output/outputSoal3.png)
program ini digunakan untuk mencari posisi suatu bilangan k pada kumpulan data integer yang sudah terurut membesar. program terlebih dahulu membaca banyak data (n) dan nilai yang ingin dicari (k), kemudian mengisi array melalui prosedur isiarray(). proses pencarian dilakukan oleh fungsi posisi() menggunakan algoritma binary search, yaitu dengan membandingkan nilai k terhadap elemen tengah array. jika nilai tengah lebih besar dari k, pencarian dilanjutkan ke bagian kiri array, sedangkan jika lebih kecil, pencarian dilanjutkan ke bagian kanan array. proses ini terus dilakukan hingga data ditemukan atau rentang pencarian habis. jika data ditemukan, fungsi mengembalikan indeks atau posisi data tersebut dalam array, sedangkan jika tidak ditemukan fungsi mengembalikan -1. pada program utama, apabila hasil pencarian bernilai -1, maka ditampilkan "tidak ada", sedangkan jika ditemukan akan ditampilkan posisi data tersebut. karena data sudah terurut, penggunaan binary search membuat proses pencarian menjadi lebih cepat dan efisien dibandingkan sequential search.



