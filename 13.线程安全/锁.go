package main

import (
	"fmt"
	"sync"
)

var num int
var wait sync.WaitGroup
var lock sync.Mutex

func add(){
	lock.Lock()
	for i := 0; i < 100000; i++{
		num++
	}
	lock.Unlock()
	wait.Done()
}

func reduce(){
	lock.Lock()
	for i := 0; i < 100000; i++{
		num--
	}
	lock.Unlock()
	wait.Done()
}

func main(){
	wait.Add(2)
	go add()
	go reduce()
	wait.Wait()
	// 这个结果不一定是1，因为num++和num--并不是原子操作，都需要先读取num，再自增/自减，再写回
	// 操作系统中所说的并发控制问题
	fmt.Println(num)
}