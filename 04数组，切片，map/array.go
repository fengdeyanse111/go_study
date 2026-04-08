package main

import "fmt"

func main() {
	// 定义数组一定要先声明长度
	var array [3]int = [3]int{1, 2, 3}
	fmt.Println(array)
	var array1 = [3]int{1, 2, 3}
	fmt.Println(array1)
	var array2 = [...]int{1, 2, 3}
	fmt.Println(array2)

	// 如果要修改某个值，只能根据索引去找然后替换
	array[0] = 10
	fmt.Println(array)

	// 切片，与与数组定义很类似，只是不说明[]长度,append的返回值一定要有变量接受，否则会报错
	var sList = []string{"a", "b"}
	sList1 := append(sList, "e")
	fmt.Println(sList)
	fmt.Println(sList1)
}
