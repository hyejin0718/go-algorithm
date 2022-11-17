package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26}
	isExist, index := binarySearch(array, -2)

	if isExist {
		fmt.Printf("찾았습니다. index: %d", index)
	} else {
		fmt.Printf("없습니다.")
	}
}

func binarySearch(sorted []int, want int) (bool, uint) {
	var start, end, mid, count uint
	start, count = 0, 0
	end = uint(len(sorted))
	mid = (start + end) / 2

	for {
		if start == end {
			break
		}

		midValue := sorted[mid]
		if want == midValue {
			return true, mid
		} else {
			if mid == 0 {
				break
			}
			if midValue > want {
				end = mid - 1
			} else if midValue < want {
				start = mid + 1
			}
			mid = (start + end) / 2
			fmt.Printf("#%d: index %d ~ %d 탐색, mid = %d\n", count, start, end, mid)
			count += 1
		}
	}
	return false, 0
}
