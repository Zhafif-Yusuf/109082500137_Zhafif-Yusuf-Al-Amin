# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>
<p align="center">Zhafif Yusuf Al Amin - 109082500137</p>

## Unguided 

### 1. [modul2A]
#### modul2A.go

```go
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
![Screenshot Output modul2A](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul2/output/outputModul2A.png)
sebuah program bahasa Go yang dimulai dengan package main yang menandakan program utama. Lalu import "fmt" digunakan untuk memanggil package fmt agar bisa menampilkan output ke layar. Fungsi func main(){} menjadi titik awal eksekusi program. Di dalamnya dideklarasikan variabel satu, dua, tiga, dan temp dengan tipe data string. Fungsi fmt.Println pertama menampilkan nilai awal dari variabel satu, dua, dan tiga. Selanjutnya dilakukan proses pertukaran nilai menggunakan variabel sementara (temp), yaitu temp = satu, satu = dua, dua = tiga, dan tiga = temp, sehingga nilai setiap variabel berubah dari posisi awalnya. Terakhir, fmt.Println kedua menampilkan output akhir setelah nilai variabel berhasil ditukar.

### 2. [modul2B]
#### modul2B.go

```go
package main

import "fmt"

func main() {
	var warna1, warna2, warna3, warna4 string
	urutan := true
	for i := 1; i <= 5; i++ {
		fmt.Print("Percobaan ", i, " : ")
		fmt.Scan(&warna1, &warna2, &warna3, &warna4)
		if warna1 == "merah" && warna2 == "kuning" && warna3 == "hijau" && warna4 == "ungu" {
			urutan = urutan
		} else {
			urutan = false
		}

	}
	fmt.Print(urutan)

}

```
### Output Unguided :

##### Output 
![Screenshot Output modul2A](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul2/output/outputModul2B.png)
package main menandakan program utama, ibaratnya kita akan membangun rumah dan package main ini seperti tembok rumah yang menjadi dasar program. import "fmt" berfungsi memanggil package fmt untuk input dan output, ibarat saat membangun rumah seperti semen, pasir, batu, dan air yang membantu proses pembangunan. func main(){…} ini bagaikan pintu rumah, karena semua program dijalankan dari sini, rumah tidak akan bisa dimasuki tanpa pintu. Lalu var warna1, warna2, warna3, warna4 string digunakan untuk mendeklarasikan variabel yang akan menampung input warna dari pengguna, serta urutan := true sebagai penanda awal bahwa urutan warna dianggap benar. Setelah itu ada for i := 1; i <= 5; i++ yang berfungsi melakukan perulangan sebanyak 5 kali untuk percobaan memasukkan warna, kemudian fmt.Print menampilkan teks percobaan dan fmt.Scan digunakan untuk menerima input warna dari pengguna. Selanjutnya program melakukan pengecekan menggunakan if apakah urutan warna yang dimasukkan adalah merah, kuning, hijau, ungu, jika sesuai maka nilai urutan tetap, tetapi jika tidak sesuai maka urutan akan berubah menjadi false. Terakhir fmt.Print(urutan) menampilkan hasil akhir apakah urutan warna yang dimasukkan selama percobaan benar atau tidak.

### 2. [modul2C]
#### modul2C.go

```go
package main

import "fmt"

func main() {
	var parsel int
	var biayag int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&parsel)
	berat := parsel / 1000
	gram := parsel % 1000

	biayak := berat * 10000
	if gram >= 500 {
		biayag = gram * 5
	} else if gram < 500 {
		biayag = gram * 15
	} 
	 if berat>10{
		biayag = gram*0
	}
	total := biayak + biayag
	
	fmt.Println("Detail berat: " ,berat, " kg + ", gram, " gr ")
	fmt.Println(gram)
	fmt.Println(biayag)
	fmt.Println(total)

}
```
### Output Unguided :

##### Output 
![Screenshot Output modul2A](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul2/output/outputModul2C.png)
package main menandakan program utama, ibaratnya kita akan membangun rumah dan package main ini seperti tembok rumah yang menjadi dasar dari program. import "fmt" berfungsi memanggil package fmt yang digunakan untuk melakukan input dan output, ibarat saat membangun rumah seperti semen, pasir, batu, dan air yang membantu proses pembangunan. func main(){…} ini bagaikan pintu rumah, karena semua program dijalankan dari sini dan tanpa fungsi main program tidak bisa berjalan. Lalu var parsel int dan var biayag int digunakan untuk mendeklarasikan variabel bertipe integer untuk menyimpan berat parsel yang diinput dan biaya tambahan gram. fmt.Print digunakan untuk menampilkan teks “Berat parsel (gram):” lalu fmt.Scan(&parsel) berfungsi menerima input berat parsel dari pengguna. Setelah itu berat := parsel / 1000 digunakan untuk mengubah berat parsel menjadi kilogram, sedangkan gram := parsel % 1000 digunakan untuk mengambil sisa gramnya. Kemudian biayak := berat * 10000 menghitung biaya berdasarkan berat kilogram. Selanjutnya dilakukan pengecekan menggunakan if, jika sisa gram ≥ 500 maka biaya gram dihitung gram * 5, jika < 500 maka dihitung gram * 15, namun jika berat lebih dari 10 kg maka biaya gram menjadi 0. Setelah itu total := biayak + biayag menghitung total biaya pengiriman. Terakhir beberapa fmt.Println digunakan untuk menampilkan detail berat parsel, sisa gram, biaya gram, dan total biaya yang harus dibayar.