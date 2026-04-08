package main
import "fmt"

// 自定义类型
type MyCode int
// 类型别名(与自定义类型很像，就是多了个=)
type AliceCode = int

const (
	SuccessMyCode MyCode = 0
	SuccessAliceCode AliceCode = 0
)

// 自定义类型可以绑定方法
func (mc MyCode) GetMyCode(){

}
// // 类型别名不能绑定方法
// func (ac AliceCode) GetAliceCode(){

// }
// func (c int) GetIntCode(){

// }
func main(){
	var i int
	// 自定义类型不能直接与原始类型比较，必须要先转换成原始类型.或者将原始类型转换成自定义类型
	// 即这两者是两个不同的类型
	fmt.Println(int(SuccessMyCode) == i)
	fmt.Println(MyCode(i) == SuccessMyCode)
	// 类型别名可以直接与原始类型比较
	fmt.Println(i == SuccessAliceCode)
}