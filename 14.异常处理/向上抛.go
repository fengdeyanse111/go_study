package main
import (
	"fmt"
	"errors"
)
func method() error {
	return errors.New("出错了")
}
func parent() error{
	err := method()
	return err
}
// 将错误交给上一级处理,一般是用于框架层，有些错误框架层面不能擅做决定，将错误向上抛不失为一个好的办法
func main(){
	fmt.Println(parent())
}