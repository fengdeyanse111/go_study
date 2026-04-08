package main

import "fmt"

func main(){
	// 匿名函数
	var add = func(a, b int) int{
		return a + b
	}
	fmt.Println(add(10, 20))
	// 高阶函数
	m := map[int]func(){
		1:login,
		2:userCenter,
		3:logOut,
	}
	fmt.Println(`1：登录
	2：用户中心
	3：注销`)
	var num int
	fmt.Scan(&num)
	m[num]()
}
func login(){
	fmt.Println("登录")
}
func userCenter(){
	fmt.Println("用户中心")
}
func logOut(){
	fmt.Println("注销")
}