package main

import (
	"fmt"
	note "goproject/gonote"
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
	// notes.RandNum()
	// notes.StrConv()
	// notes.PackageString()
	// notes.PackageUFT8()
	// notes.Time()
	// notes.FileOperation()
	// notes.FileReadAndWrite()
	// notes.Errors()
	// notes.Log()
	// notes.CmdArgs()
	// notes.PackageBuiltin()
	// notes.PackageSync()
	// notes.PackageSyncCond()
	// notes.Recursion()
	// notes.Closure()
	// notes.Sort()
	// notes.PackageSort()
	// notes.BinarySearchText()
	// m := factory.NewMes()
	// m.SetPwd("asdf")
	// note.PackageJson()

	// note.TcpServer()
	// note.TcpCli()

	note.LevelDBBasic()
}