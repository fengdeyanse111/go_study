package main
import "fmt"
import "errors"
func main(){
	// // 会报错
	// v1 := fun1()
	// fmt.Println(v1)
	// 正常
	v2 := fun2()
	fmt.Println(v2)
	// 多返回值
	v3, err := fun3()
	fmt.Println(v3, err)
	// 自动返回明明返回值
	v4 := fun4()
	fmt.Println(v4)
	
}
func fun1(){
	return
}
func fun2() int{
	return 1
}
func fun3() (int, error){
	return 0, errors.New("错误")
}
// 命名返回值
func fun4() (res string){
	res = "abc"
	return
}