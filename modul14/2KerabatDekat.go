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