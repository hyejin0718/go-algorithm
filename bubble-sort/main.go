package main

import "fmt"

func main() {
	array := []int{16, 14, 4, 21, 11, 12, 15, 22, 25, 26, 1, 5, 6, 17, 2, 3, 13}
	sort := bubbleSort(array)
	fmt.Println(sort)
}

func bubbleSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := 0; j+1 < len(a)-i; j++ {
			if a[j] < a[j+1] { // 내림차순 정렬
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}
