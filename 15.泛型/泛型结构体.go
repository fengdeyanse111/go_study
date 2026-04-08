package main
import (
	"fmt"
	"encoding/json"
)
type Response[T any] struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data T `json:"data"`
}
type User struct{
	Name string `json:"name"`
}
type UserInfo struct{
	Name string `json:"name"`
	Age int `json:"age"`
}
func main(){
	// // 实例化需要指定泛型
	// user := Response[User]{
	// 	Code: 0,
	// 	Msg: "success",
	// 	Data: User{
	// 		Name: "枫枫",
	// 	},
	// }
	// byteData, _ := json.Marshal(user)
	// fmt.Println(string(byteData))
	// userInfo := Response[UserInfo]{
	// 	Code: 0,
	// 	Msg: "success",
	// 	Data: UserInfo{
	// 		Name: "枫枫",
	// 		Age : 18,
	// 	},
	// }
	// byteData1, _ := json.Marshal(userInfo)
	// fmt.Println(string(byteData1))
	var user = Response[User]{}
	// 这里是把字符串类型强制转换为[]byte类型，其中``符号代表使用原始字符串，否则"需要加转义字符\"来表示，比较麻烦，而至于为什么字符串的格式要是下面这样的，不知道怎么解释，但是从json.Marshal得到的字符串格式就是这样的，这两个是互逆过程
	json.Unmarshal([]byte(`{"code":0,"msg":"success","data":{"name": "枫枫"}}`), &user)
	fmt.Println(user.Data.Name)
	var userInfo = Response[UserInfo]{}
	json.Unmarshal([]byte(`{"code":0,"msg":"success","data":{"name":"枫枫","age":18}}`), &userInfo)
	fmt.Println(userInfo.Data.Name, userInfo.Data.Age)
}