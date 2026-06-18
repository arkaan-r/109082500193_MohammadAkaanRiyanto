package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var huruf string
	*n = 0

	fmt.Print("Teks : ")

	for {
		fmt.Scan(&huruf)

		if huruf == "." || *n >= NMAX {
			break
		}

		t[*n] = rune(huruf[0])
		*n = *n + 1
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]), " ")
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	var temp rune

	for i := 0; i < n/2; i++ {
		temp = t[i]
		t[i] = t[n-1-i]
		t[n-1-i] = temp
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int
	var cek bool

	isiArray(&tab, &n)
	cek = palindrom(tab, n)
	balikanArray(&tab, n)

	fmt.Print("Reverse teks : ")
	cetakArray(tab, n)

	fmt.Println("Palindrom :", cek)
}