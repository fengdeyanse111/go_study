package main
import "fmt"
// 通道是由写入方和读取方共同协作操纵的，且通道有两种状态：打开和关闭
// 写入方只能在通道打开时进行操作，而读取方可以在通道打开时和关闭时都正常工作
func main(){
	var c chan int = make(chan int, 1)
	// 向c中写入数据
	c <- 1
	// // 通道已满，再写会堵塞
	// c <- 2
	fmt.Println(<- c)
	c <- 2
	num, ok := <- c
	fmt.Printf("num:%d, ok:%v\n", num, ok)
	// // 通道未关闭且通道已空，再读取就会报错
	// num1, ok1 := <- c
	// fmt.Printf("num:%d, ok:%v\n", num1, ok1)
	c <- 1
	close(c)
	// 通道已关闭且通道已空，再读取则会返回
	// 通道关闭的意义是让写入方不能再向通道里面写入数据了，并不影响读取方的操作
	num1, ok1 := <- c
	fmt.Printf("num:%d, ok:%v\n", num1, ok1)
	num2 := <- c
	fmt.Println(num2)
}