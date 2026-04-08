package main
import "fmt"
func main(){
	fmt.Println(`"1:登录
	2:用户中心
	3:注销"`)
	var num int
	fmt.Scan(&num)

	funcmap := map[int]func(){
		1: func(){
			fmt.Println("登录")
		},
		2: func(){
			fmt.Println("用户中心")
		},
		3: func(){
			fmt.Println("注销")
		},
	}
	funcmap[num]()
}