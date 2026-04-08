package main
import (
	"fmt"
	"sync"
	"time"
)
// 容量为0的缓冲区，把它理解为PV操作里的资源，起到同步功能
var c chan int = make(chan int) 
func pay(s string, money int, wait *sync.WaitGroup){
	fmt.Printf("%s 开始购物\n", s)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 购物完毕\n", s)
	c <- money
	
	wait.Done()
}
func main(){
	var wait = sync.WaitGroup{}
	wait.Add(3)
	go pay("张三", 2, &wait)
	go pay("李四", 3, &wait)
	go pay("王五", 4, &wait)

	// 启动一个匿名函数
	go func(){
		defer  close(c)
		wait.Wait()
	}()		// 一定别忘了这个括号，代表调用这个函数
	var moneyList []int
	// for {
	// 	money, ok := <- c
	// 	fmt.Println(money, ok)
	// 	if !ok {
	// 		break
	// 	}
	// }

	// 另一种写法：for 循环不仅可以遍历分片，Map，还可以遍历 通道
	// 注意这里的range c在通道c关闭的时候不会额外调用一次c = 0的过程，所以不用担心moneyList里面有0
	for money := range c{
		moneyList = append(moneyList, money)
	}
	fmt.Println(moneyList)
	// 不能将close(c)写在这里！因为如果不在for循环之前close(c)的话，for循环里的<- c读取操作会一直等待着读取数据，因为它并不知道是否还有别的进程将会写数据，所以之前的go func很优雅
	// defer close(c)
}