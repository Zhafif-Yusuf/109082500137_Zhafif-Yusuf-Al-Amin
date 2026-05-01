package main
import "fmt"

type Titik struct {
	x, y int
}

type Lingkaran struct {
	pusat Titik
	r     int
}

func dalamLingkaran(t Titik, L Lingkaran) bool {
	d := (t.x-L.pusat.x)*(t.x-L.pusat.x) + (t.y-L.pusat.y)*(t.y-L.pusat.y)
	return d <= L.r*L.r
}

func main() {
	var L1, L2 Lingkaran
	var t Titik
	fmt.Scan(&L1.pusat.x, &L1.pusat.y, &L1.r)
	fmt.Scan(&L2.pusat.x, &L2.pusat.y, &L2.r)
	fmt.Scan(&t.x, &t.y)
	dalam1 := dalamLingkaran(t, L1)
	dalam2 := dalamLingkaran(t, L2)

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