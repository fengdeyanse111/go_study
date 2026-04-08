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

func sing(obj Animal){
	obj.sing()
}
func main(){
	var animal Animal
	animal = Chicken{"ik"}

	animal.sing()
	animal.jump()
	animal.rap()
	cat := Cat{"ali"}
	sing(cat)
}