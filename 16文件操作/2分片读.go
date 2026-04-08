package main
import (
	"fmt"
	"os"
	"io"
)
func main(){
	file, _ := os.Open("16文件操作/test.txt")
	defer file.Close()
	for{
		buf := make([]byte, 3)
		_, err := file.Read(buf)
		if err == io.EOF{
			break
		}
		fmt.Printf("%s", buf)
	}
	fmt.Println()
}