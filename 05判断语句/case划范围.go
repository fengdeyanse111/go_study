package main
import "fmt"
func main(){
	var age int
	fmt.Println("请输入你的年龄")
	fmt.Scan(&age)

	switch {
	case age <= 0:
		fmt.Println("没出生")
	case age <= 18:
		fmt.Println("未成年")
	case age <= 35:
		fmt.Println("青年")
	default:
		fmt.Println("中年")
	}
}