package main
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)
func IndexHandler(res http.ResponseWriter, req *http.Request){
	switch req.Method{
	case "GET":
		data, err := os.ReadFile("19网络部署/index.html")
		if err != nil{
			fmt.Println(err)
		}
		res.Write(data)
	case "POST":
		data, err := io.ReadAll(req.Body)
		contentType := req.Header.Get("Content-Type")
		fmt.Println(contentType)
		if err != nil{
			fmt.Println(err)
		}
		ma := make(map[string]string)
		json.Unmarshal(data, &ma)
		fmt.Println(ma["username"])

		type User struct{
			Username string `json:"username"`
		}
		var user User
		json.Unmarshal(data, &user)
		// 设置响应头
		header := res.Header()
		header["token"] = []string{"sdjfalaf"}
		res.Write([]byte("hello 枫枫 POST"))
	}
}
func main(){
	// 1.绑定回调
	http.HandleFunc("/index", IndexHandler)
	// 2.注册服务
	fmt.Println("listen server on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}