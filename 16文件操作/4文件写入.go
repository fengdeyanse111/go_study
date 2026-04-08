package main
import (
	"fmt"
	"os"
)
func main(){
	err := os.WriteFile("16文件操作/test.txt", []byte("这是写入内容"), os.ModePerm)
	fmt.Println(err)
}