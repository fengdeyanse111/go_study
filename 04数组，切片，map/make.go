package main

import (
	"fmt"
	"sort"
)

func main(){
	// 这里的len:代表当前切片已占用的容量大小，cap:代表在为切片分配空间时，为切片保证的连续空间的大小，如果说之后切片会扩展，直接在新空间处赋值新元素即可，而如果切片扩展超过了cap限定的大小，则会在内存中找到新的一片连续空间把原切片数据复制上去，这样的开销会很大
	list1 := make([]int, 0, 2)
	fmt.Println(list1, len(list1), cap(list1))
	fmt.Printf("%p \n", list1)
	list1 = append(list1, 1)
	list1 = append(list1, 2)
	fmt.Printf("%p \n", list1)
	list1 = append(list1, 3)
	fmt.Printf("%p \n", list1)

	// 切片可以直接从数组中切出来
	var list = [...]string{"a", "b", "c"}
	slices := list[:]
	fmt.Println(slices)
	fmt.Println(list[1:2]) // 左开右闭合，同python

	// 切片排序
	var list2 = []int{1, 4, 3, 9}
	fmt.Println("排序前", list2)
	sort.Ints(list2)
	fmt.Println("排序后", list2)
	sort.Sort(sort.Reverse(sort.IntSlice(list2)))
	fmt.Println(list2)
}