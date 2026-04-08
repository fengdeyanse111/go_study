package main
// 如果结构体有call这个名字的方法，就执行它
import (
	"fmt"
	"reflect"
)
type Student struct{
	Name string
	Age int
}
// 方法名大写才能导出，小写则不能导出。反射只能访问导出方法
func (Student) See(name string){
	fmt.Println("see name:", name)
}
func main(){
	s := Student{
		"fengfeng",
		18,
	}
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	for i := 0; i < t.NumMethod(); i++{
		methodType := t.Method(i)
		fmt.Println(methodType.Name, methodType.Type)
		if methodType.Name != "See"{
			continue
		}
		methodValue := v.Method(i)
		methodValue.Call([]reflect.Value{reflect.ValueOf("枫枫")})
	}
}