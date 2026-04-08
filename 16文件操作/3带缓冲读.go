package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("16文件操作/test.txt")
	buf := bufio.NewReader(file)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			break
		}
		fmt.Println(string(line))
	}

	// 指定分隔符
	// 这里是另一个例子，注意NewScanner和NewReader打开的都是文件file，由于第一次打印会将文件指针移动到文件末尾，所以在进行第二次读文件之前需要将文件指针重置到文件开头
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords) // 按照单词读
	// scanner.Split(bufio.ScanLines) // 按照行读
	// scanner.Split(bufio.ScanRunes) // 按照中文字符读
	// scanner.Split(bufio.ScanBytes) // 按照字节读，中文会乱码
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
