# <h1 align="center">Laporan Praktikum Modul 14 - ... </h1>
<p align="center">Zhafif Yusuf Al Amin - 109082500137</p>

## Unguided 

### 1. [soal1]
#### soal1.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output soal1](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul14/output/outputSoal1.png)
Program ini digunakan untuk mengurutkan nomor rumah kerabat pada beberapa daerah secara membesar (ascending). Pertama, program membaca jumlah daerah yang akan diproses. Setelah itu, untuk setiap daerah program membaca banyaknya nomor rumah yang dimiliki kerabat dan menyimpannya ke dalam array. Setelah seluruh nomor rumah pada suatu daerah berhasil dibaca, program melakukan proses pengurutan menggunakan algoritma Selection Sort, yaitu dengan mencari nomor rumah yang nilainya paling kecil kemudian menempatkannya pada posisi yang seharusnya. Proses tersebut dilakukan berulang hingga seluruh nomor rumah terurut dari yang terkecil ke yang terbesar. Setelah pengurutan selesai, program menampilkan hasil urutan nomor rumah untuk setiap daerah dalam satu baris output. Langkah yang sama akan terus dilakukan sampai seluruh data daerah selesai diproses.



### 2. [modul2B]
#### modul2B.go

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output soal2](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul14/output/outputSoal2.png)
program ini digunakan untuk menampilkan nomor rumah kerabat dengan aturan nomor rumah ganjil dicetak terlebih dahulu secara terurut membesar, kemudian diikuti nomor rumah genap secara terurut mengecil. pertama, program membaca jumlah daerah yang akan diproses. setelah itu, untuk setiap daerah program membaca banyaknya nomor rumah dan memisahkan setiap nomor rumah ke dalam dua array berbeda, yaitu array ganjil untuk nomor rumah ganjil dan array genap untuk nomor rumah genap. setelah seluruh data pada suatu daerah selesai dibaca, kedua array tersebut diurutkan menggunakan algoritma selection sort melalui fungsi urutnaik(), sehingga data pada masing-masing array tersusun dari nilai terkecil ke terbesar. selanjutnya program mencetak seluruh nomor rumah ganjil dari awal hingga akhir sehingga menghasilkan urutan membesar, sedangkan nomor rumah genap dicetak dari indeks terakhir ke indeks pertama sehingga menghasilkan urutan mengecil. setelah semua nomor rumah pada daerah tersebut ditampilkan, program berpindah ke daerah berikutnya dan mengulangi proses yang sama hingga seluruh data selesai diproses.

### 3. [soal3]
#### soal3.go

```go
package main

import "fmt"

const MAX = 1000000

type DataAngka [MAX]int

func insertionSort(data *DataAngka, jumlah int) {
	var posisi, geser, sementara int

	for posisi = 1; posisi < jumlah; posisi++ {
		sementara = data[posisi]
		geser = posisi

		for geser > 0 && sementara < data[geser-1] {
			data[geser] = data[geser-1]
			geser--
		}

		data[geser] = sementara
	}
}

func main() {
	var data DataAngka
	var jumlahData int
	var angka int
	var median int

	jumlahData = 0

	for {
		fmt.Scan(&angka)

		if angka == -5313 {
			break
		}

		if angka == 0 {
			insertionSort(&data, jumlahData)

			if jumlahData%2 == 1 {
				median = data[jumlahData/2]
			} else {
				median = (data[jumlahData/2-1] + data[jumlahData/2]) / 2
			}

			fmt.Println(median)
		} else {
			data[jumlahData] = angka
			jumlahData++
		}
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output modul2A](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul14/output/outputSoal3.png)
program ini digunakan untuk mencari dan menampilkan nilai median dari sekumpulan data yang dimasukkan pengguna. program akan terus membaca bilangan satu per satu hingga menemukan nilai khusus -5313 sebagai penanda akhir input. setiap bilangan yang bukan 0 dan bukan -5313 akan disimpan ke dalam array. ketika program menemukan angka 0, seluruh data yang sudah tersimpan akan diurutkan terlebih dahulu menggunakan algoritma insertion sort secara membesar (ascending). setelah data terurut, program menghitung nilai median berdasarkan jumlah data yang ada. jika jumlah data ganjil, median diambil dari nilai yang berada tepat di tengah array, sedangkan jika jumlah data genap, median diperoleh dari rata-rata dua nilai tengah dan hasilnya dibulatkan ke bawah karena menggunakan tipe data integer. nilai median tersebut kemudian ditampilkan ke layar. proses pembacaan data akan terus berlangsung hingga program menemukan angka -5313, kemudian program berhenti dijalankan.

### 4. [soal4]
#### soal4.go

```go
package main

import "fmt"

const MAX = 1000

type DataAngka [MAX]int

func insertionSort(data *DataAngka, jumlah int) {
	var posisi, geser, sementara int

	for posisi = 1; posisi < jumlah; posisi++ {
		sementara = data[posisi]
		geser = posisi

		for geser > 0 && sementara < data[geser-1] {
			data[geser] = data[geser-1]
			geser--
		}

		data[geser] = sementara
	}
}

func main() {
	var data DataAngka
	var jumlahData int
	var angka int
	var i int

	jumlahData = 0

	for {
		fmt.Scan(&angka)

		if angka < 0 {
			break
		}

		data[jumlahData] = angka
		jumlahData++
	}

	insertionSort(&data, jumlahData)

	for i = 0; i < jumlahData; i++ {
		fmt.Print(data[i])

		if i < jumlahData-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	if jumlahData <= 2 {
		fmt.Println("Data berjarak 0")
		return
	}

	selisih := data[1] - data[0]
	jarakTetap := true

	for i = 2; i < jumlahData; i++ {
		if data[i]-data[i-1] != selisih {
			jarakTetap = false
			break
		}
	}

	if jarakTetap {
		fmt.Println("Data berjarak", selisih)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output modul2A](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul14/output/outputSoal4.png)
program ini digunakan untuk membaca sekumpulan bilangan bulat non-negatif, mengurutkannya menggunakan algoritma insertion sort, kemudian memeriksa apakah jarak antar setiap bilangan setelah diurutkan memiliki selisih yang tetap atau tidak. program akan terus menerima input bilangan dan menyimpannya ke dalam array hingga ditemukan bilangan negatif sebagai penanda akhir input. setelah seluruh data selesai dibaca, program mengurutkan data secara membesar (ascending) menggunakan algoritma insertion sort dengan cara menyisipkan setiap elemen ke posisi yang sesuai pada bagian array yang sudah terurut. selanjutnya program menampilkan seluruh data yang telah terurut. setelah itu, program menghitung selisih antara dua elemen pertama sebagai acuan dan membandingkannya dengan selisih setiap pasangan elemen berikutnya. jika seluruh selisih bernilai sama, maka program menampilkan pesan "data berjarak x", dengan x adalah nilai selisih tersebut. namun jika ditemukan selisih yang berbeda, program akan menampilkan pesan "data berjarak tidak tetap". jika jumlah data kurang dari atau sama dengan dua, program langsung menampilkan "data berjarak 0" karena belum cukup data untuk menentukan pola jarak yang tetap.

### 5. [soal5]
#### soal5.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output modul2A](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul14/output/outputSoal5.png)
program ini digunakan untuk mengelola data buku dalam sebuah perpustakaan. pertama, program membaca jumlah buku yang akan dimasukkan, kemudian setiap data buku yang terdiri dari id, judul, penulis, penerbit, jumlah eksemplar, tahun terbit, dan rating disimpan ke dalam array melalui prosedur daftarkanbuku. setelah seluruh data berhasil dibaca, program mencari buku dengan rating tertinggi melalui prosedur cetakterfavorit dan menampilkan informasi buku tersebut sebagai buku terfavorit. selanjutnya seluruh data buku diurutkan berdasarkan rating secara menurun (descending) menggunakan algoritma insertion sort pada prosedur urutbuku, sehingga buku dengan rating tertinggi berada di posisi awal array. setelah proses pengurutan selesai, prosedur cetak5terbaru menampilkan maksimal lima judul buku dengan rating tertinggi. kemudian program membaca sebuah nilai rating yang ingin dicari dan menjalankan prosedur caribuku yang menggunakan algoritma binary search untuk mencari buku dengan rating tersebut pada data yang sudah terurut. jika buku ditemukan, program akan menampilkan informasi lengkap buku tersebut, sedangkan jika tidak ditemukan program akan menampilkan pesan bahwa tidak ada buku dengan rating yang dicari. dengan demikian, program ini menggabungkan proses input data, pencarian nilai maksimum, pengurutan menggunakan insertion sort, dan pencarian menggunakan binary search untuk mengelola data buku secara lebih efisien.