package main
import "fmt"

func main(){
	// map在赋值之前一定要初始化，因为不初始化它只是nil，没有相应的方法
	var m1 map[string]string = map[string]string{}
	fmt.Println(m1 == nil)
	m1["风风"] = "知道"
	fmt.Println(m1)
	delete(m1, "风风")
	fmt.Println(m1)
	m1["凤凤"] = "知道"
	
	// 另一种初始化方法make
	var m2 map[string]string = make(map[string]string)
	fmt.Print(m2 == nil)
	age1 := m1["age1"]
	fmt.Println(age1)
	age1, ok := m1["age1"]
	fmt.Println(age1, ok)
}