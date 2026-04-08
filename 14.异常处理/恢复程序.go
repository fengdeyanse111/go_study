package main
import (
	"fmt"
	"runtime/debug"
)
func read(){
	defer func(){
		if err := recover(); err != nil{
			// 捕获异常，打印错误信息
			fmt.Println(err)
			// 打印错误的堆栈信息
			fmt.Println(string(debug.Stack()))
		}
	}()
	var list = []int{1,2}
	fmt.Println(list[2])
}
func main(){
	read()
	fmt.Println("主程序结束")
}