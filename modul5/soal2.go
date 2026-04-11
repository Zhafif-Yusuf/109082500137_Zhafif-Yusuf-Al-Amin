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