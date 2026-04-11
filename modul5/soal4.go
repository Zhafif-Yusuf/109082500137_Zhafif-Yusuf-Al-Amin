package main
import "fmt"
func main(){
	var n int
	fmt.Scan(&n)
	ret(n)
}
func ret(n int){
	for i:=n;i>=1;i--{
		fmt.Print(i," ")
	}
	for i:=2;i<=n;i++{
		fmt.Print(i," ")
	}
}
