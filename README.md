# Golang基础语法练习

> @author	Zenos
>
> 行百里者半九十

[TOC]



## 变量

### 变量类型

![image-20231226223359712](https://cdn.jsdelivr.net/gh/zhu-2002/img/image-20231226223359712.png)

![image-20231226223440808](https://cdn.jsdelivr.net/gh/zhu-2002/img/image-20231226223440808.png)

![image-20231226223504855](https://cdn.jsdelivr.net/gh/zhu-2002/img/image-20231226223504855.png)

![image-20231226223602422](https://cdn.jsdelivr.net/gh/zhu-2002/img/image-20231226223602422.png)

### rune类型

 rune本质是int32，一个rune四个字节。直观理解就是字符。

### 变量声明与定义

> 声明的变量首字母大写，可以在包外调用，小写只能在包内调用

```go
func main() {
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
}
```

## 输入输出

```go
func main() {
	var age int
	fmt.Print("输入年龄：")
	fmt.Scanf("%d", &age)
	fmt.Printf("your age is %d", age)

}
```

## iota枚举类型

```go
func main() {
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
}
```

## 强制类型转换

> golang是一个强类型语言，不会进行任何类型的隐式转换。

```go
func main() {
	// Go语言中没有隐式转换，所有类型转换必须显示声明，且转换只能发生在两种互相兼容的类型之间
	var ch byte = 97
	var a int = int(ch)
	fmt.Println(a)
}
```

## 条件语句

```go
/* if */
func main() {
	var age int = 18
	if age > 18 {
		fmt.Println("我已经成年了！")
	} else if age < 18 {
		fmt.Println("我还未成年！")
	} else {
		fmt.Println("我刚刚成年！")
	}
}
/* switch */
func main() {
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
}
```

## 循环语句

```go
func main() {
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
}
```

## 自定义函数

```go
/* 自定义函数 */
func function1(a int, b int) {
	fmt.Printf("a=%d,b=%d", a, b)
}

/* 不定长参数 */
func function2(a int, args ...int) {
	fmt.Printf("a=%d \n", a)
	for i := 0; i < len(args); i++ {
		fmt.Printf("args[%d]=%d \t", i, args[i])
	}
	fmt.Println('\n')
}

/* 返回值 */
func function3(a int) int {
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
}
```

## 自定函数变量

```go
func function1(a int) int {
	return a
}

type functionType func(a int) int

func main() {
	var newName functionType
	newName = function1
	fmt.Println(newName(1))

}
```

## 匿名函数

```go
func main() {
	var num int = 9
	//f := func() {
	//	num++
	//	fmt.Println("匿名函数：", num)
	//}
	// 调用函数
	//f()
	// 声明并直接调用
	func() {
		num++
		fmt.Println("匿名函数:", num)
	}()
	fmt.Println("main函数：", num)
}
```

## 闭包

所谓的闭包是指有权访问另一个函数作用域中的变量的函数，就是在一个函数内部创建另一个函数。**在Go语言里，所有的匿名函数(Go语言规范中称之为函数字面量)都是闭包。**

## 延迟调用defer

关键字 defer ⽤于延迟一个函数（或者当前所创建的匿名函数）的执行。一般用于文件操作场景，为了保证文件的关闭能够正确执行，可以使用defer。如果一个函数中有多个defer语句，它们会以LIFO（后进先出）的顺序执行。当函数中有返回值时不能使用defer调用

```go
func main() {
	defer fmt.Println("final")
	defer fmt.Println("second")
	defer fmt.Println("first")
	fmt.Println("normal")
}
```

## 工程管理

为了更好的管理项目中的文件，要求将文件都要放在相应的文件夹中。GO语言规定如下的文件夹如下：

- src目录：用于以代码包的形式组织并保存Go源码文件。（比如：.go .c .h .s等）

- pkg目录：用于存放经由go install命令构建安装后的代码包（包含Go库源码文件）的“.a”归档文件。

- bin目录：与pkg目录类似，在通过go install命令完成安装后，保存由Go命令源码文件生成的可执行文件。

以上目录称为工作区，工作区其实就是一个对应于特定工程的目录。目录src用于包含所有的源代码，是Go命令行工具一个强制的规则，而pkg和bin则无需手动创建，如果必要Go命令行工具在构建过程中会自动创建这些目录。

### 同级目录调包设置

- 开启gomod模式

  ```shell
  go env -w GO111MODULE=on
  ```

- 生成go.mod文件

  ```shell
  go mod init moduleName
  ```

- 调用包

  在import的时候src可以缺省

  ![image-20231227100641160](https://cdn.jsdelivr.net/gh/zhu-2002/img/image-20231227100641160.png)

  ![image-20231227100829338](https://cdn.jsdelivr.net/gh/zhu-2002/img/image-20231227100829338.png)

## 复合类型

### 数组

```go
func main() {
	// 定义
	// 初始化的元素均为0
	var a [3]int = [3]int{1, 2, 3}
	b := [3]int{2, 3, 4}
	c := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i], b[i], c[i][i])
	}
}
```

### 切片

```go
// 切片与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大，所以可以将切片理解成“动态数组”，但是，它不是数组。
func main() {
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
}
```

### Map

```go
// key value 结构，类似hash表
func main() {
	// 定义
	// 键的类型不可以使用复合类型
	// 声明后需要使用make方法
	var m map[string]int
	m = make(map[string]int, 1)
	m["first"] = 1
	m["second"] = 2
	fmt.Println(m, len(m))
}
```

### 结构体

```go
type Student struct {
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
}
```

### 指针

```go
func Swap(p *[5]int) {
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

}
```

## 泛型

## 面向对象

### 继承

```go
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
```

### 方法

```go
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
```

### 接口

```go
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
```

## 异常处理

### error接口

```go
func test(a int, b int) (int, error) {
	if b == 0 {
		err := errors.New("o不能作为除数")
		return 0, err
	}
	return a / b, nil
}
```

### panic函数

```go
// error返回的是一般性的错误，但是panic函数返回的是让程序崩溃的错误。
// 当遇到不可恢复的错误状态的时候，如数组访问越界、空指针引用等引发panic异常
// 在一般情况下，我们不应通过调用panic函数来报告普通的错误，而应该只把它作为报告致命错误的一种方式。
func test(a int, b int) int {
	return a / b
}
func main() {
	val := test(10, 0)
	fmt.Println(val)

}
```

### recover接口

```go
// 当前的程序从运行时panic的状态中恢复并重新获得流程控制权，程序能继续执行
// 只在defer调用中有效
func test(i int) {
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
}
```

## 文件操作

### 创建

```go
// 创建
func main() {

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
}
```

### 读写

```go
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
```



















## 附录

> 对应源码：
>
> 参考文章链接：
>
> - author：菜鸟教程	url：[Go 语言数据类型 | 菜鸟教程 (runoob.com)](https://www.runoob.com/go/go-data-types.html)
>
> - author：[QianLiStudent](https://blog.csdn.net/QianLiStudent)	url：http://t.csdnimg.cn/Onr44