package main

import (
  "fmt"
  "io"
  "net/http"
  "strings"
)

func main() {
  // 1️⃣ 准备 JSON 数据
  jsonStr := `{"username":"张三"}`

  // 2️⃣ 创建请求（POST + Body）
  req, _ := http.NewRequest(
    "POST",
    "http://localhost:8080/index",
    strings.NewReader(jsonStr),
  )

  // 3️⃣ 设置请求头（非常重要！）
  req.Header.Set("Content-Type", "application/json")

  // 4️⃣ 创建客户端
  client := &http.Client{}

  // 5️⃣ 发送请求
  res, _ := client.Do(req)

  defer res.Body.Close()

  // 6️⃣ 读取响应
  body, _ := io.ReadAll(res.Body)

  // 7️⃣ 打印结果
  fmt.Println("响应内容:", string(body))

  // 8️⃣ 读取响应头（token）
  token := res.Header.Get("token")
  fmt.Println("token:", token)
}