package main
import (
	"fmt"
	"testing"
	"os"
)

// test函数也没有main()函数体，所以需要在命令行运行 go test，系统会自动编译属于main包的所有.go文件，并找到_test.go结尾的测试文件，并运行里面的形如TestXxx的测试函数
func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Errorf("1 + 2 expected be 3, but %d got", ans) // 不需要最后加\n，因为打印时候会自动换行
	}
	// 注意这里and := 不会报错，因为在if里面定义的ans作用域只属于if{}里。
	if ans := Add(2, 3); ans != 5 {
		t.Errorf("2 + 3 expected be 5, but %d got", ans)
	}
	// t.Log("TessAdd Pass!")	// 如果顺利通过的话不会打印t.Log。所以没必要写
}

// 子测试.如果需要给一个函数，调用不同的测试用例，可以使用子测试。子测试里面的Fatal，是不会终止程序的
func TestAdd1(t1 *testing.T) {
	t1.Run("add1", func(t *testing.T) {
		if ans := Add(10, 20); ans != 30 {
			t.Fatalf("10 + 20 expected be 30, but %d got", ans)
		}
	})
	t1.Run("add2", func(t *testing.T) {
		if ans := Add(-1, 1); ans != 0 {
			t.Fatalf("-1 + 1 expected be 0, but %d got", ans)
		}
	})
}

// 如果测试用例很多，可以用类似表格处理
func TestAdd2(t *testing.T) {
	cases := []struct {
		Name    string
		a, b, c int
	}{
		{Name: "a1", a: 1, b: 2, c: 3},
		{Name: "a2", a: 10, b: 20, c: 30},
		{Name: "a3", a: 100, b: 200, c: 300},
	}
	for _, s := range cases {
		t.Run(s.Name, func(t1 *testing.T) {
			if ans := Add(s.a, s.b); ans != s.c {
				t1.Fatalf("%d + %d expected be %d, but %d got", s.a, s.b, s.c, ans)
			}
		})
	}
}

func setUp(){
	fmt.Println("测试前执行")
}

func tearDown(){
	fmt.Println("测试后执行")
}

// 必须叫这个名字，测试主入口
func TestMain(m *testing.M){
	//测试前执行
	setUp()
	//进行测试
	code := m.Run()
	// 测试后执行
	tearDown()
	os.Exit(code)
}