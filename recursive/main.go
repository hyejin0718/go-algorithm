package main

import "fmt"

func main() {
	for i := 1; i <= 15; i++ {
		fmt.Println("fibonacci ", i, ": ", fibonacci(uint(i)))
	}
	fmt.Println("---")

	for i := 1; i <= 15; i++ {
		fmt.Println("factorial ", i, ": ", factorial(uint(i)))
	}
	fmt.Println("---")

	for i := 1; i <= 15; i++ {
		fmt.Println("hanoi ", i, ": ", hanoi(uint(i)))
	}
}

func fibonacci(n uint) uint {
	if n <= 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func factorial(n uint) uint {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func hanoi(n uint) uint {
	if n <= 1 {
		return 1
	}
	return hanoi(n-1)*2 + 1
}

// 마지막 막대가 아닌 곳에 n-1 층의 하노이 탑을 옮긴다. -> hanoi(n-1) 회
// 가장 아래층을 마지막 막대에 옮긴다. -> 1 회
// 마지막 막대에 n-1 층의 하노이 탑을 옮긴다. -> hanoi(n-1) 회
