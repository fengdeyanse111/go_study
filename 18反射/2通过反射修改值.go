package main
import (
	"fmt"
	"reflect"
)

func refSetValue(obj any){
	// 这里得到的是指针类型的反射对象
	value := reflect.ValueOf(obj)
	// 从指针取出它指向的值elem
	elem := value.Elem()
	switch elem.Kind(){
	case reflect.String:
		elem.SetString("峰峰知道")
	}
}

func main(){
	name := "枫枫"
	// 注意这里一定要传入指针
	refSetValue(&name)
	fmt.Println(name)
}