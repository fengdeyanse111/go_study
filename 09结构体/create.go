package main
import "fmt"

type Student struct{
	name string
	age int
}
// 给结构体绑定一个方法
func (s Student) printInfo(){
	fmt.Printf("名字是：%s, 年龄是：%d\n", s.name, s.age)
}

func main(){
	s := Student{
		name: "哲哲",
		age: 18,
	}
	fmt.Println(s)
	s.printInfo()
	s.name = "校庆"
	s.age = 20
	s.printInfo()
}