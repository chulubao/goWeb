package person

/**
[1]向对象有了三大特性:封装、继承、多态.
[2]只要是大写的就是共有的,只要是小写的都是私有的.
[3]封装特性.
*/

type Person struct{
	name  string
	age   int
}

func  (ptr *Person)SetName(name string){
	ptr.name = name
}

func  (ptr *Person)SetAge(age int){
	ptr.age = age
}

func  (ptr *Person)GetAge()int {
	return  ptr.age;
}

func (ptr *Person)GetName() string{
	return   ptr.name;
}
