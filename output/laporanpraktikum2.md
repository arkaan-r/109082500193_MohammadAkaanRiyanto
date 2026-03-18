# <h1 align="center">REVIEW ALGORITMA DAN PEMROGRAMAN 1 - ... </h1>
<p align="center">[Mohammad Arkaan Riyanto] - [109082500193]</p>

## Unguided 

### Diberikan tiga buah fungsi matematika yaitu f (x) = x2, g (x) = x − 2 dan h (x) = x + Fungsi komposisi (fogoh)(x) artinya adalah f(g(h(x))). Tuliskan f(x), g(x) dan h(x)dalam bentuk function.Masukan terdiri dari sebuah bilangan bulat a, b dan c yang dipisahkan oleh spasi.Keluaran terdiri dari tiga baris. Baris pertama adalah (fogoh)(a), baris kedua (gohof)(b),dan baris ketiga adalah (hofog)(c)!
#### 

```go
package main

import "fmt"

func f(x int) int {
	return x * x
}
func g(x int) int {
	return x - 2
}
func h(x int) int {
	return x + 1
}
func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	hasil1 := f(g(h(a)))
	hasil2 := g(h(f(b)))
	hasil3 := h(f(g(c)))

	fmt.Println(hasil1)
	fmt.Println(hasil2)
	fmt.Println(hasil3)
}

```
### Output Unguided :

##### Output 

![Screenshot Output Unguided 1_1](https://github.com/arkaan-r/109082500193_MohammadAkaanRiyanto/blob/main/LAPRAKSMT-2/output/soal2.png)
[ Program ini dibuat untuk menghitung fungsi.Didalamnya terdapat 4 func, pertama ada f(x)untuk menghitung f(x) = x^2, g(x)= x - 2, h(x) = x + 1, dan terakhir ada main disini tempat untuk membaca input a,b,c dan menghitung (FOGOh)(a), (GOHOF)(b), dan (HOFOG)(c)]
