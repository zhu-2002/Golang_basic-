package main

import "fmt"

/* 方法 */
type Int int
type Worker struct {
	name string
	age  int
	sex  int
}
type Workermate struct {
	Worker
	roomNO int
}

// func （方法接收者）方法名（参数列表）返回值类型
// 方法接收者必须为已存在的数据类型别名，方法接收者就类似于对象实例
func (a Int) add(b Int) Int {
	return a + b
}

func (w Worker) PrintInfo() {
	fmt.Println(w.name, w.age, w.sex)
}
func (w *Worker) EditInfo(name string, age int, sex int) {
	// 结构体指针可以直接调用成员变量，无需（*w）
	w.name = name
	w.age = age
	w.sex = sex
}

// 方法重写
func (w Workermate) PrintInfo() {
	fmt.Println(w.name, w.age, w.sex, w.roomNO)
}

/*func main() {
	//var a Int = 10
	//fmt.Println(a.add(20))

	//w := Worker{"Zenos", 21, 1}
	//w.PrintInfo()
	//w.EditInfo("Micheal", 22, 1)
	//w.PrintInfo()

	//子类结构体继承父类结构体，允许使用父类结构体成员方法
	//w2 := Workermate{Worker{"Zenos", 21, 1}, 325}
	//w2.PrintInfo()

}*/
