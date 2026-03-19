# <h1 align="center">Laporan Praktikum Modul 3 - ... </h1>
<p align="center">Zhafif Yusuf Al Amin - 109082500137</p>

## Unguided 

### 1. [soal1]
#### soal1.go

```go
package main
import "fmt"
func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	p1 := permutasi(a, c)
	k1 := kombinasi(a, c)
	p2 := permutasi(b, d)
	k2 := kombinasi(b, d)
	fmt.Println(p1, k1)
	fmt.Println(p2, k2)
}
```
### Output Unguided :

##### Output 
![Screenshot Output soal1](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul3/output/outputSoal1.png)
sintaks bahasa pemrograman go yang digunakan untuk mencari hasil permutasi dan kombinasi dari inputan, inputan meminta 4 angka dari 4 variabel berbeda, rumus dari faktorial,kombinasi dan permutasi dibuat di dalam fungsi fungsi berbeda. pada di fungsi utama (main) nanti kita hanya memanggil fungsi fungsi yang sudah kita buat tadi

### 2. [modul2B]
#### modul2B.go

```go
package main
import "fmt"
func f(x int) int {
	return x * x
}
func g(x int) int {
	return x - 2
}
func h(x int) int {
	return x + 1
}
func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	hasil1 := f(g(h(a))) 
	hasil2 := g(h(f(b))) 
	hasil3 := h(f(g(c)))
	fmt.Println(hasil1)
	fmt.Println(hasil2)
	fmt.Println(hasil3)
}

```
### Output Unguided :

##### Output 
![Screenshot Output soal2](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul3/output/outputSoal1.png)
sebuah kode untuk menghitung hasil fungsi komposisi yang menerima inputan 3 angka dari 3 variabel berbeda dan menampilkan output sesuai rumus yang sudah dibuat padafungsi tersendiri diluar fungsi utama. Fungsi f digunakan untuk menghitung kuadrat suatu bilangan, yaitu mengalikan nilai dengan dirinya sendiri (x²). Fungsi g digunakan untuk mengurangi suatu bilangan dengan nilai 2 (x − 2). Fungsi h digunakan untuk menambahkan suatu bilangan dengan nilai 1 (x + 1).

### 3. [soal3]
#### soal3.go

```go
package main
import (
	"fmt"
	"math"
)
func jarak(x, y, cx, cy int) float64 {
	return math.Sqrt(float64((x-cx)*(x-cx) + (y-cy)*(y-cy)))
}
func dalam(x, y, cx, cy, r int) bool {
	return jarak(x, y, cx, cy) <= float64(r)
}
func main() {
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var x, y int
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	dalam1 := dalam(x, y, cx1, cy1, r1)
	dalam2 := dalam(x, y, cx2, cy2, r2)
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
![Screenshot Output modul2A](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/blob/main/modul3/output/outputSoal3.png)
sebuah kode bahasa pemrograman go yang berfungsi untuk menentukan apakah suatu titik berada di dalam atau di luar dua lingkaran dilakukan dengan cara menghitung jarak antara titik tersebut dengan pusat masing-masing lingkaran. Program menerima input berupa koordinat titik pusat dan jari-jari dari dua lingkaran, serta satu titik sembarang yang akan diuji posisinya. Selanjutnya, program menghitung jarak antara titik tersebut dengan pusat lingkaran. fungsi jarak berguna untuk menghitung jarak antara suatu titik yang diuji adalah (x, y) dengan titik pusat lingkaran (cx, cy). funsgi dalam digunakan untuk menentukan apakah suatu titik berada di dalam sebuah lingkaran atau tidak dengan cara memanggil fungsi jarak lalu menghitung jarak titik ke pusat lingkaran dan membandingkan dengan radius. percabangan if else yang akan menentukan titik ada di mananya