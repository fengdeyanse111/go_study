package main

import "fmt"

func main(){
	// Slice如果没有初始化，可以直接向里面添加元素，系统会自动帮助初始化，所以虽然未初始化的Slice值是nil，在添加元素的时候系统会自动将其初始化为Slice类型的空元素。
	// 但是系统不会给map自动初始化，map 的赋值是语法操作，不会隐式初始化，因此必须手动 make。
	var list []string
	fmt.Println(list == nil)
	fmt.Printf("%T\n", list)
	list = append(list, "哲哲")
	fmt.Println(list)

	list1 := append([]string(nil), "哲哲")
	fmt.Println(list1)
	
	// // 会报错
	// var mp map[string]string
	// fmt.Println(mp == nil)
	// mp["陈陈"] = "哲哲"
	// fmt.Println(mp)
}