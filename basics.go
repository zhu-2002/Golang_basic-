package main

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
func main() {

}
