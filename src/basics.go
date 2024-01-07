package main

import (
	"fmt"
	"os"
)

//import (
//	"fmt"
//	"testfunc"
//)
/* 变量的声明与定义 */
/*func main() {

	// 变量声明与定义
	// var 变量名 数据类型 = 值
	var a int = 25
	// 自动推导类型
	// 变量名 := 值
	b := true
	// 多重赋值
	// 变量1,变量2:=值1,值2
	c, d := 1, 2
	// 匿名变量
	_, i, _, j := 1, 2, 3, 4
	// 常量
	const PI float64 = 3.14159265
	// const PI = 3.14159265
	fmt.Println(PI)
	fmt.Println(a, b, c, d, i, j)
}*/

/* 输入输出 */
/*func main() {
	var age int
	fmt.Print("输入年龄：")
	fmt.Scanf("%d", &age)
	fmt.Printf("your age is %d", age)

}*/

/* iota枚举类型 */
/*func main() {
	// iota常量自动生成器，每个一行，自动累加1
	// iota给常量赋值使用
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a, b, c)
	// const会将iota重新置为0
	const d = iota
	// 写一个iota，后面的变量会自动赋值
	const (
		a1 = iota
		b1
		c1
	)
	fmt.Println(a1, b1, c1)
	// 同一行同时声明多个iota，其值一致
	const (
		a2         = iota
		b2, c2, d2 = iota, iota, iota
		e2         = iota
	)
	fmt.Println(a2, b2, c2, d2, e2)
}*/

/* 类型转换 */
/*func main() {
	// Go语言中没有隐式转换，所有类型转换必须显示声明，且转换只能发生在两种互相兼容的类型之间
	var ch byte = 97
	var a int = int(ch)
	fmt.Println(a)
}*/

/* 条件语句 */
/* if */
/*func main() {
	var age int = 18
	if age > 18 {
		fmt.Println("我已经成年了！")
	} else if age < 18 {
		fmt.Println("我还未成年！")
	} else {
		fmt.Println("我刚刚成年！")
	}
}*/
/* switch */
/*func main() {
	var score string = "B"
	switch score {
	case "A":
		fmt.Println("成绩优秀")
	case "B":
		fmt.Println("成绩良好")
	case "C":
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}
}*/

/* 循环 */
/*func main() {
	//for i := 0; i < 10; i++ {
	//	fmt.Println(" This is a test !")
	//}

	//var collection [10]int
	//for key, val := range collection {
	//	fmt.Println(key, val)
	//}

	var a int = 10
LOOP:
	for a < 20 {
		if a == 15 {
			// 跳过迭代
			a = a + 1
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}
}*/

/* 自定义函数 */
/*func function1(a int, b int) {
	fmt.Printf("a=%d,b=%d", a, b)
}*/

/* 不定长参数 */
/*func function2(a int, args ...int) {
	fmt.Printf("a=%d \n", a)
	for i := 0; i < len(args); i++ {
		fmt.Printf("args[%d]=%d \t", i, args[i])
	}
	fmt.Println('\n')
}*/

/* 返回值 */
/*func function3(a int) int {
	return a * 100
}
func function4(a int, b int, c int) (int, int, int) {
	return a, b, c
}
func main() {
	function1(1, 2)
	function2(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(function3(1))
	fmt.Println(function4(1, 2, 3))
}*/

/* 自定义函数变量 */
/*func function1(a int) int {
	return a
}

type functionType func(a int) int

func main() {
	var newName functionType
	newName = function1
	fmt.Println(newName(1))

}*/

/* 匿名函数 */
/*func main() {
	var num int = 9
	//f := func() {
	//	num++
	//	fmt.Println("匿名函数：", num)
	//}
	//// 调用函数
	//f()
	// 声明并直接调用
	func() {
		num++
		fmt.Println("匿名函数:", num)
	}()
	fmt.Println("main函数：", num)
}*/

/* 闭包 */
// 所谓的闭包是指有权访问另一个函数作用域中的变量的函数，就是在一个函数内部创建另一个函数。

/* 延迟调用defer */
// 关键字 defer ⽤于延迟一个函数（或者当前所创建的匿名函数）的执行.一般用于文件操作场景，为了保证文件的关闭能够正确执行，可以使用defer.
// 如果一个函数中有多个defer语句，它们会以LIFO（后进先出）的顺序执行。
/*func main() {
	defer fmt.Println("final")
	defer fmt.Println("second")
	defer fmt.Println("first")
	fmt.Println("normal")
}*/

/* 数组 */
/*func main() {
	// 定义
	// 初始化的元素均为0
	var a [3]int = [3]int{1, 2, 3}
	b := [3]int{2, 3, 4}
	c := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i], b[i], c[i][i])
	}
}*/

/* 切片 */
// 切片与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大，所以可以将切片理解成“动态数组”，但是，它不是数组。
/*func main() {
	// 定义
	// 与数组一致初始化的元素均为0
	var a []int = []int{1, 2, 3}
	// 使用make函数定义
	// 参数是: 类型 长度 容量
	s := make([]int, 2, 5)
	fmt.Println(len(a), cap(a))
	fmt.Println(len(s), cap(s))
	fmt.Println(s[0:2])
	// 切片追加,默认在已经生成的元素后面追加
	s = append(s, 1)
	fmt.Println(s[:])
}*/

/* Map */
// key value 结构，类似hash表
/*func main() {
	// 定义
	// 键的类型不可以使用复合类型
	// 声明后需要使用make方法
	var m map[string]int
	m = make(map[string]int, 1)
	m["first"] = 1
	m["second"] = 2
	fmt.Println(m, len(m))
}*/

/* 结构体 */
// 定义
/*type Student struct {
	id   int
	name string
}

func main() {
	// 初始化
	//var s Student
	//s.id = 1001
	//s.name = "zenos"

	// 按顺序赋值
	//s := Student{1, "zenos"}

	// 数组
	//stus := []Student{
	//	{1001, "Zenos"},
	//	{1002, "Micheal"},
	//}
	// println(stus[0].id)
}*/

/* 指针 */
/*func Swap(p *[5]int) {
	(*p)[0] = 89
}

type Student struct {
	id   int
	name string
}

func main() {
	// 定义
	//var pointer *int
	//i := 100
	//pointer = &i

	// 使用new函数初始化
	//pointer = new(int)
	//*pointer = 100
	//println(*pointer)

	// 数组指针
	//a := [5]int{1, 2, 3, 4, 5}
	//Swap(&a)
	//println(a[0])

	// 结构体指针
	var p *Student = &Student{1, "Zenos"}
	println((*p).id)

}*/

/* 异常 */
// error接口
/*func test(a int, b int) (int, error) {
	if b == 0 {
		err := errors.New("o不能作为除数")
		return 0, err
	}
	return a / b, nil
}*/

// panic函数
// error返回的是一般性的错误，但是panic函数返回的是让程序崩溃的错误。
// 当遇到不可恢复的错误状态的时候，如数组访问越界、空指针引用等引发panic异常
// 在一般情况下，我们不应通过调用panic函数来报告普通的错误，而应该只把它作为报告致命错误的一种方式。
/*func test(a int, b int) int {
	return a / b
}
func main() {
	val := test(10, 0)
	fmt.Println(val)

}*/

// recover接口
// 当前的程序从运行时panic的状态中恢复并重新获得流程控制权，程序能继续执行
// 只在defer调用中有效
/*func test(i int) {
	var arr [10]int
	// 通过匿名函数和recover进行错误拦截
	// 一定要在错误出现之前调用
	defer func() {
		// 返回值为接口类型
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	arr[i] = 100
}
func main() {
	test(11)
}*/

/* 文件处理 */

// 创建文件
/*func main() {

	fp, err := os.Create("./test.txt")
	// 创建失败
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	// 创建成功
	fmt.Println("文件创建成功")

	// 关闭文件
	defer fp.Close()
}*/

// 读取和写入文件
func main() {

	// 追加模式，编辑权限为777
	fp, err := os.OpenFile("./test.txt", os.O_APPEND, os.ModePerm)
	// 打开失败
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}
	// 打开成功
	fmt.Println("文件打开成功")

	// 写入
	// fp.Write() 写入字符切片
	// fp.WriteAt() 写入指定偏移量的字符切片(会覆盖原位置的字符)（配合Seek()方法查询文件字符长度）
	/* 还可以使用bufio包
	reader := bufio.NewReader(fp)
	for {
		buf,_,errReader := reader.ReadLine()
	}
	*/
	_, errWrite := fp.WriteString("This is a test !\n")
	// 写入失败
	if errWrite != nil {
		fmt.Println("文件写入失败")
		return
	}
	// 写入成功
	fmt.Println("写入成功")

	// 关闭文件
	defer fp.Close()
}
