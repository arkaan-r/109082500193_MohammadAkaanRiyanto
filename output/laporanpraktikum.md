# <h1 align="center">REVIEW ALGORITMA DAN PEMROGRAMAN 1 - ... </h1>
<p align="center">[Mohammad Arkaan Riyanto] - [109082500193]</p>

## Unguided 

### 1.Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakanmasukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yangdiperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut?
#### soal1.go

```go
package main
import "fmt"

func main() {
    var (
        satu, dua, tiga string
        temp string
    )

    fmt.Print("Masukan input string: ")
    fmt.Scanln(&satu)

    fmt.Print("Masukan input string: ")
    fmt.Scanln(&dua)

    fmt.Print("Masukan input string: ")
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

![Screenshot Output Unguided 1_1](https://github.com/arkaan-r/109082500193_MohammadAkaanRiyanto/blob/main/LAPRAKSMT-2/output/soal1.png)
[Saat program dijalankan, kita akan diminta untuk memasukkan tiga data bertipe string secara berurutan. Setiap input yang kita masukkan akan disimpan ke dalam variabel satu, dua, dan tiga sesuai dengan urutan inputnya. Setelah itu dilakukan proses pertukaran nilai, yaitu nilai pada variabel satu dipindahkan ke variabel temp. Kemudian nilai pada variabel dua menjadi nilai baru untuk variabel satu, nilai pada variabel tiga menjadi nilai baru untuk variabel dua, dan terakhir nilai yang ada di variabel temp dipindahkan ke variabel tiga. Setelah proses tersebut selesai, program akan menampilkan isi akhir dari variabel satu, dua, dan tiga.]
