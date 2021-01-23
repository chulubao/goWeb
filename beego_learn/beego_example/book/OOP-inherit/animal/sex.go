package sex

import "book/OOP/person"
/*匿名继承*/
type Boy struct{
person.Person   /*匿名结构体*/
Addr   string
}

/*有名继承:组合*/
type Gril struct{
	Ps  person.Person   /*有名继承:组合*/
	age  int
}

/*当结构体和匿名结构体有相同字段时，结构体变量采用就近原则来方法字段，如果希望采用匿名的字段，使用详细访问方式即可.*/
func  (ptr *Gril)SetAge(age int){
	ptr.age = age
}

func  (ptr *Gril)GetAge() int{
	return ptr.age
}