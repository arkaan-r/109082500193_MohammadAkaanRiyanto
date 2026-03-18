# <h1 align="center">REVIEW ALGORITMA DAN PEMROGRAMAN 1 - ... </h1>
<p align="center">[Mohammad Arkaan Riyanto] - [109082500193]</p>

## Unguided 

### Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematikadiskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, isenguntuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalianmembantu Jonas? (tidak tentunya ya :p)Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi,dengan syarat a ≥ c dan b ≥ d.Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi aterhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.Catatan: permutasi (P) dan kombinasi (C) dari n terhadap r (n ≥ r) dapat dihitung dengan menggunakan persamaan berikut!

```go
package main

import "fmt"

func faktorial(n int) int {
	hasil := 1
	for i := 2; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}
func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}
func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}
func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(permutasi(a, c), kombinasi(a, c))
	fmt.Println(permutasi(b, d), kombinasi(b, d))
}
```
### Output Unguided :

##### Output 

![Screenshot Output Unguided 1_1](https://github.com/arkaan-r/109082500193_MohammadAkaanRiyanto/blob/main/LAPRAKSMT-2/output/soal1.png)
[ Program ini digunakan untuk menghitung prmutasi dan kombinasi dari 4 inputan user. disini ada 4 func yang pertama fktorial untuk menghitung faktorial menggunakan loop dimana loop akan berjalan dari 2 ke n dan mengalikan tiap i nya. Lalu func kedua untuk permutasi yang menghitung hasil bagi dari faktorial(n)/faktorial(n-r). Kemudian func ketiga untuk kombinasi bedanya dengan permutasi hanya pembaginya yang diubah menjadi faktorial(r)*faktorial(n-r). Yang terakhir ada func main disini hanya untuk membaca inputan user dan memunculkan output program.]
