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