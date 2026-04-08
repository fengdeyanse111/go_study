package main
import (
	"fmt"
	"math"
)
func main(){
	fmt.Printf("%.0f\n", math.Pow(2, 63))
	var n1 int = 9223372036854775807
	fmt.Println(n1)
	// var n2 int = 9223372036854775808
	// fmt.Println(n2)
}