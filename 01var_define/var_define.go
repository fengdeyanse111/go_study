package main

import (
	"fmt"
	"01var_define/package1"
)
var (
	user1 string = "fengfeng"
	user2 = "zhidao"
)
func main(){
	// 先定义，在赋值
	var name string
	name = "ygy"
	fmt.Println(name)
	// 定义并赋值
	var name1 string = "ygy1"
	fmt.Println(name1)
	// 定义并赋值，省略类型说明
	var name2 = "ygy2"
	fmt.Println(name2)
	// 不使用var 直接定义并赋值
	name3 := "ygy3"
	fmt.Println(name3)

	// 定义多个变量
	var n1, n2, n3 string
	var a1, a2 = "fengfeng", "zhidao"
	a3, a4 := "fengfeng", "zhidao"
	fmt.Println(n1, n2, n3)
	fmt.Println(a1, a2)
	fmt.Println(a3, a4)

	// 定义全局变量
	fmt.Println(user1, user2)

	// 常量定义
	const c1 string = "凤凤"
	fmt.Println(c1)

	// 引用包里的变量
	fmt.Println(package1.Version)
	package1.Main()
}