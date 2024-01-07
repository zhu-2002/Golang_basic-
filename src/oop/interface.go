package main

import "fmt"

type Human struct {
	name string
	sex  int
	age  int
}
type student struct {
	Human
	score int
}
type teacher struct {
	Human
	subject string
}

func (s *student) SayHello() {
	fmt.Printf("我是%s,我这次考了%d分，一起加油！\n", s.name, s.score)
}
func (s *teacher) SayHello() {
	fmt.Printf("我是%s,任教大家的%s学科，一起加油！\n", s.name, s.subject)
}

// 定义接口
// 接口实现多态，同种方法的不同表现形式
type SayHi interface {
	SayHello()
}

/*func main() {
	//var stu student = student{Human{"Zenos", 1, 17}, 89}
	//var tea teacher = teacher{Human{"Lily", 0, 30}, "math"}
	//stu.SayHello()
	//tea.SayHello()

	// 定义接口类型
	//var stuh SayHi = &stu
	//var ttuh SayHi = &tea
	//stuh.SayHello()
	//ttuh.SayHello()

	// 空接口
	// 可以切换成不同类型，可用于函数参数
	var i interface{}
	i = 1
	fmt.Println(i)
	i = "abd"
	fmt.Println(i)
}*/
