# <h1 align="center">Laporan Praktikum Modul 4 - ... </h1>
<p align="center">Zhafif Yusuf Al Amin - 109082500137</p>

## Unguided 

### 1. [soal1]
#### soal1.go

```go
package main

import "fmt"

func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := n; i >= 1; i-- {
		*hasil *= i
	}
}
func permutasi(n, r int, hasil *int) {
	var fn, fnr int
	faktorial(n, &fn)
	faktorial(n-r, &fnr)
	*hasil = fn / fnr
}
func combinasi(n, r int, hasil *int) {
	var fn, fr, fnr int
	faktorial(n, &fn)
	faktorial(r, &fr)
	faktorial(n-r, &fnr)
	*hasil = fn / (fr * fnr)
}
func main() {
	var a, b, c, d int
	var p1, k1, p2, k2 int

	fmt.Scan(&a, &b, &c, &d)
	permutasi(a, c, &p1)
	combinasi(a, c, &k1)
	permutasi(b, d, &p2)
	combinasi(b, d, &k2)
	fmt.Println(p1, k1)
	fmt.Println(p2, k2)
}
```
### Output Unguided :

##### Output 
![Screenshot Output soal1](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul4/output/outputSoal1.png)
kodingan ini berfungsi untuk mencari hasil permutasi dan kombinasi dari 4 buah inputan bilangan bulat. pertama kita buat prosedur untuk mencari faktorial dari buah inputan dengan cara menggunakan perulangan for yang dimana sebuah bilangan akan di kurangi 1 setiap perulangan sampai nilai perulangannya bernilai false. nilai dari setiap perulangan akan dikalikan dan disimpan dalam variabel hasil yang nantinya akan digunakan untuk melihat hasil faktorialnya. lalu prosedur permutasi digunakan untuk menghitung banyaknya susunan yang mungkin dari sejumlah elemen dengan memperhatikan urutan sesuai rumus pada soal. hasil dari peermutasi akan disimpan di vaibael hasil. lalu ada prosedur combinasi yang digunakan untuk menghitung banyaknya cara memilih sejumlah elemen tanpa memperhatikan urutan. dengan isi prosedur terdapat prosedur faktorial prosedur ini memanggil prosedur faktorial untuk menghitung nilai faktorial yang dibutuhkan dalam rumus kombinasi. Prosedur faktorial digunakan sebanyak tiga kali, yaitu untuk menghitung nilai n!, r!, dan (n - r)! yang kemudian dimasukkan ke dalam rumus kombinasi. dan yang terakhir adalah fungsi utamanya yaitu func main yang dimana ini adalah tempat untuk mengeksesuki dari prosedur prosedur yang ada mita panggil seua disini. kita buat variabel baru yaitu a,b,c,d sebagai int lalu ada p1, k1, p2, dan k2 untuk menyimpan hasil perhitungan permutasi dan kombinasi. Hasil dari setiap perhitungan disimpan ke dalam variabel tujuan menggunakan pointer sehingga pada fungsi utama tidak perlu menuliskan banyak rumus dengan dampak prosedur prosedur tadi bisa berguna tidak hanyak untuk fungsi utama tapi bisa di fungsi fungsi yang lain.
### 2. [soal2.go]
#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output soal2](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul4/output/outputSoal2.png)
sebuah kode yang berfungsi untuk mencari pemenang dari daftar peserta yang diberikan. pertama kita buat prosedur hitungSkor dengan nama variabel soal dan skor menggunakan pointer int lalu kita beri nilai 0 di setiap variabelnya. buat perulangan yang bisa menerima 8 inputan karena pada soal diberikan ketentuan menjawab 8 soal. buat percabangan yang berfungsi untuk menentukan jika waktu kuran dari atau sam dengan 300 detik maka soal di anggap berhasil kalo lebih dari 300 detik makan soal ga akan di hitung. jika berhasil maka nilai dari soal akan ditambah 1. skor akan dijumlahkan dengan waktu pengerjaan. lalu pada fungsi utama, pada fungsi main, program digunakan untuk menginput nama peserta dan menentukan pemenang. variabel nama, pemenang, soal, dan skor digunakan untuk menyimpan data, sedangkan maxSoal dan minWaktu untuk membandingkan hasil terbaik. Program melakukan perulangan untuk membaca nama sampai input "Selesai". setiap peserta akan dihitung jumlah soal dan total waktunya menggunakan prosedur hitungSkor. Pemenang ditentukan dari jumlah soal terbanyak, dan jika sama maka dipilih yang memiliki waktu tercepat. Terakhir, program menampilkan nama pemenang beserta hasilnya.

### 3. [soal3]
#### soal3.go

```go
package main
import "fmt"
func cetakDeret	(n int){
	for{
		fmt.Print(n)
		if n==1{
			break
		}
		fmt.Print(" ")
		if n%2==0{
			n=n/2
		}else{
			n=3*n+1
		}
	}
}
func main() {
	var n int
	fmt.Scan(&n)

	cetakDeret(n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output modul2A](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul3/output/outputSoal3.png)
kode ini berfungsi untuk menampilkan deret bilangan Collatz dari sebuah input. Pertama kita buat prosedur cetakDeret yang menerima parameter n sebagai nilai awal. Di dalamnya digunakan perulangan terus menerus untuk mencetak nilai n. Jika n sudah sama dengan 1 maka perulangan dihentikan. Setiap angka dipisahkan dengan spasi. Lalu dibuat percabangan, jika n genap maka dibagi 2, sedangkan jika ganjil maka dihitung dengan rumus 3n + 1. Pada fungsi main, program menerima input sebuah bilangan lalu memanggil prosedur cetakDeret untuk mencetak deret sampai bernilai 1.