// 遇到错误直接停止程序,这种一般是用于初始化，一旦初始化出现错误，程序继续走下去也意义不大了，还不如中断掉
package main
import (
	"fmt"
	"os"
)

func init(){
	_, err := os.ReadFile("xxx")
	if err != nil{
		panic(err.Error())	// 直接让程序崩溃，并输出错误信息
	}
}

func main(){
	fmt.Println("主函数")
}