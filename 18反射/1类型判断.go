package main
import (
	"fmt"
	"reflect"
)
// 通过反射判断类型
func refType(obj any){
	typeObj := reflect.TypeOf(obj)
	fmt.Println(typeObj, typeObj.Kind())
	// 判断具体的类型
	switch typeObj.Kind(){
	case reflect.Slice:
		fmt.Println("切片")
	case reflect.Map:
		fmt.Println("Map")
	case reflect.Struct:
		fmt.Println("结构体")
	case reflect.String:
		fmt.Println("字符串")
	}
}
// 通过反射获取值
func refValue(obj any){
	value := reflect.ValueOf(obj)
	fmt.Println(value, value.Type())
	switch value.Kind(){
	case reflect.Int:
		fmt.Println(value.Int())
	case reflect.Struct:
		fmt.Println(value.Interface())
	case reflect.String:
		fmt.Println(value.String())
	}
}
func main(){
	refType(struct{
		Name string
	}{Name:"枫枫"})
	name := "枫枫"
	refType(name)

	refValue(struct{
		Name string
	}{Name:"枫枫"})
	refValue(name)
}