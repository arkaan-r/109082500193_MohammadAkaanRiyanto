# <h1 align="center">Laporan Praktikum Modul 5 - ... </h1>
<p align="center">[Mohammad Arkaan Riyanto] - [109082500193]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func fibonacci(n int) int{
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i <= n; i++{
		fmt.Print(fibonacci(i), " ")
	}
}
```
### Output Unguided :	

##### Output 
![Screenshot Output Unguided 1_](https://github.com/arkaan-r/109082500193_MohammadAkaanRiyanto/blob/main/modul5/output/output-soal1.png)
[Penjelasan]Program ini buat nampilin deret Fibonacci sampai angka ke-n. Fungsi fibonacci pakai cara rekursif, jadi dia manggil dirinya sendiri. Kalau n = 0 hasilnya 0, kalau n = 1 hasilnya 1, selain itu dihitung dari penjumlahan dua angka sebelumnya (n-1 dan n-2). Di main, user masukin nilai n, lalu program ngeloop dari 0 sampai n dan tiap angka dihitung pakai fungsi tadi, terus langsung ditampilin. Jadi outputnya deret Fibonacci dari awal sampai ke-n.

//

### 2. [Soal]
#### soal2.go

```go
package main

import "fmt"

func bintang(n int, i int) {
	if i > n {
		return
	}
	for j := 0; j < i; j++ {
		fmt.Print("*")
	}
	fmt.Println()
	bintang(n, i+1)
}

func main() {
	var n int
	fmt.Scan(&n)

	bintang(n, 1)
}

```
### Output Unguided :	

##### Output 
![Screenshot Output Unguided 1_](https://github.com/arkaan-r/109082500193_MohammadAkaanRiyanto/blob/main/modul5/output/output-soal2.png)
[Penjelasan] Program ini buat nampilin pola bintang segitiga ke bawah yaitu naik dari 1 sampai n. Fungsi bintang pakai rekursif. Jadi kalau i udah lebih dari n, fungsi berhenti. Kalau belum, dia cetak bintang sebanyak i, lalu pindah baris, terus manggil dirinya lagi dengan i+1. Di main, user masukin nilai n, lalu fungsi dipanggil mulai dari i = 1. Jadi hasilnya bakal kayak: 1 bintang, 2 bintang, sampai n bintang.
//

### 3. [Soal]
#### soal2.go

```go
package main

import "fmt"

func faktor(n int, i int) {
	if i > n{
		return
	}
	if n%i == 0 {
	fmt.Print(i, " ")

	}
	faktor(n, i+1)
}

func main() {
	var n int
	fmt.Scan(&n)

	faktor(n, 1)
}


```
### Output Unguided :	

##### Output 
![Screenshot Output Unguided 1_](https://github.com/arkaan-r/109082500193_MohammadAkaanRiyanto/blob/main/modul5/output/output-soal3.png)
[Penjelasan]Program ini buat nampilin semua faktor dari suatu bilangan. Fungsi faktor jalan dari i = 1 sampai n. Kalau i lebih dari n, fungsi berhenti. Setiap langkah dicek, kalau n % i == 0 berarti i adalah faktor dari n, jadi langsung dicetak. Setelah itu fungsi manggil dirinya lagi dengan i+1 buat lanjut ngecek angka berikutnya. Di main, user masukin angka n, lalu program nampilin semua faktor dari angka tersebut mulai dari 1 sampai n.