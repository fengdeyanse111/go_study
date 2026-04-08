package main
import "fmt"
func main(){
	fmt.Println("请输入你的年龄:")
	var age int
	fmt.Scan(&age)

	switch{
	case age<=0:
		fmt.Println("没出生")
		fallthrough
	case age<=18:
		fmt.Println("没成年")
		fallthrough
	case age<=35:
		fmt.Println("青年")
		fallthrough
	default:
		fmt.Println("中年")
	}
}