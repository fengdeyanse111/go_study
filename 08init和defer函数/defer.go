package main

import "fmt"
func Func(){
	defer fmt.Println("defer2")
	fmt.Println("func")
	defer fmt.Println("defer1")
}
func main(){
	// defer虽然是紧跟在return 之前执行的，但是它不能使用在defer之后定义的变量
	// defer fmt.Println("defer4" + db)
	defer fmt.Println("defer4")
	var db int
	fmt.Println("main")
	Func()
	defer fmt.Println("defer3", db)
}