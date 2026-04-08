package main
import "fmt"

func main(){
	fmt.Println("请输入年龄：")
	var age int
	t, ok := fmt.Scan(&age)
	fmt.Println(t, ok, age)

	// 分段判断
	if age <= 0{
		fmt.Println("没出生")
		return
	}
	if age <= 18{
		fmt.Println("没成年")
		return
	}
	if age <= 35{
		fmt.Println("青年")
		return
	}
	fmt.Println("中年")
}