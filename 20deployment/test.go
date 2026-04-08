package main
import (
	"fmt"
)
func main(){
	fmt.Println("hello world")
	var s string
	fmt.Println("请任意输入,q代表退出")
	for{
		fmt.Scan(&s)
		if s == "q"{
			break
		}
		fmt.Println(s)
	}
}