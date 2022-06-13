package note

import "fmt"

const(
	Version = 1000
)
func SayHello() {
	fmt.Println("hello")
}

// 变量常量
func VarConst() {
	var v1 int
	v1 = 1
	var (
		v2 = 2
		v3 = 3
	)
	fmt.Printf("v1 = %v, v2 = %v v3 = %v\n", v1, v2, v3)
	const (
		c1 = 8
		c2 = iota 
		c3 // 默认为上一行的值
		c4
	)
	fmt.Printf("c1 = %v, c2 = %v, c3 = %v, c4 = %v\n", c1, c2, c3, c4)
}

// 基本数据类型
func BasicData() {
	var (
		n1 = 5
		n2 int8 = 127
		n3 uint16 
	)
	fmt.Printf("n1 = %v, type is %T,", n1, n1)
	fmt.Printf("n2 = %v, type is %T,", n2, n2)
	fmt.Printf("n3 = %v, type is %T\n", n3, n3)

	const (
		c1 = 5.1
		c2 float32 = 127
		c3 float32 = 1 
	)
	fmt.Printf("c1 = %v, type is %T,", c1, c1)
	fmt.Printf("c2 = %v, type is %T,", c2, c2)
	fmt.Printf("c3 = %v, type is %T\n", c3, c3)

	var (
		d1 byte
		d2 = 'a'
		d3 = "中"
	)
	fmt.Printf("d1 = %v, type is %T,", d1, d1)
	fmt.Printf("d2 = %v, type is %T,", d2, d2)
	fmt.Printf("d3 = %v, type is %T\n", d3, d3)

	var s1 = "hello"
	fmt.Println(s1,"world", len(s1))
}

// 指针
func increase(n *int) {
	*n++
	fmt.Printf("n = %v, type is %T\n", *n, *n)

}

func Pointer() {
	var src = 2022;
	increase(&src)
	fmt.Printf("n = %v, type is %T\n", src, src)

	var ptr = new(int)
	fmt.Printf("ptr = %v\n", ptr)
}

// 格式字符
func FmtVerbs() {
	fmt.Printf("%%\n");
	fmt.Printf("%.2f\n", 2.234234)
	fmt.Printf("%t\n", true)
	fmt.Printf("%q\n", "hello")
	fmt.Printf("%x\n", "hello")
}

// switch.case
func SwitchCase() {
	var weekday = 1
	fmt.Println("请输入你的年龄")
	fmt.Scanln(&weekday)
	switch weekday {
	case 1:
		fmt.Printf("%v", weekday)
		fallthrough
	case 2:
		fmt.Printf("%v", weekday)
	}
}

// for循环
func ForCase() {
	//无线循环
	for {
		fmt.Println("a")
	}
	i := 0
	for i < 10 {
		fmt.Println("abc")
		i++ 
	}

	for i := 0; i < 10; i++ {
		fmt.Println("cba")
	}
}

// label与goto
func LabelAndGoto() {
	outside:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print("+ ")
			if i == 9 && j == 4 {
				break outside
			}
		}
		fmt.Println()
	}
}
