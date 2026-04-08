package main

import "fmt"

func main(){
	add(10, 20)
	add1(10, 20)
	add2([]int{10, 20, 30})
	add3(10, 20, 30)
}
func add(n1 int, n2 int){
	fmt.Println(n1 + n2)
}
func add1(n1, n2 int){
	fmt.Println(n1 + n2)
}
// 注意add2和add3在输入参数上的差别
func add2(numList []int){
	fmt.Printf("%T\n", numList)
	fmt.Println(numList)
}
func add3(numList ...int){
	fmt.Printf("%T\n", numList)
	fmt.Println(numList)
}