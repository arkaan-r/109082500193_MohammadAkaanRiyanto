# <h1 align="center">REVIEW ALGORITMA DAN PEMROGRAMAN 1 - ... </h1>
<p align="center">[Mohammad Arkaan Riyanto] - [109082500193]</p>

## Unguided 

### Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakanpraktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang manasusunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa dimintauntuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunanwarna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah ‘merah’, ‘kuning’,‘hijau’, dan ‘ungu’ selama 5 kali percobaan berulang.Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksisebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutanwarna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan falseuntuk urutan warna lainnya.Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawahadalah input/read):
#### soal1.go

```go
package main

import "fmt"

func main() {
	var a, b, c, d string
	berhasil := true
	for i := 1; i <= 5; i++ {
		fmt.Print("percobaan ", i, ": ")
		fmt.Scan(&a, &b, &c, &d)
		if !(a == "merah" && b == "kuning" && c == "hijau" && d == "ungu") {
			berhasil = false
		}
	}
	if berhasil {
		fmt.Print("BERHASIL : true")
	} else {
		fmt.Print("BERHASIL : false")
	}
}

```
### Output Unguided :

##### Output 

![Screenshot Output Unguided 1_1](https://github.com/arkaan-r/109082500193_MohammadAkaanRiyanto/blob/main/LAPRAKSMT-2/output/soal2.png)
[ Program ini dibuat untuk mengecek kombinasi 4 warna dalam percobaan. Jika warnanya merah, kuning, hijau, ungu maka true. Sebaliknya jika tidak sesuai maka false.]
