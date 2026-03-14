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

### Output Unguided : dua tiga satu

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul2/output/outputModul2A.png)
package main: menandakan program utama ini ibaratnya kita akan membangun rumah & package main ini sebagai temboknya.

import "fmt": berfungsi memanggil package “fmt” ini ibarat bangun rumah sebagai semen, pasir, batu, air.

func main(){…}: ini bagaikan akses dengan filosofi pintu rumah, sebuah rumah tidak akan bisa di masuki tanpa pintu, masa iya rumah full tembok.

var (satu, dua, tiga string
temp            string)
:deklarasi variabel satu, dua, tiga dan temp menggunakan tipe data string.

fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
:fungsi untuk menampilkan sebuah string di barengi dengan variabel satu, dua dan tiga. Output dari variabel akan tampil sesuai inputan pertama








temp = satu
satu = dua
dua = tiga
tiga = temp
:sebuah fungsi untuk mengubah nilai dari satu variabel yang diambil dari variabel lain, disini menggunakan variabel temporary mengambil nilai dari variabel satu lalu satu mengambil nilai dari variabel dua dan seterusnya sehingga input pertama dan output akhir akan berbeda.

fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
:fungsi untuk menampilkan output akhir yang sudah kita ubah nilai masing masing variabelnya menggunakan variabel temporary.

