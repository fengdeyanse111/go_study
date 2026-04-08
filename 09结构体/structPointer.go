package main
import "fmt"
type Student struct{
	name string
	age int
}

func valuePass(s Student, age int){
	s.age = age
}

func addressPass(s *Student, age int){
	// 注意这里直接用的s.age，而没有用s->age
	s.age = age
}

func main(){
	s := Student{
		name : "枫枫",
		age : 18,
	}
	fmt.Println(s.age)
	valuePass(s, 20)
	fmt.Println(s.age)
	addressPass(&s, 22)
	fmt.Println(s.age)
}