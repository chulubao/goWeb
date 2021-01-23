package main

import (
	"book/OOP/person"
	"fmt"
)

/**
[1]向对象有了三大特性:封装、继承、多态.
[2]只要是大写的就是共有的,只要是小写的都是私有的.
*/
func main() {
	p :=  person.Person{}
	p.SetName("zhu su juan")
	p.SetAge(24)
	fmt.Println(p.GetName())
	fmt.Println(p.GetAge())
}
