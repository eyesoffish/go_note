package note

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
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

func PackageString() {
	str := "abcd"
	fmt.Println("is Contains = ", strings.Contains(str, "dd"))
	fmt.Println("str index = ", strings.Index(str, "dd"))
	fmt.Println("str index = ", strings.Replace(str, "d", " hello", 10))
	fmt.Println("str = ", strings.Repeat(str, 5))	
	fmt.Println("str = ", strings.Split(str, ""))
	// 字符串修剪
	fmt.Println("str = ", strings.Trim("#*\nwww.www.www&%$", "!@#$%^&*\n"))
}


// 中文字符操作
func PackageUFT8() {
	str := "hello 世界, 🦴"
	fmt.Println(utf8.ValidString(str))
}

// 时间操作

func Time() {
	for i := 0; i < 5; i++ {
		fmt.Print(".")
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Println()
	d1, err := time.ParseDuration("1000s")
	if err != nil {
		panic(err)
	}
	fmt.Printf("time = %v, type = %T\n", d1, d1)

	// parse
	t1, err := time.Parse("2006年2月2日, 15点4分", "2022年6月5日, 18点8分")
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("time = %v, type = %T", t1, t1)

	// 超时控制
	var intChan chan int = make(chan int)
	// 理解: chan 等待数据流入1s之后结束
	select {
	case  <- intChan:
		fmt.Println("收到了用户发送的验证码")
	case <- time.After(time.Second):
		fmt.Println("验证码1s超时")
	}

	// 时区
	l1, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	fmt.Println("时区---", l1)
	fmt.Println(time.Now().Format("2006/01/02 15:04:00"))

	// 返回时区缩写 CST
	z,_ := time.Now().Zone()
	fmt.Println("时区---", z)

	go func() {
		time.Sleep(time.Second)
		intChan <- 1
	}()
	Ticker:
	for {
		select {
		case <- intChan:
			fmt.Println()
			break Ticker
		case <- time.NewTicker(100 * time.Millisecond).C:
			fmt.Print(".")
		}
	}
	// go func() {
	// 	time.Sleep(time.Millisecond * 900)
	// 	intChan <- 1
	// }()
	// 单词计数器
	select {
	case <- intChan:
		fmt.Println("用户收到了验证啊")
	case <- time.NewTicker(time.Second).C:
		fmt.Println("验证码已过期")
	}
}
