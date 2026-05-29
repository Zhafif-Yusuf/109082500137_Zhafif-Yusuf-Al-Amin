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