# <h1 align="center">REVIEW ALGORITMA DAN PEMROGRAMAN 1 - ... </h1>
<p align="center">[Mohammad Arkaan Riyanto] - [109082500193]</p>

## Unguided 

### Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya. Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur. prosedure hitungSkor(in/out soal, skor : integer) Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah 8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal. Jika tidak berhasil atau tidak mengirimkan jawaban maka otomatis dianggap menyelesaikan dalam waktu 5 jam 1 menit (301 menit). Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil diselesaikan.
#### 

```go
package main

import "fmt"

func hitungSkor(soal *int, skor *int) {
	var time int
	*soal = 0
	*skor = 0
	for i := 0; i < 8; i++ {
		fmt.Scan(&time)
		if time <= 300 {
			*soal++
			*skor += time
		}
	}
}

func main() {
	var name, winner string
	var soal, skor, jmlsoal, jmlskor int
	jmlskor = 999999

	fmt.Scan(&name)

	for name != "selesai" {
		hitungSkor(&soal, &skor)

		if soal > jmlsoal || (soal == jmlsoal && skor < jmlskor) {
			jmlsoal = soal
			jmlskor = skor
			winner = name
		}
		fmt.Scan(&name)
	}

	if winner != "" {
		fmt.Printf("%s %d %d\n", winner, jmlsoal, jmlskor)
	}
}

```
### Output Unguided :

##### Output 

![Screenshot Output Unguided 1_1](https://github.com/arkaan-r/109082500193_MohammadAkaanRiyanto/blob/main/LAPRAKSMT-2/output/soal2.png)
[ program membaca nama peserta satu per satu sampai memasukkan "selesai" agar program berhenti. Untuk setiap peserta, program memanggil fungsi hitungSkor dengan i <= 8 karna ada 8 soal yang diberikan dan waktu <= 300 ini juga karna diberikan waktu 5 jam. Di fungsi hitungSkor, program membaca waktu dari 8 soal. Jika waktu ≤ 300, soal dihitung benar dan akan ditambahkan dengan waktu. Program ini memakai pointer supaya nilai soal dan skor bisa langsung berubah tanpa return. Setelah itu, program membandingkan hasil setiap peserta. Pemenang adalah yang soalnya paling banyak, dan pengerjaan yang cepat.]
