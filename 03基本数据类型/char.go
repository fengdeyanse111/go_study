package main
import "fmt"
import "unsafe"
func main(){
	// 单字节字符byte 完全等价于 uint8
	var n1 byte = 'a'
	var n2 uint8 = 97
	fmt.Println(n1, n2)
	fmt.Printf("%c, %d\n", n1, n2)
	fmt.Printf("%d, %c\n", n1, n2)
	
	// 多字节字符rune
	var n3 rune = '中'
	var n4 rune = 'c'
	fmt.Printf("%c, %d, %c, %d\n", n3, n3, n4, n4)
	fmt.Println(unsafe.Sizeof(n3))
	fmt.Println(unsafe.Sizeof(n4))
}