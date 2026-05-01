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