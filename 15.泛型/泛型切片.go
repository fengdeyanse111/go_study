package main
import "fmt"
// 自定义类型MySlice，同时也是泛型
type MySlice [T any] []T
func main(){
	var stringSlice MySlice[string]
	stringSlice = append(stringSlice, "枫枫")
	fmt.Println(stringSlice[0])

	var intSlice MySlice[int]
	intSlice = append(intSlice, 18)
	fmt.Println(intSlice[0])
}