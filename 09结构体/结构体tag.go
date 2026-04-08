package main
import (
	"fmt"
	"encoding/json"
)

// 这里的json tag是为了将属性名改名，不写的话就会原样写入json文件，且属性名要大写开头
type Student struct{
	Name string `json:"name"`
	Age int		`json:"age"`
	Password string `json:"-"`	// 如果不想将某个字段转换出来的话就用"-"
	Num int 	`json:"num,omitempty"`	// 如果某个字段是空值的话也省略
}

func main(){
	s := Student{
		Name: "枫枫",
		Age: 18,
		Password: "123456",
		Num : 0,			// 空值会被省略
	}
	byteData, err := json.Marshal(s)
	fmt.Println(string(byteData))
	fmt.Println(err)
}