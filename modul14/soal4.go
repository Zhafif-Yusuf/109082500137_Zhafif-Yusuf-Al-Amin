package main

import "fmt"

const MAX = 1000

type DataAngka [MAX]int

func insertionSort(data *DataAngka, jumlah int) {
	var posisi, geser, sementara int

	for posisi = 1; posisi < jumlah; posisi++ {
		sementara = data[posisi]
		geser = posisi

		for geser > 0 && sementara < data[geser-1] {
			data[geser] = data[geser-1]
			geser--
		}

		data[geser] = sementara
	}
}

func main() {
	var data DataAngka
	var jumlahData int
	var angka int
	var i int

	jumlahData = 0

	for {
		fmt.Scan(&angka)

		if angka < 0 {
			break
		}

		data[jumlahData] = angka
		jumlahData++
	}

	insertionSort(&data, jumlahData)

	for i = 0; i < jumlahData; i++ {
		fmt.Print(data[i])

		if i < jumlahData-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	if jumlahData <= 2 {
		fmt.Println("Data berjarak 0")
		return
	}

	selisih := data[1] - data[0]
	jarakTetap := true

	for i = 2; i < jumlahData; i++ {
		if data[i]-data[i-1] != selisih {
			jarakTetap = false
			break
		}
	}

	if jarakTetap {
		fmt.Println("Data berjarak", selisih)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}