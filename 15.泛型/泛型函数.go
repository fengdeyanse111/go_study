package main
import "fmt"
func add[T int | float32 | int32](a, b T) T{
	return a + b
}
func main(){
	fmt.Println(add(10, 20))
	fmt.Println(add(float32(10), float32(20)))
	fmt.Println(add(int32(10), int32(20)))
	// 会报错
	// fmt.Println(add(int64(10), int64(20)))
}