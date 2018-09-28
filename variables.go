package main

import "fmt"

var globalVal = "我是全局变量"

// 一起定义
var (
	globalBool bool
	globalInt int
	globalString string
)

func main() {
	fmt.Println("变量练习")

	// 变量声明，若没指定初始值，则为各类型默认的初始值
	// var variableName type, 定义一个名称为“variableName”，类型为"type"的变量
	var iAmBool bool
	var iAmInt int
	var iAmString string

	// var vname1, vname2, vname3 type  多个变量，一行定义

	// 变量声明，并初始化
	var iAmBool2 bool = false
	var iAmInt2 int = 15
	var iAmString2 string = "Hello World"

	// 不指定类型的初始化，go编译器自动识别
	var iAmBool3 = true
	var iAmInt3 = 15
	var iAmString3 = "Hello Golang"

	// var vname1,vname2 = 15, "hello world"  Go会根据其相应值的类型来帮你初始化它们

	// 简短声明变量, 只能定义局部变量
	iAmBool4 := false
	iAmInt4 := 17
	iAmString4 := "Hello Golang2"

	// vname1, vname2 := 16, "Hello world"  编译器会根据初始化的值自动推导出相应的类型

	_, b := 35, 24  // 下划线是个特殊的变量名，任何赋予它的值都会被丢弃

	fmt.Println("变量声明:", "iAmBool:", iAmBool, "iAmInt:", iAmInt, "iAmString:", iAmString)
	fmt.Println("变量声明，并初始化:", "iAmBool2:", iAmBool2, "iAmInt2:", iAmInt2, "iAmString2:", iAmString2)
	fmt.Println("不指定类型的初始化:", "iAmBool3:", iAmBool3, "iAmInt3:", iAmInt3, "iAmString3:", iAmString3)
	fmt.Println("简短声明变量:", "iAmBool4:", iAmBool4, "iAmInt4:", iAmInt4, "iAmString4:", iAmString4)
	fmt.Println("下划线:", b)
	fmt.Println("全局变量", globalVal)
	fmt.Println("分组定义全局变量", "globalBool:", globalBool, "globalInt:", globalInt, "globalString:", globalString)
}
