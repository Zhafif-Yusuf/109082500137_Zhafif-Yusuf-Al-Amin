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
![Screenshot Output Unguided 1_1](https://github.com/Zhafif-Yusuf/109082500137_Zhafif-Yusuf-Al-Amin/modul2/output/outputModul2A.png)
[penjelasan]
