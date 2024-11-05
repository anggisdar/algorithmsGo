package main

import "fmt"

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		// inner lop untuk pembanding yang berdekatan
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// tukar elemen yang sebelumnya jika lebih besar yang sesudah
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}

		}
	}
	return arr
}

func main() {
	// daftar angkanya
	data := []int{20, 24, 01, 04, 30, 9, 24}

	fmt.Println("data menjadi terurut:", bubbleSort(data))
}
