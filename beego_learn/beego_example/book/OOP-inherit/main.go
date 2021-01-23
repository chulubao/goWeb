package main

import (
	"book/OOP/animal"
	"fmt"
)

/**
[1]向对象有了三大特性:封装、继承、多态.
[2]只要是大写的就是共有的,只要是小写的都是私有的.
*/
func main() {

	boy :=new(sex.Boy)
	boy.SetAge(12)
	boy.SetName("*zhuluji")
	boy.Addr="beijing"
	fmt.Println(boy.GetAge())
	fmt.Println(boy.GetName())
	fmt.Println(boy.Addr)

	gril :=new(sex.Gril)
	gril.SetAge(11)
	fmt.Println(gril.GetAge())/*就近原则*/
	gril.Ps.SetAge(23)
	fmt.Println(gril.Ps.GetAge())
	gril.Ps.SetName("zhang san")
	fmt.Println(gril.Ps.GetName())



}
