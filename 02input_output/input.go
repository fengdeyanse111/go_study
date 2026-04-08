package main

import "fmt"

func main(){
	fmt.Println("请输入你的名字")
	var name string
	fmt.Scan(&name)
	fmt.Println(name)

	fmt.Println("请输入你的年龄")
	var age int
	t, err := fmt.Scan(&age)
	fmt.Println(t, err, age)
}