package main

import "fmt"

func main(){
	fmt.Println("请输入你的年龄:")
	var age int
	fmt.Scan(&age)
	if age <= 18{
		if age <= 0{
			fmt.Println("没出生")
		}else{
			fmt.Println("没成年")
		}
	}else{
		if age <= 35{
			fmt.Println("青年")
		}else{
			fmt.Println("中年")
		}
	}
}