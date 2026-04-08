package main
import (
	"fmt"
	"time"
	"sync"
)
var wait = sync.WaitGroup{}
func sing(){
	fmt.Println("唱歌")
	time.Sleep(1 * time.Second)
	fmt.Println("唱歌结束")
	wait.Done()
}
func main(){
	wait.Add(3)
	go sing()
	go sing()
	go sing()
	wait.Wait()
	fmt.Println("主线程结束")
}