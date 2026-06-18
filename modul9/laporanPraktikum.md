# <h1 align="center">Laporan Praktikum Modul 9 - ... </h1>
<p align="center">[Mohammad Arkaan Riyanto] - [109082500193]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import (
	"fmt"
	"math"
)

type titik struct {
	x int
	y int
}

type lingkaran struct {
	pusat  titik
	radius int
}

func jarak(p, q titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

func didalam(c lingkaran, p titik) bool {
	if jarak(c.pusat, p) <= float64(c.radius) {
		return true
	}
	return false
}

func main() {
	var l1, l2 lingkaran
	var p titik
	var d1, d2 bool

	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.radius)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.radius)
	fmt.Scan(&p.x, &p.y)

	d1 = didalam(l1, p)
	d2 = didalam(l2, p)

	if d1 && d2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if d1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if d2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :	

##### Output 
![Screenshot Output Unguided 1_](https://github.com/MuhammadJovan/109082500024_Muhammad-Jovan/blob/main/modul9/output/output-soal1.png)
[Penjelasan]Program ini dipakai untuk mengecek apakah suatu titik masuk ke lingkaran pertama, lingkaran kedua, keduanya, atau malah di luar. Caranya dengan menghitung jarak titik ke pusat lingkaran, lalu dibandingkan dengan nilai radius. Kalau jaraknya lebih kecil atau sama dengan radius, berarti titik tersebut ada di dalam lingkaran.

//

### 2. [Soal]
#### soal2.go

```go
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

```
### Output Unguided :	

##### Output 
![Screenshot Output Unguided 1_](https://github.com/MuhammadJovan/109082500024_Muhammad-Jovan/blob/main/modul9/output/output-soal2.png)
[Penjelasan] Program ini digunakan untuk menyimpan beberapa angka ke dalam array lalu menampilkan isi array sesuai kebutuhan. Data bisa ditampilkan semua, berdasarkan indeks ganjil, indeks genap, atau indeks kelipatan tertentu. Selain itu program ini juga bisa menghapus data, mencari rata", dan menghitung berapa kali suatu angka muncul di dalam array.
//

### 3. [Soal]
#### soal2.go

```go
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


```
### Output Unguided :	

##### Output 
![Screenshot Output Unguided 1_](https://github.com/MuhammadJovan/109082500024_Muhammad-Jovan/blob/main/modul9/output/output-soal3.png)
[Penjelasan]Program ini dipakai untuk mencatat hasil pertandingan dua klub. Setiap skor yang dimasukkan akan dibandingkan untuk menentukan klub mana yang menang. Nama klub yang menang disimpan ke dalam array, lalu di akhir program semua hasil pertandingan ditampilkan satu per satu.
//

### 4. [Soal]
#### soal2.go

```go
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


```
### Output Unguided :	

##### Output 
![Screenshot Output Unguided 1_](https://github.com/MuhammadJovan/109082500024_Muhammad-Jovan/blob/main/modul9/output/output-soal4.png)
[Penjelasan]Program ini digunakan untuk membalik urutan huruf yang disimpan di dalam array. Setelah dibalik, program mengecek apakah kata tersebut termasuk palindrom atau tidak. Jika huruf dari depan dan belakang sama, maka kata tersebut termasuk palindrom.