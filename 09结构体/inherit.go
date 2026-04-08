package main
import "fmt"

type People struct{
	time string
}

func (p People) pInfo(){
	fmt.Println(p.time)
}

type Student struct{
	People
	name string
	age int
}

func (s Student) printInfo(){
	fmt.Printf("name:%s, age:%d\n", s.name, s.age)
}

func main(){
	p := People{
		time : "2026-04-03 20:12",
	}
	fmt.Println(p)
	p.pInfo()
	s := Student{
		People: p,
		name: "fengfeng",
		age: 18,
	}
	fmt.Println(s)
	s.printInfo()
	// 子类可以直接或者通过父类间接地访问父类结构体的属性
	fmt.Println(s.People.time, s.time)
	// 子类可以直接或者通过父类间接地访问父类结构体的方法
	s.People.pInfo()
	s.pInfo()
}