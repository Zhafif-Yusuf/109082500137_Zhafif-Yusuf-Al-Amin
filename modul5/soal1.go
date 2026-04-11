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