package note

import (
	"fmt"
	"goproject/gonote/util"
	"math/rand"
	"sort"
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

// 冒泡排序
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

// 选择排序
func selectionSort(s [] int) {
	lastIndex:=len(s) - 1;
	for i := 0; i < lastIndex; i++ {
		max := lastIndex - i
		for j := 0; j < lastIndex - i; j++ {
			if s[j] < s[max] {
				max = j
			}
		}
		if max != lastIndex - i {
			t := s[lastIndex - i]
			s[lastIndex - i] = s[max]
			s[max] = t
		}
	}
}

// 插入排序
func insertSort(s []int) {
	for i := 1; i < len(s); i++ {
		t:= s[i]
		j:=i-1
		for ; j >= 0 && s[j] > t; j-- {
			s[j+1] = s[j]
		}
		if j != i-1 {
			s[j+1] = t
			fmt.Println("s=",s)
		}
	}
}

// 快速排序
func fastSort(s []int, leftIndex, rightIndex int) {
	if leftIndex < rightIndex {
		pivot := s[rightIndex]
		var rs []int
		l := leftIndex
		for i := leftIndex; i < rightIndex; i++ {
			if s[i] > pivot {
				rs = append(rs, s[i])
			} else {
				s[l] = s[i]
				l++
			}
		}
		s[l] = pivot
		copy(s[l+1:], rs)
		if leftIndex < l-1 {
			fastSort(s, leftIndex, l - 1)
		}
		if l+1 < rightIndex {
			fastSort(s, l+1, rightIndex)
		}
	}
}

func Sort() {
	n:=10
	s:=make([]int,n)
	seedNum := time.Now().UnixNano()
	for i := 0; i < n; i++ {
		rand.Seed(seedNum)
		s[i] = rand.Intn(10001)
		seedNum ++
	}
	fmt.Println(s)
	// bobbleSort(s)
	// selectionSort(s)
	// insertSort(s)
	fastSort(s,0, len(s)- 1)
	fmt.Println("排序后",s)
}

// 二分查找
func BinarySearch(s[]int, key int) int {
	startIndex := 0;
	endIndex := len(s) - 1 
	midIndex := 0
	for startIndex <= endIndex {
		midIndex = startIndex + (endIndex - startIndex) / 2
		if s[midIndex] < key {
			startIndex = midIndex + 1
		} else if s[midIndex] > key {
			endIndex = midIndex - 1
		} else {
			return midIndex
		}
	}
	return -1
}

func BinarySearchText() {
	s:= make([]int, util.RandInt(1000) + 1)
	for i := 0; i < len(s); i++ {
		s[i] = util.RandInt(1000)
	}
	sort.Ints(s)
	i:=BinarySearch(s, 555)
	fmt.Println("找到下标=",i)
}