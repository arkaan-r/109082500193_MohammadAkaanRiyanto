# <h1 align="center">Laporan Praktikum Modul 4 - ... </h1>
<p align="center">[Mohammad Arkaan Riyanto] - [109082500193]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt" 
func factorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return	hasil
}

func permutation (n,r int) int {
	return factorial(n)/factorial(n-r)
}

func combination (n,r int) int {
	return factorial(n)/(factorial(r) * factorial(n-r))
}

func main() {
	var a,b,c,d int
	
	fmt.Scan(&a, &b, &c, &d)

	p1 := permutation(a,c)
	c1 := combination(a,c)

	p2 := permutation(b,d)
	c2 := combination(b,d)

	fmt.Println(p1, c1)
	fmt.Println(p2, c2)
}
```
### Output Unguided :	

##### Output 
![Screenshot Output Unguided 1_](https://github.com/MuhammadJovan/109082500024_Muhammad-Jovan/blob/main/modul4/output/output-soal1.png)
[Penjelasan]Program ini intinya buat ngitung permutasi sama kombinasi dari angka yang kita masukin. Jadi pertama ada fungsi faktorial yang dipakai buat nyari nilai faktorial (perkalian dari 1 sampai n). Nah dari situ dipakai lagi di fungsi permutasi buat hitung permutasi pakai rumus n!/(n−r)!, sama di fungsi kombinasi pakai rumus n!/(r!(n−r)!). Di bagian main, program minta 4 angka yaitu a, b, c, dan d. Terus dia hitung permutasi dan kombinasi dari (a, c) sama (b, d). Hasilnya ditampilin dalam dua baris, baris pertama buat a dan c, baris kedua buat b dan d. Jadi simpel aja, input angka  dihitung keluar hasil.


//

### 2. [Soal]
#### soal2.go

```go
package main

import	(
	"fmt"
)

func hitungSkor(soal *int, skor *int) {

	*soal = 0
	*skor = 0

	var waktu int
	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)
		
		if waktu <= 300 {
			*soal ++ 
			*skor += waktu
		}
	}
}	

func main() {
	var nama, pemenang string
	var maxSoal, minSkor int
	var soal, skor int

	maxSoal = -1
	minSkor = 999999

	for {
		fmt.Scan(&nama)
		if nama == "Selesai"{
		break
	}

	hitungSkor(&soal, &skor)

	if soal > maxSoal || (soal == maxSoal && skor < minSkor) {
		maxSoal = soal
		minSkor = skor
		pemenang = nama
		}
	}
	fmt.Println(pemenang, maxSoal, minSkor)
}

```
### Output Unguided :	

##### Output 
![Screenshot Output Unguided 1_](https://github.com/MuhammadJovan/109082500024_Muhammad-Jovan/blob/main/modul4/output/output-soal2.png)
[Penjelasan] Program ini buat nentuin pemenang dari beberapa peserta. Setiap peserta masukin nama, lalu 8 waktu pengerjaan soal. Kalau waktu ≤ 300 detik, berarti soalnya dihitung benar, jadi jumlah soal nambah dan waktunya masuk ke total skor. Program bakal terus jalan sampai input nama “Selesai”. Pemenangnya dipilih dari yang soal benarnya paling banyak, kalau sama maka yang waktunya paling kecil yang menang. Terakhir ditampilin nama pemenang, jumlah soal, dan total waktunya.

