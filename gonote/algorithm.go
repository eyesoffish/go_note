package note

import "fmt"

// 递归
var fibonacciRes []int
func fibonacci(n int) int {
	if n < 3 {
		return 1
	}
	if fibonacciRes[n] == 0 {
		fibonacciRes[n] = fibonacci(n-2) + fibonacci(n-1)
	}
	return fibonacciRes[n]
}

func Recursion() {
	n := 5;
	fibonacciRes = make([]int, n+1)
	fmt.Printf("第%v位斐波那契额数是%v\n", n, fibonacci(n))
}