# <h1 align="center">REVIEW ALGORITMA DAN PEMROGRAMAN 1 - ... </h1>
<p align="center">[Mohammad Arkaan Riyanto] - [109082500193]</p>

## Unguided 

### Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N. Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).
#### 

```go
package main

import "fmt"

func faktor(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Printf("%d ", i)
	}
	faktor(n, i+1)
}
func main() {
	var n int
	fmt.Scan(&n)
	faktor(n, 1)
	fmt.Println()
}

```
### Output Unguided :

##### Output 

![Screenshot Output Unguided 1_1](https://github.com/arkaan-r/109082500193_MohammadAkaanRiyanto/blob/main/LAPRAKSMT-2/output/soal3.png)
[ Program ini bertujuan untuk menampilkan semua faktor dari suatu bilangan n yang diinput oleh pengguna dengan menggunakan pendekatan rekursif. Fungsi faktor(n, i) bekerja dengan memeriksa setiap bilangan i mulai dari 1 hingga n, di mana jika i lebih besar dari n maka proses dihentikan (basis rekursi). Pada setiap langkah, program mengecek apakah n habis dibagi i (n % i == 0), dan jika iya maka nilai i dicetak sebagai faktor dari n. Setelah itu, fungsi memanggil dirinya sendiri dengan nilai i+1 untuk melanjutkan pengecekan hingga seluruh kemungkinan faktor diperiksa. Dengan cara ini, semua faktor dari n akan ditampilkan secara berurutan.]
