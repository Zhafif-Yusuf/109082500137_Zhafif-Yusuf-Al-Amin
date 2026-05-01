# <h1 align="center">Laporan Praktikum Modul 5 - ... </h1>

<p align="center">Zhafif Yusuf Al Amin - 109082500137</p>

## Unguided

### 1. [soal1]

#### soal1.go

```go
package main
import "fmt"

type Titik struct {
	x, y int
}

type Lingkaran struct {
	pusat Titik
	r     int
}

func dalamLingkaran(t Titik, L Lingkaran) bool {
	d := (t.x-L.pusat.x)*(t.x-L.pusat.x) + (t.y-L.pusat.y)*(t.y-L.pusat.y)
	return d <= L.r*L.r
}

func main() {
	var L1, L2 Lingkaran
	var t Titik
	fmt.Scan(&L1.pusat.x, &L1.pusat.y, &L1.r)
	fmt.Scan(&L2.pusat.x, &L2.pusat.y, &L2.r)
	fmt.Scan(&t.x, &t.y)
	dalam1 := dalamLingkaran(t, L1)
	dalam2 := dalamLingkaran(t, L2)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```

### Output Unguided :

##### Output

![Screenshot Output soal1](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul9/output/outputSoal1.png)
sebuah program untuk menentukan posisi suatu titik terhadap dua buah lingkaran, dimana program diawali dengan deklarasi package main lalu import fmt. kemudian kita membuat tipe data struct yaitu titik untuk menyimpan koordinat x dan y, dan lingkaran yang berisi pusat berupa titik dan jari-jari r. setelah itu kita membuat sebuah fungsi bernama dalamlingkaran yang berfungsi untuk mengecek apakah suatu titik berada di dalam lingkaran atau tidak, dengan cara menghitung jarak antara titik dengan pusat lingkaran menggunakan rumus jarak, namun disini tidak menggunakan akar melainkan langsung bentuk kuadrat agar lebih sederhana. jika hasil perhitungan jarak tersebut lebih kecil atau sama dengan kuadrat jari-jari maka titik dinyatakan berada di dalam lingkaran. selanjutnya pada fungsi main kita deklarasikan dua lingkaran dan satu titik lalu meminta inputan dari user untuk masing-masing lingkaran dan titik tersebut. setelah itu kita panggil fungsi dalamlingkaran untuk masing-masing lingkaran dan hasilnya disimpan dalam variabel boolean. terakhir kita menggunakan percabangan if else untuk menentukan output, jika kedua kondisi bernilai true maka titik berada di dalam kedua lingkaran, jika hanya salah satu yang true maka titik berada di salah satu lingkaran, dan jika keduanya false maka titik berada di luar kedua lingkaran.

### 2. [soal2]

#### soal2.go

```go
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
```

### Output Unguided :

##### Output

![Screenshot Output soal1](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul9/output/outputSoal2.png)
sebuah program untuk mengolah data array yang berisi bilangan bulat, dimana program diawali dengan deklarasi package main lalu import fmt dan math, dimana math digunakan untuk perhitungan akar pada standar deviasi. pertama tama program meminta input jumlah data n dari user, kemudian dibuat array dengan kapasitas maksimal 100 elemen dan dilakukan perulangan untuk mengisi array sebanyak n data sesuai input user, setelah itu seluruh isi array ditampilkan menggunakan perulangan. selanjutnya program menampilkan elemen dengan indeks ganjil dimulai dari indeks 1 dan bertambah 2 setiap perulangan serta elemen dengan indeks genap dimulai dari indeks 0 dan bertambah 2, kemudian program meminta input nilai x untuk menampilkan elemen dengan indeks kelipatan x dengan cara mengecek apakah indeks habis dibagi x. berikutnya program meminta input indeks yang ingin dihapus lalu elemen array digeser ke kiri mulai dari indeks tersebut sehingga elemen setelahnya menggantikan posisi sebelumnya dan jumlah data n dikurangi satu, kemudian array ditampilkan kembali. setelah itu program menghitung rata-rata dengan menjumlahkan seluruh elemen array lalu dibagi jumlah data, dilanjutkan dengan menghitung standar deviasi dengan cara mencari selisih tiap elemen terhadap rata-rata, dikuadratkan, dijumlahkan, lalu diakar menggunakan fungsi sqrt. terakhir program meminta input sebuah angka untuk dicari frekuensinya dan dilakukan perulangan untuk menghitung jumlah kemunculan angka tersebut di dalam array, kemudian hasilnya ditampilkan sebagai frekuensi.



### 3. [soal3.go]

#### soal3.go

```go
package main
import "fmt"

func main() {
	var klubA, klubB string
	fmt.Print("klub a: ")
	fmt.Scan(&klubA)
	fmt.Print("klub b: ")
	fmt.Scan(&klubB)

	var skorA, skorB int
	var hasil [100]string
	var pemenang [100]string
	var n int = 0

	for i := 1; ; i++ {
		fmt.Print("pertandingan ", i, ": ")
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil[n] = klubA
			pemenang[n] = klubA
		} else if skorB > skorA {
			hasil[n] = klubB
			pemenang[n] = klubB
		} else {
			hasil[n] = "draw"
		}
		n = n + 1
	}
	for j := 0; j < n; j++ {
		fmt.Println("hasil", j+1, ":", hasil[j])
	}
	fmt.Println("pertandingan selesai")
}
```

### Output Unguided :

##### Output

![Screenshot Output soal2](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul9/output/outputSoal3.png)
sebuah program untuk merekap hasil pertandingan antara dua klub sepak bola, dimana program diawali dengan deklarasi package main lalu import fmt untuk input dan output. pertama tama program meminta input nama dua klub yaitu klubA dan klubB yang akan bertanding. kemudian dideklarasikan variabel skorA dan skorB untuk menyimpan skor masing-masing klub, serta array hasil dan pemenang untuk menyimpan hasil pertandingan, dan variabel n sebagai penghitung jumlah pertandingan yang valid. selanjutnya program menggunakan perulangan tanpa batas untuk memasukkan skor tiap pertandingan, dimana user akan terus memasukkan skor sampai salah satu atau kedua skor bernilai negatif, yang menandakan input berhenti. di dalam perulangan dilakukan percabangan untuk menentukan pemenang, jika skorA lebih besar maka klubA disimpan, jika skorB lebih besar maka klubB disimpan, dan jika sama maka hasilnya adalah draw. setiap hasil disimpan ke array dan nilai n ditambah satu. setelah proses input selesai, program menampilkan seluruh hasil pertandingan menggunakan perulangan dari indeks 0 sampai n-1 sesuai urutan pertandingan, lalu diakhiri dengan menampilkan pesan bahwa pertandingan telah selesai.

### 4. [soal4]

#### soal4.go

```go
package main
import "fmt"

const NMAX int = 127
type tabel [NMAX]rune
func isiArray(t *tabel, n *int) {
	var huruf rune
	*n = 0

	for {
		fmt.Scanf("%c", &huruf)
		if huruf == ' ' || huruf == '\n' {
			continue
		}
		if huruf == '.' || *n >= NMAX {
			break
		}
		t[*n] = huruf
		*n = *n + 1
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]), " ")
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	var temp rune

	for i := 0; i < n/2; i++ {
		temp = t[i]
		t[i] = t[n-1-i]
		t[n-1-i] = temp
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var data tabel
	var n int

	fmt.Print("teks: ")
	isiArray(&data, &n)

	if palindrom(data, n) {
		fmt.Println("palindrom ? true")
	} else {
		fmt.Println("palindrom ? false")
	}
	fmt.Print("reverse teks: ")
	balikanArray(&data, n)
	cetakArray(data, n)
}
```

### Output Unguided :

##### Output

![Screenshot Output modul2A](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul9/output/outputSoal4.png)
sebuah program untuk mengolah array karakter dan mengecek apakah suatu teks termasuk palindrom, dimana program diawali dengan deklarasi package main lalu import fmt. pertama tama dibuat konstanta NMAX sebagai batas maksimal array dan tipe data tabel berupa array rune untuk menyimpan karakter. kemudian terdapat prosedur isiArray yang digunakan untuk membaca input karakter satu per satu dari user sampai ditemukan tanda titik (.) sebagai penanda akhir atau jumlah data sudah mencapai batas maksimum, dimana spasi dan enter akan dilewati. setelah itu terdapat prosedur cetakArray untuk menampilkan isi array karakter ke layar menggunakan perulangan. selanjutnya terdapat prosedur balikanArray yang berfungsi untuk membalik urutan isi array dengan cara menukar elemen depan dan belakang secara bertahap hingga ke tengah. kemudian terdapat fungsi palindrom yang digunakan untuk mengecek apakah susunan karakter sama jika dibaca dari depan maupun belakang dengan cara membandingkan elemen dari kedua sisi, jika ada yang berbeda maka bernilai false, jika semua sama maka bernilai true. pada fungsi main, program meminta input teks dari user lalu memanggil fungsi palindrom untuk menentukan apakah teks tersebut termasuk palindrom atau tidak dan menampilkan hasilnya, kemudian array dibalik menggunakan prosedur balikanArray dan hasilnya ditampilkan menggunakan cetakArray.

