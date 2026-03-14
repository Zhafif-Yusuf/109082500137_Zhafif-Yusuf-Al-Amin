package main

import "fmt"

func main() {
	var parsel int
	var biayag int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&parsel)
	berat := parsel / 1000
	gram := parsel % 1000

	biayak := berat * 10000
	if gram >= 500 {
		biayag = gram * 5
	} else if gram < 500 {
		biayag = gram * 15
	} 
	 if berat>10{
		biayag = gram*0
	}
	total := biayak + biayag
	
	fmt.Println("Detail berat: " ,berat, " kg + ", gram, " gr ")
	fmt.Println(gram)
	fmt.Println(biayag)
	fmt.Println(total)

}
