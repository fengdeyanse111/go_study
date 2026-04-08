package main
import "fmt"
func main(){
	var week int
	fmt.Println("请输入周几：")
	fmt.Scan(&week)

	switch week {
	case 1:fmt.Println("Monday")
	case 2:fmt.Println("Tuesday")
	case 3:fmt.Println("wednesday")
	case 4:fmt.Println("Thursday")
	case 5:fmt.Println("Friday")
	case 6:fmt.Println("Saturday")
	case 7:fmt.Println("Sunday")
	default:fmt.Println("Wrong Input")
	}
}