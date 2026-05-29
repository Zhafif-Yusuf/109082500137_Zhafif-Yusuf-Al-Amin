# <h1 align="center">Laporan Praktikum Modul 10 - ... </h1>

<p align="center">Zhafif Yusuf Al Amin - 109082500137</p>

## Unguided

### 1. [soal1]

#### soal1.go

```go
package main
import "fmt"
func main() {
	var N int
	var berat [1000]float64

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&berat[i])
	}

	min := berat[0]
	max := berat[0]

	for i := 1; i < N; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}

	fmt.Println(min, max)
}
```

### Output Unguided :

##### Output

![Screenshot Output soal1](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul10/output/outputSoal1.png)
sebuah program untuk mencari berat kelinci paling kecil dan paling besar dari sejumlah data berat kelinci yang dimasukkan pengguna. pertama program membaca nilai n yang menyatakan banyaknya data berat kelinci, kemudian menyimpan seluruh data berat tersebut ke dalam array berat berkapasitas 1000. setelah semua data tersimpan, nilai minimum (min) dan maksimum (max) diinisialisasi dengan elemen pertama array, yaitu berat[0]. selanjutnya program melakukan perulangan dari data kedua hingga data terakhir untuk membandingkan setiap berat dengan nilai minimum dan maksimum saat ini. jika ditemukan berat yang lebih kecil dari min, maka nilai min diperbarui, dan jika ditemukan berat yang lebih besar dari max, maka nilai max juga diperbarui. setelah seluruh data selesai diperiksa, program menampilkan berat kelinci terkecil dan terbesar yang ditemukan dalam array.


### 2. [soal2]

#### soal2.go

```go
package main
import "fmt"
func main() {
	var x, y int
	var berat [1000]float64
	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	jumlahWadah := 0
	totalSemua := 0.0
	for i := 0; i < x; i += y {
		total := 0.0

		for j := i; j < i+y && j < x; j++ {
			total += berat[j]
		}

		fmt.Printf("%.2f ", total)
		totalSemua += total
		jumlahWadah++
	}
	fmt.Println()
	fmt.Printf("%.2f\n", totalSemua/float64(jumlahWadah))
}
```

### Output Unguided :

##### Output

![Screenshot Output soal1](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul10/output/outputSoal2.png)
menghitung total berat ikan pada setiap wadah dan mencari rata-rata berat per wadah. pertama, program membaca dua bilangan bulat, yaitu x yang menyatakan jumlah ikan dan y yang menyatakan kapasitas atau banyaknya ikan yang dimasukkan ke setiap wadah. selanjutnya, program membaca x data berat ikan dan menyimpannya ke dalam array berat berkapasitas 1000. setelah itu, program melakukan perulangan dengan langkah sebesar y untuk mengelompokkan ikan ke dalam beberapa wadah sesuai urutan input. pada setiap wadah, program menjumlahkan berat ikan yang berada dalam kelompok tersebut dan langsung menampilkannya. nilai total berat setiap wadah juga ditambahkan ke variabel totalsemua, sedangkan jumlahwadah digunakan untuk menghitung banyaknya wadah yang terbentuk. setelah seluruh data selesai diproses, program menghitung rata-rata berat per wadah dengan membagi totalsemua dengan jumlahwadah, kemudian menampilkan hasilnya sebagai keluaran pada baris terakhir.



### 3. [soal3.go]

#### soal3.go

```go
package main
import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int) (float64, float64) {
	min := arrBerat[0]
	max := arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < min {
			min = arrBerat[i]
		}
		if arrBerat[i] > max {
			max = arrBerat[i]
		}
	}

	return min, max
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64

	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}

	return total / float64(n)
}

func main() {
	var data arrBalita
	var n int

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d : ", i+1)
		fmt.Scan(&data[i])
	}

	min, max := hitungMinMax(data, n)
	rata := rerata(data, n)

	fmt.Printf("\nBerat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}
```

### Output Unguided :

##### Output

![Screenshot Output soal2](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul10/output/outputSoal3.png)
program untuk mengolah data berat balita yang disimpan dalam sebuah array. pengguna terlebih dahulu memasukkan jumlah data balita, kemudian memasukkan berat masing-masing balita satu per satu. setelah seluruh data tersimpan, program menggunakan fungsi hitungminmax untuk mencari berat balita paling kecil dan paling besar dengan membandingkan setiap data dalam array. selanjutnya, fungsi rerata digunakan untuk menghitung rata-rata berat balita dengan menjumlahkan seluruh data berat dan membaginya dengan banyaknya data yang dimasukkan. pada akhir program, ditampilkan berat minimum, berat maksimum, dan rata-rata berat balita berdasarkan data yang telah diinputkan pengguna.



