package main

import "fmt"

type Buku struct {
	id       string
	judul    string
	penulis  string
	penerbit string
	eksemplar int
	tahun    int
	rating   int
}

func insertionSort(buku []Buku) {
	n := len(buku)

	for i := 1; i < n; i++ {
		temp := buku[i]
		j := i - 1

		for j >= 0 && buku[j].rating < temp.rating {
			buku[j+1] = buku[j]
			j--
		}

		buku[j+1] = temp
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	data := make([]Buku, n)

	for i := 0; i < n; i++ {
		fmt.Scan(
			&data[i].id,
			&data[i].judul,
			&data[i].penulis,
			&data[i].penerbit,
			&data[i].eksemplar,
			&data[i].tahun,
			&data[i].rating,
		)
	}

	insertionSort(data)

	fmt.Println("Buku Terfavorit:")
	fmt.Println(data[0].judul, data[0].penulis, data[0].penerbit, data[0].tahun)

	fmt.Println("5 Buku Rating Tertinggi:")

	limit := 5
	if n < 5 {
		limit = n
	}

	for i := 0; i < limit; i++ {
		fmt.Println(data[i].judul)
	}

	var cari int
	fmt.Scan(&cari)

	found := false

	for i := 0; i < n; i++ {
		if data[i].rating == cari {
			fmt.Println(data[i])
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Tidak ada buku dengan rating tersebut")
	}
}