package main
import "fmt"

type Student struct{
	name string
	age int
}
// 这里当Student类型的对象调用setAge1时候也发生了值的拷贝！！！
func (s Student) setAge1(age int){
	s.age = age
}
func (s *Student) setAge2(age int){
	s.age = age
}

func main(){
	s := Student{
		name : "fengfeng",
		age : 18,
	}
	s.setAge1(20)
	fmt.Println(s.age)
	s.setAge2(22)
	fmt.Println(s.age)
}