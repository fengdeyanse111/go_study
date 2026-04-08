package main
import (
	"fmt"
	"time"
)

func sing(){
	fmt.Println("唱歌")
	time.Sleep(1 * time.Second)
	fmt.Println("唱歌结束")
}

func main(){
	go sing()
	go sing()
	go sing()
	// 停止几秒钟，否则主程序直接结束会结束协程
	time.Sleep(2 * time.Second)
}