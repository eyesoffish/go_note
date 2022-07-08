package note

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"goproject/gonote/util"
	"math/rand"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
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


// 文件读写
func FileReadAndWrite() {
	file, err := os.OpenFile("gonote/f1.txt", os.O_WRONLY | os.O_CREATE, 0666);
	if err != nil {
		panic(err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	for i := 1; i <= 4; i++ {
		fileName := fmt.Sprintf("gonote/f%d.txt",i)
		data, err := os.ReadFile(fileName)
		if err != nil {
			panic(err)
		}
		data = append(data, '\n')
		write.Write(data) // 写入缓冲区
	}
	write.Flush() // 写入硬盘
}

// 错误 err 默认值nil
func Errors() {
	// 
	defer func() {
		err := recover()
		if err != nil {fmt.Println("捕捉到了错误", err)}
	}()
	err1 := errors.New("可爱的错误")
	fmt.Println(err1)
	err2 := fmt.Errorf("温柔的错误 , %v", 2)
	fmt.Println(err2)
}

// 日志
func Log() {
	defer func ()  {
		err := recover()
		if err != nil {
			fmt.Println("捕捉到了错误->Log")
		}
	}()
	err1 := errors.New("可爱的错误1")
	// err2 := errors.New("可爱的错误2")
	err3 := errors.New("可爱的错误3")
	util.INFO.Println(err1)
	// panic 和fatal 都会导致程序退出,所以两个一起用会有一个错误无法记录
	// util.WARN.Panicln(err2)
	util.ERROR.Fatalln(err3)
}

func IsNotNegative( n int) bool {
	return n > -1
}

// 命令行参数
func CmdArgs() {
	fmt.Printf("接收到了%v个参数", len(os.Args));
	for i, v := range os.Args {
		fmt.Printf("第%v各参数是%v\n", i, v);
	}
	fmt.Println()
	vPtr := flag.Bool("v", false, "Gonote版本")
	var userName string
	flag.StringVar(&userName, "u", "", "用户名")
	flag.Func("f", "", func(s string) error {
		fmt.Println("s=", s)
		return nil
	})
	flag.Parse()
	if *vPtr {
		fmt.Print("Gonote版本是 v0.0.0")
	}
	fmt.Println("当前用户为", userName)
}


// builtin package
func PackageBuiltin() {
	c1 := complex(12.34, 45.67)
	fmt.Println("c1 = ", c1)
	r1 := real(c1)
	i1 := imag(c1)
	fmt.Println("r1 = ", r1, "i1=", i1)
}

// runtime()
func PackageRuntime() {
	// 返回计算机逻辑处理数量
	if runtime.NumCPU() > 7 {
		// 允许go最多调几个处理器
		runtime.GOMAXPROCS(runtime.NumCPU() - 1 )
	}
}

// sync
func PackageSync() {
	var c int = 0
	var lock sync.Mutex
	var wg sync.WaitGroup
	// 线程
	primeNum := func(n int) {
		defer wg.Done()
		for i := 2; i < n; i++ {
			if n % i == 0 {
				return
			}
		}
		fmt.Printf("%v\t", n)
		lock.Lock()
		c++
		lock.Unlock()
	}
	
	for i := 2; i < 100001; i++ {
		wg.Add(1)
		go primeNum(i)
	}
	// 阻塞
	wg.Wait()
	fmt.Printf(",\n共找到%v", c)
}

// sync cond
func PackageSyncCond() {
	var lock sync.Mutex
	var wg sync.WaitGroup
	cond := sync.NewCond(&lock) // 提供了同时控制多个携程阻塞的能力
	for i := 0; i < 10; i++ {
		go func(n int) {
			cond.L.Lock()
			cond.Wait()
			fmt.Printf("携程%v,被唤醒了\n", n)
			cond.L.Unlock()
		}(i)
	}

	for i := 0; i < 15; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		if i == 4 {
			fmt.Println()
			cond.Signal()
		}
		if i == 9 {
			fmt.Println()
			cond.Broadcast()
		}
	}

	// once
	var once sync.Once
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			once.Do(func() {
				fmt.Println("只有一次机会")
			})
			wg.Done()
		}()
	}
	wg.Wait()

	// map
	var m sync.Map
	m.Store(1, 100)
	m.Store(2, 200)
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("m[%v] = %v\n", key, value.(int))
		return true	
	})
}

type Person struct {
	Name string
	Age int
}

type PersonSlice[] Person
func (ps PersonSlice) Len() int {
	return len(ps)
}
func (ps PersonSlice) Less(i, j int) bool {
	return ps[i].Age < ps[j].Age
}
func (ps PersonSlice) Swap(i,j int) {
	ps[i],ps[j] = ps[j],ps[i]
}
func PackageSort() {
	// 下标插入搜索
	is:=[]int{2,3,4,5,8, 10}
	v:=6
	i:=sort.SearchInts(is, v)
	fmt.Println("ss=",i)

	/// 自定义排序
	p := []Person{{"小小1", 18},{"小小0", 14},{"小小2", 16},}
	// sort.Slice(p, func(i, j int) bool {
	// 	return p[i].Age < p[j].Age
	// })
	fmt.Println("persons = ", p)
	// 自定义查找, 第一次出现不小于6的w
	sort.Search(len(is), func(i int) bool {
		return is[i] >= v
	})
	fmt.Printf("%v中第一次出现不小于%v的位置是%v\n", is, v, i)
	
	fmt.Println("sort interface")
	sort.Sort(sort.Reverse(PersonSlice(p)))
	fmt.Println("persons1 = ", p)
}