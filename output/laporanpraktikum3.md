# <h1 align="center">REVIEW ALGORITMA DAN PEMROGRAMAN 1 - ... </h1>
<p align="center">[Mohammad Arkaan Riyanto] - [109082500193]</p>

## Unguided 

### [Lingkaran] Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y)berdasarkan dua lingkaran tersebut.Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusatdan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titiksembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan denganbilangan bulat.Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titikdi dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".
#### 

```go
package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}
func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}
func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y float64

	fmt.Scanln(&cx1, &cy1, &r1)
	fmt.Scanln(&cx2, &cy2, &r2)
	fmt.Scanln(&x, &y)

	in1 := didalam(cx1, cy1, r1, x, y)
	in2 := didalam(cx2, cy2, r2, x, y)

	if in1 && in2 {
		fmt.Println("Titik didalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik didalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik didalam lingkaran 2")
	} else {
		fmt.Println("Titik diluar lingkaran 1 dan 2")
	}
}



```
### Output Unguided :

##### Output 

![Screenshot Output Unguided 1_1](https://github.com/arkaan-r/109082500193_MohammadAkaanRiyanto/blob/main/LAPRAKSMT-2/output/soal3.png)
[Program ini dibuat untuk menentukan lokasi titik sembarang.Di program ini terdapat 3 func, Pertama untuk menghitung jarak titik sembarang dari titik pusat,kedua untuk membandingkan jarak titik sembarang terhadap jari" lingkaran, dan yang terakhir main untuk membaca inputan user dan menentukan output yang akan dikeluarkan program.]
