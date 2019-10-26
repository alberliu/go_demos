package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type Persons []Person

// 获取此 slice 的长度
func (p Persons) Len() int { return len(p) }

// 根据元素的年龄降序排序 （此处按照自己的业务逻辑写）
func (p Persons) Less(i, j int) bool {
	return p[i].Age > p[j].Age
}

// 交换数据
func (p Persons) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func main() {
	persons := Persons{
		{
			Name: "test1",
			Age:  20,
		},
		{
			Name: "test2",
			Age:  22,
		},
		{
			Name: "test3",
			Age:  21,
		},
	}

	fmt.Println("排序前")
	for _, person := range persons {
		fmt.Println(person.Name, ":", person.Age)
	}
	sort.Sort(persons)
	fmt.Println("排序后")
	for _, person := range persons {
		fmt.Println(person.Name, ":", person.Age)
	}
}
