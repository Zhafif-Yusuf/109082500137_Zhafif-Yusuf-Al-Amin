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