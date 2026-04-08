package main
import (
	"fmt"
	"sync"
	"time"
)
// 容量为0的缓冲区，把它理解为PV操作里的资源，起到同步功能
var moneyChan = make(chan int) 
var nameChan = make(chan string)
var done = make(chan struct{})
func pay(s string, money int, wait *sync.WaitGroup){
	fmt.Printf("%s 开始购物\n", s)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 购物完毕\n", s)

	moneyChan <- money
	nameChan <- s
	
	wait.Done()
}
func main(){
	var wait sync.WaitGroup
	startTime := time.Now()
	wait.Add(3)
	go pay("张三", 2, &wait)
	go pay("李四", 3, &wait)
	go pay("王五", 4, &wait)

	// 启动一个匿名函数
	go func(){
		defer func(){
			close(nameChan)
			fmt.Println("nameChan已关闭")
		}()

		defer func(){
			close(moneyChan)
			fmt.Println("moneyChan已关闭")
		}()
		defer close(done)
		wait.Wait()
	}()		// 一定别忘了这个括号，代表调用这个函数
	var moneyList []int
	var nameList []string
	// 注意这里用event函数的原因：
	// 1.通过让主程序调用一个event函数，使得上面的匿名函数上的两个close()函数能够有时间执行
	// 2.如果直接用for 循环，case <-done:后面如果跟break,则只能跳出select而不能跳出for循环，而若后面跟return则与原因1相违背，因为return 会让主程序更快结束，更没有时间让协程close()
	event := func(){
		for{
			select{
			case money := <- moneyChan:
				moneyList = append(moneyList, money)
			case name := <- nameChan:
				nameList = append(nameList, name)
			case <-done:
				return
			}
		}
	}
	event()
	fmt.Println(nameList)
	fmt.Println(moneyList)
	fmt.Println("购买完成", time.Since(startTime))
}