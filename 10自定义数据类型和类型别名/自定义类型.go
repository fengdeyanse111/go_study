package main
import "fmt"
type code int
const (
	SuccessCode code = 0
	ValidCode   code = 7		//校验失败的错误
	ServiceErrCode code = 8		//服务错误
)

func (c code)GetMsg() string{
	return "成功"
}

func main(){
	fmt.Println(SuccessCode.GetMsg())
	var i int
	fmt.Println(i == int(SuccessCode))
}