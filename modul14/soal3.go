package main

import "fmt"

const MAX = 1000000

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
	var median int

	jumlahData = 0

	for {
		fmt.Scan(&angka)

		if angka == -5313 {
			break
		}

		if angka == 0 {
			insertionSort(&data, jumlahData)

			if jumlahData%2 == 1 {
				median = data[jumlahData/2]
			} else {
				median = (data[jumlahData/2-1] + data[jumlahData/2]) / 2
			}

			fmt.Println(median)
		} else {
			data[jumlahData] = angka
			jumlahData++
		}
	}
}