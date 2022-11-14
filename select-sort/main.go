package main

import "fmt"

func main() {
	array := []int{16, 21, 14, 4, 11, 12, 15, 22, 25, 26, 1, 5, 6, 17, 2, 4, 4, 3, 13}
	//sort := selectSort(array)
	//fmt.Println(sort)
	sort2 := selectSort2(array)
	fmt.Println(sort2)
}

func selectSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		minIndex := i
		for j := i; j < len(a); j++ {
			if a[j] <= a[minIndex] {
				minIndex = j
			}
		}
		fmt.Println(i, minIndex, a[i], a[minIndex])
		a[i], a[minIndex] = a[minIndex], a[i]
	}
	return a
}

// selectSort 최적화
func selectSort2(a []int) []int {
	// 같은 index 의 요소를 비교하지 않도록 반복 범위 변경
	for i := 0; i < len(a)-1; i++ { // 불필요한 반복 제거1
		minIndex := i
		for j := i + 1; j < len(a); j++ { // 불필요한 반복 제거2
			if a[j] < a[minIndex] { // 값이 같을 때 minIndex 갱신하지 않음
				minIndex = j
			}
		}
		fmt.Println(i, minIndex, a[i], a[minIndex])
		a[i], a[minIndex] = a[minIndex], a[i]
	}
	return a
}
