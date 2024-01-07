package main

type Person struct {
	name string
	age  int
	sex  int
}

/* 继承 */
type Student struct {
	// 通过匿名字段实现继承操作
	// 结构体作为结构体成员
	Person
	// 重名字段信息，遵循就近原则，现在子类找，找不到去找父类
	name  string
	id    int
	score int
}

/* 指针匿名字段 */
type Student1 struct {
	*Person
	id    int
	score int
}

/*func main() {
	// var stu Student = Student{Person{"Zenos", 21, 1}, "Micheal", 1001, 90}
	// fmt.Println(stu.name)

	//var stu Student1
	//stu.Person = new(Person)
	//s-tu.name = "Zenos"
	//stu.age = 18
	//stu.sex = 1
	//stu.id = 1001
	//stu.score = 89
	//fmt.Println(stu)

}
*/
