package main
import (
	"fmt"
	"io"
	"os"
)
func main(){
	read, _ := os.Open("16文件操作/test.txt")
	write, _ := os.Create("16文件操作/test1.txt")	//注意这里是创建文件
	n, err := io.Copy(write, read)
	fmt.Println(n, err)
}