package note

import (
	"fmt"
	"math/rand"
	"time"
)

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

// 闭包
func Closure() {
	f := closureFunc()
	f(10)
	f(11)
	f = closureFunc()
	f(12)
}

func closureFunc() func(int) int {
	i := 0
	return func(n int) int {
		fmt.Printf("本次调用接收到n=%v\n", n)
		i++
		fmt.Printf("匿名函数被%v次调用\n", i)
		return i
	}
}

func bobbleSort(s []int) {
	lastIndex:=len(s) - 1;
	for i := 0; i < lastIndex; i++ {
		for j := 0; j < lastIndex - i; j++ {
			if s[j] > s[j+1] {
				t:=s[j]
				s[j] = s[j+1]
				s[j+1] = t
			}
		}
	}
}

func Sort() {
	n:=100
	s:=make([]int,n)
	seedNum := time.Now().UnixNano()
	for i := 0; i < n; i++ {
		rand.Seed(seedNum)
		s[i] = rand.Intn(10001)
		seedNum ++
	}
	fmt.Println(s)
	bobbleSort(s)
	fmt.Println("排序后",s)
}