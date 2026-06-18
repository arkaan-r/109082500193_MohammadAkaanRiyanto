package main

import (
	"fmt"
	"math"
)

const NMAX = 100

func main() {
	var arr [NMAX]int
	var n, i int
	var x, hapus int
	var jumlah int
	var rata, sd float64
	var cari, frek int

	fmt.Print("Jumlah data: ")
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Print("Data ke-", i, ": ")
		fmt.Scan(&arr[i])
	}

	fmt.Println("Isi array:")
	for i = 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Println("Indeks ganjil:")
	for i = 0; i < n; i++ {
		if i%2 != 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	fmt.Println("Indeks genap:")
	for i = 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	fmt.Print("Kelipatan indeks: ")
	fmt.Scan(&x)

	fmt.Println("Isi pada indeks kelipatan", x)
	for i = 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	fmt.Print("Indeks yang dihapus: ")
	fmt.Scan(&hapus)

	for i = hapus; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	n--

	fmt.Println("Array setelah dihapus:")
	for i = 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	for i = 0; i < n; i++ {
		jumlah += arr[i]
	}

	rata = float64(jumlah) / float64(n)
	fmt.Println("Rata-rata =", rata)

	for i = 0; i < n; i++ {
		sd = sd + math.Pow(float64(arr[i])-rata, 2)
	}

	sd = math.Sqrt(sd / float64(n))
	fmt.Println("Standar deviasi =", sd)

	fmt.Print("Bilangan yang dicari: ")
	fmt.Scan(&cari)

	for i = 0; i < n; i++ {
		if arr[i] == cari {
			frek++
		}
	}

	fmt.Println("Frekuensi =", frek)
}