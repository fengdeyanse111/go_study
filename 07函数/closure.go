package main
import (
	"fmt"
	"time"
)
func main(){
	v := await(2)(1, 2, 3)
	fmt.Println(v)
}
func await(t int) func(...int) int{
	time.Sleep(time.Duration(t) * time.Second)
	return func(numList ...int) int{
		sum := 0
		for _, num := range numList{
			sum += num
		}
		return sum
	}
}