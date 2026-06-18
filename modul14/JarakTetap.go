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