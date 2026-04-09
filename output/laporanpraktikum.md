# <h1 align="center">REVIEW ALGORITMA DAN PEMROGRAMAN 1 - ... </h1>
<p align="center">[Mohammad Arkaan Riyanto] - [109082500193]</p>

## Unguided 

### Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematikadiskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, isenguntuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalianmembantu Jonas? (tidak tentunya ya :p)Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi,dengan syarat a ≥ c dan b ≥ d.Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi aterhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.Catatan: permutasi (P) dan kombinasi (C) dari n terhadap r (n ≥ r) dapat dihitung denganmenggunakan persamaan berikut!
#### 

```go
package main

import "fmt"

fpackage main

import "fmt"

func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutasi(n, r int, hasil *int) {
	var nf, nrf int
	faktorial(n, &nf)
	faktorial(n-r, &nrf)
	*hasil = nf / nrf
}

func kombinasi(n, r int, hasil *int) {
	var nf, rf, nrf int
	faktorial(n, &nf)
	faktorial(r, &rf)
	faktorial(n-r, &nrf)
	*hasil = nf / (rf * nrf)
}

func main() {
	var a, b, c, d, p1, c1, p2, c2 int
	fmt.Scan(&a, &b, &c, &d)

	permutasi(a, c, &p1)
	kombinasi(a, c, &c1)
	permutasi(b, d, &p2)
	kombinasi(b, d, &c2)

	fmt.Println(p1, c1)
	fmt.Println(p2, c2)
}
```
### Output Unguided :

##### Output 

![Screenshot Output Unguided 1_1](https://github.com/arkaan-r/109082500193_MohammadAkaanRiyanto/blob/main/LAPRAKSMT-2/output/soal1.png)
[Program ini digunakan untuk menghitung permutasi dan kombinasi dari dua pasang bilangan yang diinputkan pengguna. Fungsi faktorial digunakan untuk menghitung n! dengan perulangan, dan menggunakan pointer agar hasil langsung disimpan ke variabel tanpa rturn(karna pointer berguna untuk menyimpan nilai sementara jadi tidak perlu mengembalikan nilai ke func lain). Fungsi permutasi dan kombinasi mengambil nilai dari func faktorial karna menggunakan pointer. Di fungsi main, program membaca empt bilangan, menghitung hasil untuk dua pasangan(permutasi dan kombinasi), lalu menampilkannya dalam dua baris.]
