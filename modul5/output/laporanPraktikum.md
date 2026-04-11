# <h1 align="center">Laporan Praktikum Modul 5 - ... </h1>

<p align="center">Zhafif Yusuf Al Amin - 109082500137</p>

## Unguided

### 1. [soal1]

#### soal1.go

```go
package main
import "fmt"
func main(){
	var x int
	fmt.Scan(&x)
	fibonacci(x)
}
func fibonacci(x int) {
	a := 0
	b := 1

	for i := 0; i <= x; i++ {
		fmt.Print(a, " ")
		sn := a + b
		a = b
		b = sn
	}
}
```

### Output Unguided :

##### Output

![Screenshot Output soal1](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul4/output/outputSoal1.png)
sebuah kode untuk mencari bilangan fibonacci dari suku ke n, suku ke n ini adalah inputan dari user. pertama tama seperti biasa kita buat struktur golang lalu deklarasi variabel x sebagai int lalu minta inputan user baru kita panggil fungsi rekursifnya. lalu dibawahnya kita buat struktur fungsi rekursif dengan nama fibonacci dengan x sebagai int. lanjut kita deklarasi variabel a dengan nilai 0 dan b dengan nilai 1 yang dimana fibonacci adalah penjumlahan suku pertama dengan suku kedua lalu jumlah dari suku pertama dan kedua akan digunakan untuk penjumlahan suku ketiga, dan setrusnya. disini kita buat perulangan i dengan nilai 0, jika i lebih kecil atau sama dengan x maka nilai i akan ditambah 1 setiap perulangan. kita outputkan a, nilai a diambil dari luar perulangan. lalu disini kita buat temp untuk hasil penjumlahan a+b lalu kita ganti nilai dari a diambil dari b dan b diambil dari temp sn tadi. dan begitu seterusnya sampai i bernilai false, bernilai false ketika i lebih besar dari x.

### 2. [soal2]

#### soal2.go

```go
package main
import "fmt"
func main(){
	var n int
	fmt.Scan(&n)
	bintang(n)
}
func bintang(n int){
	for i:=1;i<=n;i++{
		for j:=1;j<=i;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}
```

### Output Unguided :

##### Output

![Screenshot Output soal1](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul4/output/outputSoal1.png)
sebuah kodingan rekursif untuk menampilkan output berupa bintang berbentuk segitiga rata kiri, program diawali dengan package main lalu import fmt, func main dulu baru func untuk rekursifnya. didalam func main terdapat deklarasi variabel n sebagai integer lalu meminta inputan dari n dan akan menampilkan rekursif dari func dibawahnya, disini saya memakai variabel bintang. lalu di func bintangnya kita deklarasi n sebagai int juga kemudian membuat perulangan i dengan nilai 1, jika i lebih kecil atau sama dengan n maka nilai i akan ditambah 1 disetiap perulangan. didalam perulangan i teredapat prulangan j dengan nilai 1, jika j lebih kecil atau sama dengan n maka nilai j akan ditambah 1 disetiap perulangan. didalam perulangan j terdapat fungsi untuk menampilkan output berupa string \*, inilah yang berfungsi untuk ditampilkan pada output. didalam perulangan i ada fungsi println yang berfungsi untuk membuat output dari j akan berbaris baru disetiap perulangannya.

### 3. [soal3.go]

#### soal3.go

```go
package main
import "fmt"
func main(){
	var n int
	fmt.Scan(&n)
	faktor(n)
}
func faktor(n int){
	for i:=1;i<=n;i++{
		if n%i==0{
			fmt.Print(i," ")
		}
	}
}
```

### Output Unguided :

##### Output

![Screenshot Output soal2](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul4/output/outputSoal2.png)
sebuah porgram rekursif untuk menampilkan faktor dari n, yang dimana n merupakan inputan dari user. program diawali deklarasi package utama lalu import liblary fmt dan func main. kita deklarasi variabel n sebagai integer lalu meminta inputannya dan akan ditampilkan menggunakan func rekursif yang bernama faktor. deklarasi variabel yang akan digunakan yaitu n dengan tipe data int. buat perulangan dengan nilai i adalah 1 jika i lebih kecil atau sama dengan n jumlah i akan di tambah 1 setiap perulangan. lalu kita buat percabangan yang menentukan jika n dimodulus i sama dengan 0 maka akan menampikan nilai dari i yang bisa me modulus n sampai bernilai 0 atau nama lainnya adalah bilangan yang menjadi faktor dari N

### 4. [soal4]

#### soal4.go

```go
package main
import "fmt"
func main(){
	var n int
	fmt.Scan(&n)
	ret(n)
}
func ret(n int){
	for i:=n;i>=1;i--{
		fmt.Print(i," ")
	}
	for i:=2;i<=n;i++{
		fmt.Print(i," ")
	}
}


```

### Output Unguided :

##### Output

![Screenshot Output modul2A](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul3/output/outputSoal3.png)
program ini digunakan untuk menampilkan pola bilangan turun lalu naik berdasarkan inputan user, dimana program diawali dengan deklarasi package main dan import fmt, kemudian pada func main dibuat variabel n bertipe integer untuk menampung inputan user yang selanjutnya diproses oleh fungsi ret, pada fungsi tersebut terdapat dua perulangan, perulangan pertama digunakan untuk mencetak bilangan dari n sampai 1 secara menurun dengan cara nilai i dimulai dari n lalu dikurangi hingga 1, kemudian perulangan kedua digunakan untuk mencetak bilangan naik kembali dari 2 sampai n agar tidak terjadi pengulangan angka 1 di tengah, sehingga jika kedua perulangan digabungkan akan menghasilkan pola bilangan yang turun lalu naik secara berurutan.

### 5. [soal5]

#### soal5.go

```go
package main
import "fmt"
func main(){
	var n int
	fmt.Scan(&n)
	ret(n)
}
func ret(n int){
	for i:=n;i>=1;i--{
		fmt.Print(i," ")
	}
	for i:=2;i<=n;i++{
		fmt.Print(i," ")
	}
}


```

### Output Unguided :

##### Output

![Screenshot Output modul2A](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul3/output/outputSoal3.png)
program ini digunakan untuk menampilkan bilangan ganjil dari 1 sampai n berdasarkan inputan user, dimana program diawali dengan deklarasi package main dan import library fmt yang berfungsi untuk input dan output, kemudian pada func main dibuat variabel n bertipe integer yang digunakan untuk menyimpan nilai yang dimasukkan oleh user melalui fmt.Scan(&n), setelah itu nilai n akan diproses oleh fungsi ganjil, pada fungsi tersebut dibuat perulangan dengan variabel i dimulai dari 1 hingga n dan akan terus bertambah 1 setiap perulangan, di dalam perulangan terdapat percabangan menggunakan operator modulus (%) untuk mengecek apakah suatu bilangan termasuk ganjil atau tidak, dimana jika hasil dari i%2 tidak sama dengan 0 maka bilangan tersebut adalah bilangan ganjil, sehingga nilai i akan dicetak menggunakan fmt.Print, proses ini akan berlangsung secara berulang sampai kondisi perulangan tidak terpenuhi yaitu ketika i lebih besar dari n, sehingga seluruh bilangan ganjil dari 1 sampai n berhasil ditampilkan.

### 6. [soal6]

#### soal6.go

```go
package main
import "fmt"
func main(){
	var n int
	fmt.Scan(&n)
	ret(n)
}
func ret(n int){
	for i:=n;i>=1;i--{
		fmt.Print(i," ")
	}
	for i:=2;i<=n;i++{
		fmt.Print(i," ")
	}
}


```

### Output Unguided :

##### Output

![Screenshot Output modul2A](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul3/output/outputSoal3.png)
program ini digunakan untuk menghitung hasil perpangkatan dari suatu bilangan x dengan pangkat y berdasarkan inputan user, dimana program diawali dengan deklarasi package main dan import library fmt yang berfungsi untuk input dan output, kemudian pada func main dibuat dua variabel yaitu x dan y bertipe integer yang digunakan untuk menyimpan nilai yang dimasukkan oleh user melalui fmt.Scan(&x,&y), setelah itu nilai tersebut akan diproses oleh fungsi pangkat, pada fungsi tersebut dibuat variabel hasil dengan nilai awal 1 yang berfungsi sebagai penampung hasil perkalian, kemudian dilakukan perulangan dengan variabel i dimulai dari 1 hingga y, dimana setiap perulangan nilai hasil akan dikalikan dengan x sehingga proses ini sama dengan mengalikan x sebanyak y kali, setelah perulangan selesai maka nilai hasil akan berisi hasil dari x pangkat y dan ditampilkan menggunakan fmt.Print.
