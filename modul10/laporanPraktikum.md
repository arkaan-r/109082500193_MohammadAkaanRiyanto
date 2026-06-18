# <h1 align="center">Laporan Praktikum Modul 10 - ... </h1>
<p align="center">[Mohammad Arkaan Riyanto] - [109082500193]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func hitung(kelinci[1000]float64, A int) {
	for i := 0; i < A; i++{
		fmt.Printf("Anak kelinci ke %d: ",i+1)
		fmt.Scan(&kelinci[i])
	}
	max := kelinci[0]
	min := kelinci[0]
	for i := 1; i < A; i++{
	if kelinci[i] > max {
		max = kelinci[i]
	}else if kelinci[i] < min {
		min = kelinci[i]
	}
}
	fmt.Println("Berat kelinci terbesar: ",max)
	fmt.Print("Berat kelinci terkecil: ",min)

}

func main() {
	var kelinci[1000] float64
	var N int
	fmt.Print("Jumlah anak kelinci: ")
	fmt.Scan(&N)

	hitung(kelinci, N)

}
```
### Output Unguided :	

##### Output 
![Screenshot Output Unguided 1_](https://github.com/MuhammadJovan/109082500024_Muhammad-Jovan/blob/main/modul10/output/output-soal1.png)
[Penjelasan]Program ini berfungsi untuk mengumpulkan input berat badan sejumlah kelinci lalu menyimpannya ke dalam sebuah array. Setelah datanya terkumpul maka program akan membandingkan setiap nilai secara otomatis untuk menentukan dan menampilkan mana berat yang paling besar dan mana yang paling kecil.

//

### 2. [Soal]
#### soal2.go

```go
package main

import "fmt"

func beratIkan(ikan [1000]float64) {
	var x, y int
	fmt.Print("Ikan yang akan dijual (x): ")
	fmt.Scan(&x)
	fmt.Print("Wadah (y): ")
	fmt.Scan(&y)

		fmt.Println()

	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&ikan[i])
	}
	fmt.Println("Pembagian wadah: ")
	for i := 0; i < x; i += y {
		var total float64 = 0
		var jumlah int = 0

		for j := i; j < i+y && j < x; j++ {
			total += ikan[j]
			jumlah++
		}
		fmt.Printf("Wadah %d: %.2f kg\n", (i/y)+1, total/float64(jumlah))
	}
}
func main() {
	var ikan [1000]float64
	beratIkan(ikan)
}

```
### Output Unguided :	

##### Output 
![Screenshot Output Unguided 1_](https://github.com/MuhammadJovan/109082500024_Muhammad-Jovan/blob/main/modul10/output/output-soal2.png)
[Penjelasan] Program ini berfungsi untuk membagi data berat ikan ke dalam beberapa wadah berdasarkan jumlah yang kamu tentukan. Setelah menginput berat tiap ikan, kode tersebut akan menghitung rata-rata berat ikan di setiap wadah secara otomatis.
//

### 3. [Soal]
#### soal2.go

```go
package main

import "fmt"

type arrBalita [100]float64
func hitungMinMax(arrBerat *arrBalita, bMin, bMax *float64, balita *int) {
	
	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(balita)
	for i := 0; i < *balita; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ",i+1)
		fmt.Scan(&arrBerat[i])
	}
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for i := 1; i < *balita; i++ {
		if arrBerat[i] < *bMin {	
			*bMin = arrBerat[i]
		}else if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
		
	}
		fmt.Printf("Berat balita minimum: %.2f", *bMin)
		fmt.Println()
		fmt.Printf("Berat balita maksimum: %.2f", *bMax)
		fmt.Println()
}
	
func rerata(arrBerat *arrBalita, balita int) float64 {
	var total float64

	if balita == 0 {
		return 0
	}
	for i := 0; i < balita; i++ {
		total = total + arrBerat[i]
	}
	return total / float64(balita)
	
}
func main() {
	var arrBerat arrBalita
	var bMin, bMax float64
	var balita int
	hitungMinMax(&arrBerat, &bMin, &bMax, &balita)
	fmt.Printf("Rata-rata berat balita: %.2f", rerata(&arrBerat, balita))
}


```
### Output Unguided :	

##### Output 
![Screenshot Output Unguided 1_](https://github.com/MuhammadJovan/109082500024_Muhammad-Jovan/blob/main/modul10/output/output-soal3.png)
[Penjelasan]Program tersebut digunakan untuk ngolah data berat badan balita dengan mencari nilai ekstrem sekaligus menghitung rata-ratanya. Program akan meminta input berat badan satu per satu, lalu secara otomatis nentuin berat yang paling ringan dan yang paling berat lalu menampilkan hasil perhitungan rata" seluruh data itu.
//

