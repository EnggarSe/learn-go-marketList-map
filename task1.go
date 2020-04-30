package main

import "fmt"

func main() {
	namaBarang := make(map[int]string)
	hargaBarang := make(map[int]int)
	stokBarang := make(map[int]int)

	namaBarang[0] = ("Coca-cola")
	namaBarang[1] = ("Fanta")
	namaBarang[2] = ("Sprite")

	hargaBarang[0] = (10000)
	hargaBarang[1] = (15000)
	hargaBarang[2] = (20000)

	stokBarang[0] = (9)
	stokBarang[1] = (11)
	stokBarang[2] = (5)

	for index, element := range stokBarang {
		if element < 10 {
			fmt.Println("-------------------------")
			fmt.Println("Nama Barang : ", namaBarang[index])
			fmt.Println("Harga Barang :", hargaBarang[index])
			fmt.Println("Stok Barang :", stokBarang[index])
		}

	}
}
