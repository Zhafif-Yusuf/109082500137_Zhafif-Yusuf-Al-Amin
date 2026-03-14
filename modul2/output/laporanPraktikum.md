# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">Zhafif Yusuf Al Amin - 109082500137</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

package main
import "fmt"
func main() {
	var (satu, dua, tiga string
		temp            string
)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :

##### Output 
![Screenshot Output modul2A]https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul2/output/outputModul2A.png
sebuah program bahasa Go yang dimulai dengan package main yang menandakan program utama. Lalu import "fmt" digunakan untuk memanggil package fmt agar bisa menampilkan output ke layar. Fungsi func main(){} menjadi titik awal eksekusi program. Di dalamnya dideklarasikan variabel satu, dua, tiga, dan temp dengan tipe data string. Fungsi fmt.Println pertama menampilkan nilai awal dari variabel satu, dua, dan tiga. Selanjutnya dilakukan proses pertukaran nilai menggunakan variabel sementara (temp), yaitu temp = satu, satu = dua, dua = tiga, dan tiga = temp, sehingga nilai setiap variabel berubah dari posisi awalnya. Terakhir, fmt.Println kedua menampilkan output akhir setelah nilai variabel berhasil ditukar.
