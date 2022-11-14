package main

import "fmt"

type Arr struct {
	data  []int
	left  int
	right int
	pivot int
}

func NewArr(array []int) Arr {
	return Arr{
		data:  array,
		left:  0,
		right: len(array) - 2,
		pivot: len(array) - 1,
	}
}

func main() {
	arr := NewArr([]int{15, 12, 13, 17, 1, 14, 3, 6, 19, 16, 2, 20, 4, 11, 5, 7, 8, 18, 9, 10})
	arr.Split(0, len(arr.data))
	fmt.Println("result:", arr.data)

	arr2 := NewArr([]int{1, 2, 3, 4, 2, 2, 1, 1, 4, 4, 3, 3})
	arr2.Split(0, len(arr2.data))
	fmt.Println("result:", arr2.data)
}

func (a *Arr) Split(start, end int) {
	if end <= start {
		return
	}
	a.left = start
	a.right = end - 2
	a.pivot = end - 1
	sorted := a.pivot
	fmt.Println(a.data[start:end])

	for {
		a.leftMove(end)
		a.rightMove(start)
		if a.left >= a.right {
			break
		}
		a.swap(a.left, a.right)
	}
	if a.data[a.left] > a.data[a.pivot] {
		a.swap(a.left, a.pivot)
		sorted = a.left
	}
	fmt.Println(a.data[start:end], "sorted: ", a.data[sorted])

	a.Split(start, sorted)
	a.Split(sorted+1, end)
}

func (a *Arr) leftMove(end int) {
	//왼쪽값 > 피벗값이 될때까지 하나씩 오른쪽으로 옮긴다.
	for {
		if a.left < end-1 && a.data[a.left] <= a.data[a.pivot] {
			a.left++
		} else {
			break
		}
	}
}

func (a *Arr) rightMove(start int) {
	//오른쪽값 < 피벗값이 될때까지 하나씩 왼쪽으로 옮긴다.
	for {
		if a.right > start && a.data[a.right] >= a.data[a.pivot] {
			a.right--
		} else {
			break
		}
	}
}

func (a *Arr) swap(left, right int) {
	a.data[left], a.data[right] = a.data[right], a.data[left]
}
