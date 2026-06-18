package main

import "fmt"

const NMAX = 100

type daftarKlub [NMAX]string

func main() {
	var klubA, klubB string
	var pemenang daftarKlub
	var skorA, skorB int
	var n, i int

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)

	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	n = 0
	i = 1

	for {
		fmt.Print("Pertandingan ", i, " : ")
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang[n] = klubA
		} else if skorB > skorA {
			pemenang[n] = klubB
		} else {
			pemenang[n] = "Draw"
		}

		n++
		i++
	}

	for i = 0; i < n; i++ {
		fmt.Println("Hasil", i+1, ":", pemenang[i])
	}

	fmt.Println("Pertandingan selesai")
}