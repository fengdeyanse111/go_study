package main
import "fmt"
func main(){
	// while()
	// doWhile()
	// traverseSlices()
	// traverseMaps()
	nineNineTable()
}
func while(){
	i := 0
	sum := 0
	for i <= 100{
		sum += i
		i += 1
	}
	fmt.Println(sum)
}
func doWhile(){
	i := 0
	sum := 0
	for{
		sum += i
		i += 1
		if i == 101{break}
	}
	fmt.Println(sum)
}
func traverseSlices(){
	s := []string{"xiaoqing", "jiajia"}
	for index, s2 := range s{
		fmt.Println(index, s2)
	}
}
func traverseMaps(){
	m := map[int]string{1:"xiaoqing", 2:"jiajia"}
	for key, val := range m{
		fmt.Println(key, val)
	}
}
func nineNineTable(){
	for i := 1; i <= 9; i++{
		for j := 1; j <= i; j++{
			fmt.Printf("%d * %d = %d\t", i, j, i * j)
		}
		fmt.Println()
	}
}