package main
import (
	"fmt"
	"reflect"
)
// 结构体反射，读取json标签对应的值，如果没有就用属性的名称
type Student struct{
	Name string
	Age int `json:"age"`
}

func main(){
	s := Student{
		Name : "枫枫",
		Age : 18,
	}
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	for i := 0; i < t.NumField(); i++{
		field := t.Field(i)
		jsonField := field.Tag.Get("json")
		if jsonField == ""{
			jsonField = field.Name
		}
		fmt.Printf("Name: %s, Type: %s, json: %s, Value: %v\n", field.Name, field.Type, jsonField, v.Field(i))
	}
}