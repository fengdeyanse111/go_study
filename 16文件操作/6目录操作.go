package main
import (
	"fmt"
	"os"
)
func main(){
	// 目录也可以用for循环来遍历
	dir, _ := os.ReadDir("16文件操作")
	for _, entry := range dir{
		info, _ := entry.Info()
		fmt.Println(entry.Name(), info.Size())	//文件名，文件大小，单位比特
	}
}