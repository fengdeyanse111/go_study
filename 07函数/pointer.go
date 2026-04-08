package main

import "fmt"
func main(){
	num := 5
	fmt.Println(&num)
	valuePass(num)
	fmt.Println(num)

	addPass(&num)
	fmt.Println(num)
}
func valuePass(num int){
	fmt.Println(&num)
	num = 10
}
func addPass(num *int){
	fmt.Println(num)
	fmt.Println(&num)
	fmt.Printf("%T\n", num)
	*num = 20
}