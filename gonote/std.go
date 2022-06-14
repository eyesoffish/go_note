package note

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// 随机数
func RandNum() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(10) + 1)
}


// 字符串类型转换
func StrConv() {
	i1 := 124
	s1 := "dasdfjkl"
	s2 := fmt.Sprintf("%d@%s", i1, s1)
	fmt.Println(s2)
	var (
		i2 int
		s3 string
	)
	n,err := fmt.Sscanf(s2, "%d@%s", &i2, &s3)
	if err != nil {
		panic(err)
	}
	fmt.Printf("success = i2 = %v, s3 = %v", n, s3)

	s4 := strconv.FormatInt(123, 4)
	u1, err := strconv.ParseUint(s4, 4, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("u1 = %T", u1)
}