package main

import "fmt"

func main() {
	var warna1, warna2, warna3, warna4 string
	urutan := true
	for i := 1; i <= 5; i++ {
		fmt.Print("Percobaan ", i, " : ")
		fmt.Scan(&warna1, &warna2, &warna3, &warna4)
		if warna1 == "merah" && warna2 == "kuning" && warna3 == "hijau" && warna4 == "ungu" {
			urutan = urutan
		} else {
			urutan = false
		}

	}
	fmt.Print(urutan)

}
