package main
// 修改结构体中某些值,例如，结构体tag中有big的标签，就将值大写
import (
	"fmt"
	"reflect"
	"strings"
)
type Student struct{
	Name1 string `big:"-"`
	Name2 string
}

func main(){
	s := Student{
		Name1 : "fengfeng",
		Name2 : "zhangsan",
	}
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(&s).Elem()
	for i := 0; i < t.NumField(); i++{
		field := t.Field(i)
		bigField := field.Tag.Get("big")
		if field.Type.Kind() != reflect.String || bigField == ""{
			continue
		}
		valueField := v.Field(i)
		valueField.SetString(strings.ToTitle(valueField.String()))
	}
	fmt.Println(s)
}
