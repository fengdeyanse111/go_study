package main
import "fmt"
// 任何实现了以下三种函数的类型都可以是Animal
type Animal interface{
	sing()
	jump()
	rap()
}
// 实现接口：一个类型实现了该接口的所有方法，就说该类型实现了该方法
type Chicken struct{
	name string
}
func (c Chicken) sing(){
	fmt.Printf("Chicken %s sing\n", c.name)
}
func (c Chicken) jump(){
	fmt.Printf("Chicken %s jump\n", c.name)
}
func (c Chicken) rap(){
	fmt.Printf("Chicken %s rap\n", c.name)
}

type Cat struct{
	name string
}
func (c Cat) sing(){
	fmt.Printf("Cat %s sing\n", c.name)
}
func (c Cat) jump(){
	fmt.Printf("Chicken %s jump\n", c.name)
}
func (c Cat) rap(){
	fmt.Printf("Chicken %s rap\n", c.name)
}

// 通过断言来获取此时的具体类型
func sing(obj Animal){
	switch obj.(type){
	case Chicken:fmt.Println("鸡")
	case Cat:fmt.Println("猫")
	}
	obj.sing()
}
// 或者是断言某个类型
func jump(obj Animal){
	c, ok := obj.(Chicken)
	fmt.Println(c, ok)
	d := obj.(Cat)		// 如果只用一个参数接收，且类型不匹配的话会报错
	fmt.Println(d)
}
func main(){
	chicken := Chicken{"ik"}
	cat := Cat{"ali"}
	sing(chicken)
	sing(cat)
	// jump(chicken)
	jump(cat)
}