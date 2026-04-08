package main
import "fmt"
func main(){
	// 这里不能只编译运行lala.go文件，因为这样会报错说找不到Add()函数，因为Add()函数是在calc.go文件里定义的，所以需要将calc.go联合起来一起编译才能找到Add()函数并顺利运行。正确运行方式：在命令行里：go run .
	fmt.Println(Add(1, 2))
}