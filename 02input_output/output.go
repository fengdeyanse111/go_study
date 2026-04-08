package main

import "fmt"

func main(){
	fmt.Println("fengfeng")
	fmt.Println(true)
	fmt.Println(1)
	fmt.Println("什么", "都", "可以", "输出")

	fmt.Printf("%v \n", "anytype is ok") // %v :可作为任何类型值的占位符输出
	fmt.Printf("%T, %T, %T \n", "fengfeng", 123, 1.23) // %T :输出类型
	fmt.Printf("%d, %.1f, %s \n", 123, 1.23, "fengfeng")

	name := fmt.Sprintf("%d, %v", 123, "fengfeng")	// fmt.Sprintf不打印结果，将结果转换成字符串返回
	fmt.Println(name)
	fmt.Printf("%p\n", &name) // 地址
	fmt.Println(&name)
}