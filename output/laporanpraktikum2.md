# <h1 align="center">REVIEW ALGORITMA DAN PEMROGRAMAN 1 - ... </h1>
<p align="center">[Mohammad Arkaan Riyanto] - [109082500193]</p>

## Unguided 

### Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user.
#### 

```go
package main

import "fmt"

func sangbintang(n int) {
	if n == 0 {
		return
	}
	sangbintang(n - 1)
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}
func main() {
	var n int
	fmt.Scan(&n)
	sangbintang(n)
}


```
### Output Unguided :

##### Output 

![Screenshot Output Unguided 1_1](https://github.com/arkaan-r/109082500193_MohammadAkaanRiyanto/blob/main/LAPRAKSMT-2/output/soal2.png)
[Program ini bertujuan untuk menampilkan pola bintang berbentuk segitiga dengan jumlah baris sesuai nilai n yang diinput oleh pengguna. Fungsi sangbintang menggunakan konsep rekursi, di mana jika n sama dengan 0 maka fungsi berhenti (basis rekursi), sedangkan jika n lebih dari 0 maka fungsi akan memanggil dirinya sendiri dengan nilai n-1 terlebih dahulu. Setelah pemanggilan rekursif selesai, program mencetak n buah karakter "*" pada satu baris menggunakan perulangan, sehingga setiap baris memiliki jumlah bintang yang bertambah secara bertahap dari 1 hingga n. Dengan cara ini, pola yang dihasilkan akan berbentuk segitiga siku-siku yang tersusun dari atas ke bawah.]
