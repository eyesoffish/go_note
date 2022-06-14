package main

import (
	"fmt"
	notes "goproject/gonote"
)

func init() {
	fmt.Println("初始化函数执行顺序")
	fmt.Println("被依赖包全局变量 -> 被依赖包init函数 -> main函数全局变量 -> main函数init -> main函数")
}

func main() {
	// 第一部分语法
	// notes.SayHello()
	// notes.VarConst()
	// println(notes.Version)
	// notes.BasicData()
	// notes.Pointer()
	// notes.FmtVerbs()
	// // note.SwitchCase()
	// // note.LabelAndGoto()
	// notes.DefferFunc()
	// notes.Slice()
	// notes.Map()
	// notes.TypeDefineAndTypeAlias()
	// notes.Struct()
	// notes.Method()
	// notes.Interface()
	// // notes.Goroutine()
	// notes.Channel()

	// 第二部分 
	notes.RandNum()
	notes.StrConv()
}