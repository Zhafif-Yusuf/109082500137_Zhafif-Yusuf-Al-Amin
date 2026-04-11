	package main
	import "fmt"
	func main(){
		var x, y int
		fmt.Scan(&x,&y)
		pangkat(x,y)

	}
	func pangkat(x int, y int){
		hasil:=1
		for i:=1;i<=y;i++{
			hasil *=x

		}
		fmt.Print(hasil)

	}