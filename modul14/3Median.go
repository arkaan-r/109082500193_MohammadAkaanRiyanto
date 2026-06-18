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