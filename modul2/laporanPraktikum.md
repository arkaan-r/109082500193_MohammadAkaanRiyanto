# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>
<p align="center">[Mohammad Arkaan Riyanto] - [109082500193]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp string
	)
	fmt.Print("Masukkan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukkan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukkan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :	

##### Output 
![Screenshot Output Unguided 1_](https://github.com/MuhammadJovan/109082500024_Muhammad-Jovan/blob/main/modul2/output/output-soal1.png)
[Penjelasan] Program tersebut akan membuat algoritma yang membuat urutan angka tertukar, jadi disini terdapat 4 variable yaitu satu, dua, tiga dan temp dengan menggunakan tipe data yaitu string, lalu kita membuat sebuah inputan satu - tiga yang nanti akan dibuat untuk kita memasukkan inputan, setelah itu disini terdapat "temp" yang dimana fungsinya hanya untuk sebuah perpindahan inputan sementara jika tidak terdapat "temp" maka program akan error, nah hasil outputnya nanti jika kita memasukkan 1 2 3 maka outputnya akan menjadi 3 2 1.

//

### 2. [Soal]
#### soal2.go

```go
package main

import "fmt"

func main() {
	var warna1, warna2, warna3, warna4 string
	var BERHASIL = true

	for i:= 1; i <= 5; i++{

		fmt.Print("Percobaan", i, ": ")
		fmt.Scanln(&warna1, &warna2, &warna3, &warna4)

		if !(warna1 == "merah" && warna2 == "kuning" && warna3 == "hijau" && warna4 == "ungu") {
		BERHASIL = false
		}	
	}
	fmt.Println("BERHASIL:", BERHASIL)
}

```
### Output Unguided :	

##### Output 
![Screenshot Output Unguided 1_](https://github.com/MuhammadJovan/109082500024_Muhammad-Jovan/blob/main/modul2/output/output-soal2.png)
[Penjelasan] Program di atas adalah sebuah program looping atau perulangan, tugas program tersebut yaitu untuk mencocokkan inputan yang nanti output akan false atau true, di dalam program tersebut terdapat 4 variable warna type data string dan 1 variable yang berbentuk type data TRUE, lalu disini algoritma looping akan melakukan perulangan 1 sampai dengan 5, jika inputan sudah melampaui 5 maka algoritma akan berhenti dan jika if ! atau jika tidak sama dengan merah, kuning, hijau, ungu, maka program tersebut akan berubah menjadi false.

//

### 3. [Soal]
#### soal2.go

```go
package main

import	"fmt"

func main() {
	
	var gram int
	var kg, sisa int
	var biayaKg, biayaSisa, total int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&gram)

	kg = gram / 1000
	sisa = gram % 1000

	biayaKg = kg * 10000

	if kg < 10 {
		if sisa >= 500 {
			biayaSisa = sisa * 5
		} else {
			biayaSisa = sisa * 15
		}

	} else {
		biayaSisa = 0
	}
	total = biayaKg + biayaSisa
	fmt.Println("Detail berat:", kg, "kg +", sisa, "gr")
	fmt.Println("Detail biaya: Rp.", biayaKg, "+ Rp.", biayaSisa)
	fmt.Println("Total biaya: Rp.", total)
}

```
### Output Unguided :	

##### Output 
![Screenshot Output Unguided 1_](https://github.com/MuhammadJovan/109082500024_Muhammad-Jovan/blob/main/modul2/output/output-soal3.png)
[Penjelasan] Program di atas merupakan program if else, jadi dalam program tersebut terdapat 6 variable yaitu gram, kg, sisa, biayaKg, biayaSisa, total, dengan type data integer, program pertama akan di minta inputan "gram" lalu algoritma akan menghitung kg = gram / 1000 dan sisa = gram % 1000, lalu disini biayaKg adalah kg x 10000, masuk ke dalam if yaitu jika kg kurang dari 10 maka jika sisa lebih dari sama dengan 500 maka biayaSisa = sisa x 5 dan jika tidak maka biayaSisa = sisa x 15, lalu disini biayaSisa adalah 0 dan untuk perhitungannya yaitu total = biayaKg ditambahkan dengan biaya Sisa.