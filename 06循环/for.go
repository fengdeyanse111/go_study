package main
import "fmt"
import "time"
func main(){
	sum := 0
	for i := 0; i <= 100; i++{
		sum += i
	}
	fmt.Println(sum)
	deadCircle()
}
func deadCircle(){
	for{
		time.Sleep(1 * time.Second)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	}
}