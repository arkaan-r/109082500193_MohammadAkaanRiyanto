# <h1 align="center">Laporan Praktikum Modul 14 - ... </h1>
<p align="center">[Mohammad Arkaan Riyanto] - [109082500193]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		min := i

		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		arr[i], arr[min] = arr[min], arr[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		arr := make([]int, m)

		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j])
		}

		selectionSort(arr)

		for j := 0; j < m; j++ {
			fmt.Print(arr[j])
			if j != m-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
```
### Output Unguided :	

##### Output 
![Screenshot Output Unguided 1_](https://github.com/MuhammadJovan/109082500024_Muhammad-Jovan/blob/main/modul14/output/output-soal1.png)
[Penjelasan]Program ini digunakan untuk mengurutkan nomor rumah kerabat Hercules pada setiap daerah menggunakan algoritma Selection Sort. Program diawali dengan membaca jumlah daerah yang akan diproses, kemudian untuk setiap daerah dibaca banyaknya rumah beserta nomor rumah para kerabat. Seluruh nomor rumah disimpan ke dalam array dan diurutkan secara menaik. Algoritma Selection Sort bekerja dengan mencari elemen terkecil dari bagian array yang belum terurut, kemudian menukarnya dengan elemen pada posisi saat ini. Proses tersebut dilakukan berulang hingga seluruh data tersusun dari nilai terkecil ke terbesar. Setelah proses pengurutan selesai, program menampilkan nomor rumah yang telah terurut untuk setiap daerah.
//

### 2. [Soal]
#### soal2.go

```go
package main

import "fmt"

func insertionAsc(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}
}

func insertionDesc(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] < key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		var ganjil []int
		var genap []int

		for j := 0; j < m; j++ {
			var x int
			fmt.Scan(&x)

			if x%2 == 0 {
				genap = append(genap, x)
			} else {
				ganjil = append(ganjil, x)
			}
		}

		insertionAsc(ganjil)
		insertionDesc(genap)

		for _, v := range ganjil {
			fmt.Print(v, " ")
		}

		for _, v := range genap {
			fmt.Print(v, " ")
		}

		fmt.Println()
	}
}

```
### Output Unguided :	

##### Output 
![Screenshot Output Unguided 1_](https://github.com/MuhammadJovan/109082500024_Muhammad-Jovan/blob/main/modul14/output/output-soal2.png)
[Penjelasan] Program ini dibuat untuk membantu Hercules mengunjungi kerabatnya dengan jumlah penyeberangan jalan yang seminimal mungkin. Program menerima sejumlah nomor rumah pada setiap daerah, kemudian memisahkan nomor rumah ganjil dan genap ke dalam kelompok yang berbeda. Nomor rumah ganjil diurutkan secara menaik menggunakan algoritma Insertion Sort, sedangkan nomor rumah genap diurutkan secara menurun. Setelah proses pengurutan selesai, program mencetak seluruh nomor rumah ganjil terlebih dahulu, kemudian diikuti nomor rumah genap. Dengan susunan tersebut, Hercules dapat mengunjungi seluruh rumah pada satu sisi jalan terlebih dahulu sebelum berpindah ke sisi jalan yang lain.
//

### 3. [Soal]
#### soal2.go

```go
package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}
}

func median(arr []int) int {
	temp := make([]int, len(arr))
	copy(temp, arr)

	insertionSort(temp)

	n := len(temp)

	if n%2 == 1 {
		return temp[n/2]
	}

	return (temp[n/2-1] + temp[n/2]) / 2
}

func main() {
	var data []int
	var x int

	for {
		fmt.Scan(&x)

		if x == -5313 {
			break
		}

		if x == 0 {
			fmt.Println(median(data))
		} else {
			data = append(data, x)
		}
	}
}


```
### Output Unguided :	

##### Output 
![Screenshot Output Unguided 1_](https://github.com/MuhammadJovan/109082500024_Muhammad-Jovan/blob/main/modul14/output/output-soal3.png)
[Penjelasan]Program ini digunakan untuk menghitung nilai median dari sekumpulan data bilangan bulat. Data dibaca satu per satu dan disimpan ke dalam array. Setiap kali program menemukan angka 0, seluruh data yang telah tersimpan akan diurutkan menggunakan algoritma Insertion Sort, kemudian median dari data tersebut dihitung dan ditampilkan. Jika jumlah data ganjil, median adalah elemen yang berada tepat di tengah setelah data diurutkan. Jika jumlah data genap, median diperoleh dari rata-rata dua nilai tengah dan hasilnya dibulatkan ke bawah sesuai ketentuan soal. Proses ini dapat terjadi beberapa kali selama pembacaan data. Program akan berhenti ketika menemukan bilangan penanda akhir yaitu -5313. Dengan demikian, setiap kemunculan angka 0 berfungsi sebagai perintah untuk menampilkan median dari seluruh data yang telah dibaca hingga saat itu.
//

### 1. [Soal]
#### soal2.go

```go
package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}
}

func main() {
	var x int
	var arr []int

	for {
		fmt.Scan(&x)

		if x < 0 {
			break
		}

		arr = append(arr, x)
	}

	insertionSort(arr)

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])

		if i != len(arr)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	if len(arr) <= 2 {
		fmt.Println("Data berjarak tetap")
		return
	}

	selisih := arr[1] - arr[0]
	tetap := true

	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != selisih {
			tetap = false
			break
		}
	}

	if tetap {
		fmt.Printf("Data berjarak %d\n", selisih)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}


```
### Output Unguided :	

##### Output 
![Screenshot Output Unguided 1_](https://github.com/MuhammadJovan/109082500024_Muhammad-Jovan/blob/main/modul14/output/output-soal4.png)
[Penjelasan]Program ini digunakan untuk membaca sekumpulan bilangan bulat non-negatif yang diakhiri oleh bilangan negatif. Seluruh bilangan non-negatif yang dimasukkan pengguna disimpan ke dalam array, kemudian diurutkan menggunakan algoritma Insertion Sort. Algoritma ini bekerja dengan mengambil satu elemen dan menyisipkannya ke posisi yang sesuai pada bagian array yang telah terurut sebelumnya. Setelah seluruh data terurut, program memeriksa apakah selisih antara setiap elemen dengan elemen sebelumnya selalu sama. Jika seluruh selisih bernilai sama, maka program menampilkan status bahwa data memiliki jarak tetap beserta nilai jaraknya. Sebaliknya, jika terdapat selisih yang berbeda, maka program menampilkan informasi bahwa data memiliki jarak yang tidak tetap.
//

### 2. [Soal]
#### soal2.go

```go
package main

import "fmt"

type Buku struct {
	id       string
	judul    string
	penulis  string
	penerbit string
	eksemplar int
	tahun    int
	rating   int
}

func insertionSort(buku []Buku) {
	n := len(buku)

	for i := 1; i < n; i++ {
		temp := buku[i]
		j := i - 1

		for j >= 0 && buku[j].rating < temp.rating {
			buku[j+1] = buku[j]
			j--
		}

		buku[j+1] = temp
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	data := make([]Buku, n)

	for i := 0; i < n; i++ {
		fmt.Scan(
			&data[i].id,
			&data[i].judul,
			&data[i].penulis,
			&data[i].penerbit,
			&data[i].eksemplar,
			&data[i].tahun,
			&data[i].rating,
		)
	}

	insertionSort(data)

	fmt.Println("Buku Terfavorit:")
	fmt.Println(data[0].judul, data[0].penulis, data[0].penerbit, data[0].tahun)

	fmt.Println("5 Buku Rating Tertinggi:")

	limit := 5
	if n < 5 {
		limit = n
	}

	for i := 0; i < limit; i++ {
		fmt.Println(data[i].judul)
	}

	var cari int
	fmt.Scan(&cari)

	found := false

	for i := 0; i < n; i++ {
		if data[i].rating == cari {
			fmt.Println(data[i])
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Tidak ada buku dengan rating tersebut")
	}
}


```
### Output Unguided :	

##### Output 
![Screenshot Output Unguided 1_](https://github.com/MuhammadJovan/109082500024_Muhammad-Jovan/blob/main/modul14/output/output-soal5.png)
[Penjelasan]Program ini digunakan untuk mengelola data buku pada sebuah perpustakaan. Setiap buku memiliki beberapa atribut seperti id, judul, penulis, penerbit, jumlah eksemplar, tahun terbit, dan rating. Data buku yang dimasukkan pengguna disimpan ke dalam array bertipe struct. Selanjutnya data buku diurutkan berdasarkan nilai rating secara menurun menggunakan algoritma Insertion Sort sehingga buku dengan rating tertinggi berada pada posisi pertama. Setelah proses pengurutan selesai, program menampilkan buku terfavorit, yaitu buku yang memiliki rating tertinggi, kemudian menampilkan lima buku dengan rating tertinggi. Selain itu, program juga dapat melakukan pencarian buku berdasarkan rating yang dimasukkan pengguna. Jika ditemukan buku dengan rating yang dicari, seluruh informasi buku akan ditampilkan, sedangkan jika tidak ditemukan maka program akan menampilkan pesan bahwa buku dengan rating tersebut tidak tersedia.
//
