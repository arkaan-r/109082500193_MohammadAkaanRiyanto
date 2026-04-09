# <h1 align="center">REVIEW ALGORITMA DAN PEMROGRAMAN 1 - ... </h1>
<p align="center">[Mohammad Arkaan Riyanto] - [109082500193]</p>

## Unguided 

### Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan Sn = Sn−1 + Sn−2 . Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci tersebut.
#### 

```go
package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Print("n: ")
	for i := 0; i <= n; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	fmt.Print("Sn: ")
	for i := 0; i <= n; i++ {
		fmt.Print(fibonacci(i), " ")
	}

}

```
### Output Unguided :

##### Output 

![Screenshot Output Unguided 1_1](https://github.com/arkaan-r/109082500193_MohammadAkaanRiyanto/blob/main/LAPRAKSMT-2/output/soal1.png)
[Program ini bertujuan untuk menampilkan deret indeks dari 0 hingga n serta deret bilangan Fibonacci hingga suku ke-n dengan menggunakan bahasa pemrograman Go. Nilai n diinput oleh pengguna, kemudian program mencetak indeks menggunakan perulangan. Untuk menghitung bilangan Fibonacci, digunakan fungsi rekursif, di mana jika n ≤ 1 maka nilai dikembalikan langsung, sedangkan jika n > 1 maka dihitung sebagai penjumlahan dua suku sebelumnya, yaitu Fibonacci(n-1) + Fibonacci(n-2)]
