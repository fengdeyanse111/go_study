package main
import (
	"fmt"
	"os"
	"runtime"
)
// 获取当前go文件的路径。可以通过获取当前go文件的路径，然后用相对于当前go文件的路径去打开文件
func getCurrentFilePath() string{
	_, file, _, _ := runtime.Caller(1)
	return file
}
func main(){
	fmt.Println(getCurrentFilePath())
	// 绝对路径是：e:/Language/golang/go_study/16文件操作/test.txt
	// Go 的相对路径是基于“当前工作目录”，而不是 .go 文件所在目录,当前工作目录是e:/Language/golang/go_study/
	byteData, ok := os.ReadFile("16文件操作/test.txt") // 相对路径
	fmt.Println(ok)
	fmt.Println(string(byteData))
}