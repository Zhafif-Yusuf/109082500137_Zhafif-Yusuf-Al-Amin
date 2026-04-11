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
