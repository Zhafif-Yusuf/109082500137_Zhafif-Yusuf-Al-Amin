package main
import "fmt"
func main(){
	var n int
	fmt.Scan(&n)
	pro(n)
}
func pro(n int){
	for i:=1;i<=n;i=i+2{
		if n%2!=0 {
			fmt.Println(i)
			
		}else if n%2==0 {
			fmt.Print(i," ")
			
		}
	}
}