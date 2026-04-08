package main
import "fmt"
type MyMap[K string|int, V any] map[K]V
func main(){
	var MyMap = make(MyMap[string, string])
	MyMap["名字"] = "枫枫"
	fmt.Println(MyMap)
}