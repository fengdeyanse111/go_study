package main
import (
	"fmt"
	"time"
)
var done = make(chan struct{})
func event(){
	fmt.Println("协程开始执行")
	time.Sleep(2 * time.Second)
	fmt.Println("协程结束执行")
	close(done)
}
func main(){
	go event()
	select{
	case <-done:
		fmt.Println("协程正常结束")
	case <- time.After(1 * time.Second):
		fmt.Println("超时")
	}
	return
}