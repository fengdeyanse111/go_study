package main
import (
	"fmt"
	"sync"
	"time"
)

var wait sync.WaitGroup
var mp = map[string]string{}

func reader(){
	for{
		fmt.Println(mp["time"])
	}
	wait.Done()
}
func writer(){
	for{
		mp["time"] = time.Now().Format("15:04:05")
	}
	wait.Done()
}
// 这里会报错：concurrent map read and map write
// 因为两个线程同时访问map["time"],且一个线程读，一个线程写，权限冲突了
func main(){
	wait.Add(2)
	go reader()
	go writer()
	wait.Wait()
}