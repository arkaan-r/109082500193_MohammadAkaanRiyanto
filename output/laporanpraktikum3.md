# <h1 align="center">REVIEW ALGORITMA DAN PEMROGRAMAN 1 - ... </h1>
<p align="center">[Mohammad Arkaan Riyanto] - [109082500193]</p>

## Unguided 

### PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka,buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuansebagai berikut!Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalamgram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yangkurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawahadalah input/read):
#### 

```go
package main

import "fmt"

func main() {
	var parsel, biaya, detail, total int

	fmt.Print("Berat parsel dalam gram: ")
	fmt.Scan(&parsel)

	kg := parsel / 1000
	gr := parsel % (kg * 1000)
	biaya = kg * 10000

	fmt.Println("Detail berat:", kg, "kg", "+", gr, "gr")

	if kg > 10 {
		detail = gr * 5
		total = biaya
	} else if gr >= 500 {
		detail = gr * 5
		total = biaya + detail
	} else {
		detail = gr * 15
		total = biaya + detail
	}

	fmt.Println("detail biaya: ", "Rp.", biaya, " + Rp.", detail)
	fmt.Print("Total biaya: Rp.", total)
}


```
### Output Unguided :

##### Output 

![Screenshot Output Unguided 1_1](https://github.com/arkaan-r/109082500193_MohammadAkaanRiyanto/blob/main/LAPRAKSMT-2/output/soal3.png)
[ Program ini dibuat untuk mengecek kombinasi 4 warna dalam percobaan. Jika warnanya merah, kuning, hijau, ungu maka true. Sebaliknya jika tidak sesuai maka false.]
