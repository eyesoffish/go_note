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

// deffer

func DefferFunc() {
	var sum = func(a, b int) int {
		return a + b
	}
	defer fmt.Println(sum(1, 0))
	defer fmt.Println(sum(2, 0))
	fmt.Println(sum(3, 0))

	// deferRecover 错误处理
	
	defer func ()  {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	n := 0
	fmt.Println(3 / n)
}

// 数组
func Slice() {
	array := [5]int{1,2,3,4,5}
	var s1 []int = array[0: len(array) - 1]
	s1[1] = 10
	fmt.Printf("array = %v\n", array)
	s2 := s1[1:]
	s2[0] = 0;
	fmt.Printf("array = %v\n", array)
	var s3 []int
	fmt.Println("s3 == true", s3 == nil)
	// s3为最小长度3, 最大长度5,, 
	s3 = make([]int, 3, 5)
	s1 = append(s1, 5,6,7,8)// 底层创建了新的数组,不再引用原数组
	fmt.Println(s1) 
	s4 := make([]int, 1)
	copy(s3, s4) // 容量能复制多少就接收多少
	fmt.Println("s4 = ", s4) 
	str := "hello 世界"
	for i,v := range str {
		fmt.Printf("str = %d, %c\n", i,v)
	}
}

// map map和数组默认为nil
func Map() {
	var m1 map[string]string
	fmt.Println(m1)
	m1 = make(map[string]string, 2)// 容量会自动增长, size如果省略, 长度默认为1
	m1["1"] = "m1"
	m1["2"] = "m2"
	m1["3"] = "m3"
	fmt.Println("m1 = ", m1)
	m2 := map[string]string{
		"4":"m4",
		"5":"m5",
	}
	fmt.Println("m2 = ", m2)

	v, ok := m2["14"]
	if ok {
		fmt.Println("key = 4 exist ", v)
	} else {
		fmt.Println("key = 4 not exist ", v == "")
	}
	delete(m1, "1")
	fmt.Println("m1 delete ", m1)
	// m1 = nil
	m2 = make(map[string]string)
	fmt.Println("m1 nil ", m1, m2)

	for key,value := range m1 {
		fmt.Println("key, value", key, value)
	}
}

// 自定义数据类型

func TypeDefineAndTypeAlias() {
	type mesType uint16
	var u1000 uint16 = 1000;
	var textMsg mesType = mesType(u1000)
	fmt.Printf("textMsg = %v, Type of textMes = %T\n", textMsg, textMsg)

	// 别名
	type myUint = uint16
	var  u2000 myUint = u1000
	fmt.Printf("textMsg = %v, Type of textMsg = %T\n", u2000, u2000)
}

// 结构体
type User struct {
	Name string `json:"name"`
	Id uint32	`json:"id"`
}
// 继承
type Account struct {
	User
	password string
}

func Struct() {
	var u1 User = User{Name: "张三", Id: 10000}
	fmt.Printf("user = %v", u1)
	var u2 Account = Account{User: User{Name: "李四", Id: 2000}, password:"666"}
	fmt.Printf("user2 = %v", u2)
}

// 方法
func (u User) printName() {
	fmt.Println(u.Name)
}

func (u *User) setId() {
	(*u).Id = 2000;
}

func Method() {
	u1 := User{Name: "张三", Id: 10000}
	u1.printName()
	u1.setId()
	fmt.Println("res == ",u1.Name, u1.Id)
}


// 接口
type textMsg struct {
	Text string
	Type string
}
func (tm *textMsg) setText() {
	tm.Text = "newText"
}
type imageMsg struct {
	Image string
	Type string
}
func (img *imageMsg) setImage() {
	img.Image = "清明上河图"
}

type Msg interface {
	setType()
}

func (tm *textMsg) setType(){tm.Type = "文字消息"}
func (tm *imageMsg) setType(){tm.Type = "图片消息"}

func SendMsg(m Msg) {
	m.setType()
	switch mptr:=m.(type) {
	case *textMsg:
		mptr.setText()
	case *imageMsg:
		mptr.setImage()
	}
	fmt.Println("m=",m)
}

func Interface() {
	tm:=textMsg{}
	im:=imageMsg{}
	SendMsg(&tm)
	SendMsg(&im)

	var n1 int = 1
	n1interface := interface{}(n1)
	n2,ok := n1interface.(string)
	if ok {
		fmt.Println("n2 = ", n2)
	} else {
		fmt.Println("n2 = ", n2)
	}
}