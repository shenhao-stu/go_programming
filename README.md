# Go 语言自学系列

[TOC]

## byte 和 rune 类型

组成每个字符串的元素叫做“字符”，可以通过遍历或者单个获取字符串元素获得字符。 字符用单引号（'）包裹起来，如：

```go
package main

import "fmt"

func main() {
    var a = '华'
    var b = 'a'
    fmt.Printf("a: %v,%c\n", a, a)
    fmt.Printf("b: %v,%c\n", b, b)
}
```

运行结果

```
a: 21326,华
b: 97,a
```

Go 语言的字符有以下两种：

- uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符。
  - byte 类似 uint8
- rune 类型，代表一个 UTF-8 字符。
  - rune 类似 int32
  - 当需要处理中文、日文或者其他复合字符时，则需要用到 rune 类型。rune 类型实际是一个 int32。

Go 使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode 的文本处理更为方便，也可以使用 byte 型进行默认字符串处理，性能和扩展性都有照顾。

## 占位符

### 整数占位符

| 占位符 |                     说明                     |         举例         |  输出  |
| :----: | :------------------------------------------: | :------------------: | :----: |
|   %b   |                  二进制表示                  |   Printf("%b", 5)    |  101   |
|   %c   |        相应 Unicode 码点所表示的字符         | Printf("%c", 0x4E2D) |   中   |
|   %d   |                  十进制表示                  |  Printf("%d", 0x12)  |   18   |
|   %o   |                  八进制表示                  |   Printf("%o", 10)   |   12   |
|   %q   | 单引号围绕的字符字面值，由 Go 语法安全地转义 | Printf("%q", 0x4E2D) |  '中'  |
|   %x   |       十六进制表示，字母形式为小写 a-f       |   Printf("%x", 13)   |   d    |
|   %X   |       十六进制表示，字母形式为大写 A-F       |   Printf("%x", 13)   |   D    |
|   %U   |    Unicode 格式：U+1234，等同于 "U+%04X"     | Printf("%U", 0x4E2D) | U+4E2D |

### 浮点数和复数的组成部分（实部和虚部）

| 占位符 |                                                 说明                                                 |          举例          |     输出     |
| :----: | :--------------------------------------------------------------------------------------------------: | :--------------------: | :----------: |
|   %b   | 无小数部分的，指数为二的幂的科学计数法，与 strconv.FormatFloat 的 'b' 转换格式一致。例如 -123456p-78 |                        |              |
|   %e   |                                    科学计数法，例如 -1234.456e+78                                    |   Printf("%e", 10.2)   | 1.020000e+01 |
|   %E   |                                    科学计数法，例如 -1234.456E+78                                    |   Printf("%e", 10.2)   | 1.020000E+01 |
|   %f   |                                    有小数点而无指数，例如 123.456                                    |   Printf("%f", 10.2)   |  10.200000   |
|   %g   |                        根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的 0）输出                        |  Printf("%g", 10.20)   |     10.2     |
|   %G   |                        根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的 0）输出                        | Printf("%G", 10.20+2i) |  (10.2+2i)   |

### 字符串与字节切片

| 占位符 |                   说明                   |              举例              |     输出     |
| :----: | :--------------------------------------: | :----------------------------: | :----------: |
|   %s   |   输出字符串表示（string 类型或[]byte)   | Printf("%s", []byte("多课网")) |    多课网    |
|   %q   | 双引号围绕的字符串，由 Go 语法安全地转义 |     Printf("%q", "多课网")     |   "多课网"   |
|   %x   |    十六进制，小写字母，每字节两个字符    |     Printf("%x", "golang")     | 676f6c616e67 |
|   %X   |    十六进制，大写字母，每字节两个字符    |     Printf("%X", "golang")     | 676F6C616E67 |

### 指针

| 占位符 |         说明          |        举例         |   输出   |
| :----: | :-------------------: | :-----------------: | :------: |
|   %p   | 十六进制表示，前缀 0x | Printf("%p", &site) | 0x4f57f0 |

## 快捷键

`fint`
`ff`
`fp`
`forr`
`s.print!`
`s.var!`
`pkm`
`tyi`

## go 语言类型定义和类型别名的区别

- 类型定义相当于定义了一个全新的类型，与之前的类型不同；但是类型别名并没有定义一个新的类型，而是使用一个别名来替换之前的类型
- 类型别名只会在代码中存在，在编译完成之后并不会存在该别名
- 因为类型别名和原来的类型是一致的，所以原来类型所拥有的方法，类型别名中也可以调用，但是如果是重新定义的一个类型，那么不可以调用之前的任何方法

### go 语言类型定义

类型定义的语法

```python
type NewType Type
```

```go
package main

import "fmt"

func main() {
    // 类型定义
    type MyInt int
    // i为MyInt类型
    var i MyInt
    i = 100
    fmt.Printf("i: %v i: %T\n", i, i)
}

// i: 100 i: main.MyInt
```

### go 语言类型别名

类型别名的语法

```python
type NewType = Type
```

```go
package main

import "fmt"

func main() {
    // 类型别名定义
    type MyInt2 = int
    // i其实还是int类型
    var i MyInt2
    i = 100
    fmt.Printf("i: %v i: %T\n", i, i)
}

// i: 100 i: int
```

## Struct

### golang 结构体作为函数参数

go 结构体可以像普通变量一样，作为函数的参数，传递给函数，这里分为两种情况：

1. 直接传递结构体，这是一个副本（拷贝），在函数内部不会改变外面结构体内容 **值传递**
2. 传递结构体指针，这时在函数内部，能够改变外部结构体内容 **指针传递**

#### 直接传递结构体

```go
package main

import "fmt"

type Person struct {
    id   int
    name string
}

func showPerson(person Person) {
    person.id = 1
    person.name = "kite"
    fmt.Printf("person: %v\n", person)
}

func main() {
    person := Person{1, "tom"}
    fmt.Printf("person: %v\n", person)
    fmt.Println("-------------------")
    showPerson(person)
    fmt.Println("-------------------")
    fmt.Printf("person: %v\n", person)
}
```

```
person: {1 tom}
-------------------
person: {1 kite}
-------------------
person: {1 tom}
```

#### 传递结构体指针

```go
package main

import "fmt"

type Person struct {
    id   int
    name string
}

func showPerson(person *Person) {
    person.id = 1
    person.name = "kite"
    fmt.Printf("person: %v\n", *person)
}

func main() {
    person := Person{1, "tom"}
    fmt.Printf("person: %v\n", person)
    fmt.Println("-------------------")
    showPerson(&person)
    fmt.Println("-------------------")
    fmt.Printf("person: %v\n", person)
}
```

```
person: {1 tom}
-------------------
person: {1 kite}
-------------------
person: {1 kite}
```

### golang 嵌套结构体

go 语言没有面向对象编程思想，也没有继承关系，但是可以通过结构体嵌套来实现这种效果。

下面通过实例演示如何实现结构体嵌套，假如有一个人 Person 结构体，这个人还养了一个宠物 Dog 结构体

```go
import "fmt"

type Dog struct {
    name  string
    color string
    age   int
}

type person struct {
    dog  Dog
    name string
    age  int
}

func main() {
    var tom person
    tom.dog.name = "花花"
    tom.dog.color = "黑白花"
    tom.dog.age = 2

    tom.name = "tom"
    tom.age = 20

    fmt.Printf("tom: %v\n", tom)
}

// tom: {{花花 黑白花 2} tom 20}
```

### golang 方法

go 语言没有面向对象的特性，也没有类对象的概念。但是，可以使用结构体来模拟这些特性，我们都知道面向对象里面有类方法等概念。我们也可以声明一些方法，属于某个结构体。

#### go 语言方法的语法

Go 中的方法，是一种特殊的函数，定义于`struct`之上（与`struct`关联、绑定），被称为`struct`的接收者（`receiver`）。

通俗的讲，方法就是有接收者的函数。

```go
type mytype struct{}

func (recv mytype) my_method(para) retrun_type {}
func (recv *mytype) my_method(para) return_type {}
```

`mytype`：定义一个结构体

`recv`：接收该方法的结构体（receiver）

`my_method`：方法名称

`para`：参数列表

`return_type`：返回值类型

从语法格式可以看出，一个方法和一个函数非常相似，多了一个接收类型。

```go
package main

import "fmt"

type Person struct {
    name string
}

type Customer struct{
    name string
}

func (per Person) eat() {
    fmt.Println(per.name + " eating....")
}

func (per Person) sleep() {
    fmt.Println(per.name + " sleep....")
}

func (customer Customer) login(name string, pwd string) bool {
    fmt.Printf("customer.name: %v\n", customer.name)
    if name == "tom" && pwd == "123" {
        return true
    } else {
        return false
    }
}

func main() {
    var per Person
    per.name = "tom"
    per.eat()
    per.sleep()

    cus := Customer{"tom"}
    b := cus.login("tom", "123")
    fmt.Println(b)
}

// tom eating....
// tom sleep....
// customer.name: tom
// true
```

#### go 语言方法的注意事项

1. 方法的`receiver type`并非一定要是`struct`类型，`type`定义的类型别名、`slice`、`map`、`channel`、`func`类型等都可以。
2. `struct`结合它的方法就等价于面向对象中的类。只不过`struct`可以和它的方法分开，并非一定要属于同一个文件，**但必须属于同一个包**。
3. 方法有两种接收类型：`(T Type)`和`(T *Type)`，它们之间有区别。
4. 方法就是函数，所以 Go 中没有方法重载（overload）的说法，也就是说同一个类型中的所有方法名必须都唯一。
5. 如果`receiver`是一个指针类型，则会**自动解除引用**。
6. 方法和`type`是分开的，意味着**实例的行为（behavior）和数据存储（field）是分开的**，但是它们通过 receiver 建立起关联关系。

### golang 方法接收者类型

结构体实例，有值类型和指针类型，那么方法的接收者是结构体，那么也有值类型和指针类型。区别就是接收者是否复制结构体副本。值类型复制，指针类型不复制。

#### 值类型结构体和指针类型结构体

```go
package main

import "fmt"

type Person struct {
    name string
}

func main() {
    p1 := Person{name: "tom"}
    fmt.Printf("p1: %T\n", p1)
    p2 := &Person{name: "tom"}
    fmt.Printf("p2: %T\n", p2)
}

// 从运行结果，我们可以看出p1是值类型，p2是指针类型。
// p1: main.Person
// p2: *main.Person
```

```go
package main

import "fmt"

type Person struct {
    name string
}

func showPerson(per Person) {
    fmt.Printf("per: %p\n", &per)
    per.name = "kite"
    fmt.Printf("per: %v\n", per)
}

func showPerson2(per *Person) {
    fmt.Printf("per: %p\n", per)
    // per.name 自动解引用
    per.name = "kite"
    fmt.Printf("per: %v\n", per)
}

func main() {
    p1 := Person{name: "tom"}
    fmt.Printf("p1: %p\n", &p1)
    showPerson(p1)
    fmt.Printf("p1: %v\n", p1)
    fmt.Println("---------------------")
    p2 := &Person{name: "tom"}
    fmt.Printf("p2: %p\n", p2)
    showPerson2(p2)
    fmt.Printf("p2: %v\n", p2)
}

/*
从运行结果，我们看到p1是值传递，拷贝了副本，地址发生了改变，而p2是指针类型，地址没有改变。
p1: 0xc000050230
per: 0xc000050240
per: {kite}
p1: {tom}
---------------------
p2: 0xc000050270
per: 0xc000050270
per: &{kite}
p2: &{kite}
*/
```

#### 方法的值类型和指针类型接收者

值类型和指针类型接收者，本质上和函数传参道理相同。

```go
package main

import "fmt"

type Person struct {
    name string
}

func (per Person) showPerson() {
    fmt.Printf("per: %p\n", &per)
    per.name = "kite"
    fmt.Printf("per: %v\n", per)
}

func (per *Person) showPerson2() {
    fmt.Printf("per: %p\n", per)
    per.name = "kite"
    fmt.Printf("per: %v\n", per)
}

func main() {
    p1 := Person{name: "tom"}
    fmt.Printf("p1: %p\n", &p1)
    p1.showPerson()
    fmt.Printf("p1: %v\n", p1)
    fmt.Println("---------------------")
    p2 := &Person{name: "tom"}
    fmt.Printf("p2: %p\n", p2)
    p2.showPerson2()
    fmt.Printf("p2: %v\n", p2)
}

/*
p1: 0xc000050230
per: 0xc000050240
per: {kite}
p1: {tom}
---------------------
p2: 0xc000050270
per: 0xc000050270
per: &{kite}
p2: &{kite}
*/
```

### golang 接口

go 语言的接口，是一种新的类型定义，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。

#### 接口的语法格式

```go
/* 定义接口 */
type interface_name interface {
    method_name1 [return_type]
    method_name2 [return_type]
    method_name3 [return_type]
    ...
    method_namen [return_type]
}

/* 定义结构体 */
type struct_name struct {
    /* variables */
}

/* 实现接口方法 */
func (struct_name_variable struct_name) method_name() [return_type] {
    /* 方法实现 */
}
...
/* 实现接口方法 */
func (struct_name_variable struct_name) method_name() [return_type] {
    /* 方法实现 */
}
```

#### 接口实例

下面我定义一个 USB 接口，有读 read 和写 write 两个方法，再定义一个电脑 Computer 和一个手机 Mobile 来实现这个接口。

**USB 接口**

```go
type USB interface {
    read()
    write()
}
```

**Computer 结构体**

```haskell
type Computer struct {
}
```

**Mobile 结构体**

```go
type Mobile struct {
}
```

**Computer 实现 USB 接口方法**

```go
func (c Computer) read() {
    fmt.Println("computer read...")
}

func (c Computer) write() {
    fmt.Println("computer write...")
}
```

**Mobile 实现 USB 接口方法**

```go
func (c Mobile) read() {
    fmt.Println("mobile read...")
}

func (c Mobile) write() {
    fmt.Println("mobile write...")
}
```

```go
package main

import "fmt"

type USB interface {
    read()
    write()
}

type Computer struct {
}

type Mobile struct {
}

func (c Computer) read() {
    fmt.Println("computer read...")
}

func (c Computer) write() {
    fmt.Println("computer write...")
}

func (c Mobile) read() {
    fmt.Println("mobile read...")
}

func (c Mobile) write() {
    fmt.Println("mobile write...")
}

func main() {
    c := Computer{}
    m := Mobile{}

    c.read()
    c.write()
    m.read()
    m.write()

    /*
    接口d
    var usb USB
    usb = Computer{}
    usb.read()
    usb.write()
    */
}

/*
computer read...
computer write...
mobile read...
mobile write...
*/
```

#### 实现接口必须实现接口中的所有方法

下面我们定义一个 OpenClose 接口，里面有两个方法 open 和 close，定义个 Door 结构体，实现其中一个方法。

```go
package main

import "fmt"

type OpenClose interface {
    open()
    close()
}

type Door struct {
}

func (d Door) open() {
    fmt.Println("open door...")
}

func main() {
    var oc OpenClose
    oc = Door{} // 这里编译错误，提示只实现了一个接口
}
```

### golang 接口接收者类型

#### 值类型接收者

这个话题，本质上和方法的值类型接收者和指针类型接收者的思考方法是一样的，值接收者是一个拷贝，是一个副本，而指针接收者，传递的是指针。

```go
package main

import "fmt"

type Pet interface {
    eat()
}

type Dog struct {
    name string
}

func (dog Dog) eat() {
    fmt.Printf("dog: %p\n", &dog)
    fmt.Println("dog eat...")
    dog.name = "黑黑"
}

func main() {
    dog := Dog{name: "花花"}
    fmt.Printf("dog: %p\n", &dog)
    dog.eat()
    fmt.Printf("dog: %v\n", dog)
}

/*
dog: 0xc000050230
dog: 0xc000050240
dog eat...
dog: {花花}
*/
```

从运行结果，我们看出 dog 的地址变了，说明是复制了一份，dog 的 name 没有变说明外面的 dog 变量没有被改变。

#### 指针类型接收者

将 Pet 接口改为指针接收者

```go
package main

import "fmt"

type Pet interface {
    eat()
}

type Dog struct {
    name string
}

func (dog *Dog) eat() {
    fmt.Printf("dog: %p\n", dog)
    fmt.Println("dog eat...")
    dog.name = "黑黑"
}

func main() {
    dog := &Dog{name: "花花"}
    fmt.Printf("dog: %p\n", dog)
    dog.eat()
    fmt.Printf("dog: %v\n", dog)
}

/*
dog: 0xc000050230
dog: 0xc000050230
dog eat...
dog: &{黑黑}
*/
```

## Channel

### channel 的特点

- channel 本质就是一个数据结构-队列
- channel 是**引用类型**
- channel 必须初始化才能写入数据，即 make 后才能使用，且**只能存放指定的数据类型，除非使用空接口**
  - intChan := make(chan int, capacity)

### channel 死锁和阻塞 tips

- channel 的数据放满后，就不能放入了，如果继续放入，会导致死锁
- 如果从 channel 取出数据，可以继续放入
- 在没有使用协程的情况下，如果我们的管道数据已经全部去除，再取就会报告 **deadlock** (不要用 testing 框架测试，会变成**runtime 超时**)
- 在使用协程的情况下，如果我们的管道数据已经全部去除，再取就会**阻塞**

### channel 的关闭

使用内置函数 close 可以关闭 channel，当 channel 关闭后，就不能再向 channel 写数据了，但是仍然可以从该 channel 读取数据

### channel 的遍历

channel 支持 for-range 的方式进行遍历，请注意两个细节

- 在遍历时，如果 channel 没有关闭，则会出现 deadlock 的错误
- 在遍历时，如果 channel 已经关闭，则会正常遍历数据，遍历完成后，就会退出遍历

### channel 的阻塞

- 在没有使用协程的情况下，如果 channel 的长度小于写入数据的长度，那么会发生死锁的现象
- 在使用协程的情况下，如果 channel 的长度小于写入数据的长度，那么会阻塞，直到数据被读取，才会被唤醒
- 如果关闭了协程，v,ok := <-chan，那么 v 为默认值，ok 为 false

### channel 使用细节和注意事项

#### 1.channel 可以声明为只读或者只写性质

```go
package main

import "fmt"

func main() {
	// 管道可以声明为只读或者只写

	// 1. 在默认情况下，管道是双向的
	var chan1 chan int // 可读可写
	chan1 <- 20
	fmt.Printf("<-chan1: %v\n", (<-chan1))

	// 2. 声明为只写
	var chan2 chan<- int = make(chan int, 2)
	chan2 <- 20
	// 打印chan2的内容，其为一个地址，指向一个队列空间
	fmt.Printf("chan2: %v\n", chan2)

	// 3. 声明为只读
	var chan3 <-chan int
	<-chan3

}

```

应用案例

```go
package main

import "fmt"

// 形参指定模式
// ch chan<- int, only write
func send(ch chan<- int, exitChan chan struct{}) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	var a struct{}
	exitChan <- a
}

// ch <-chan int, only read
func recv(ch <-chan int, exitChan chan struct{}) {
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Printf("v: %v\n", v)
	}
	var a struct{}
	exitChan <- a
}

func main() {
	var ch = make(chan int, 10)
	exitChan := make(chan struct{}, 2)
	go send(ch, exitChan)
	go recv(ch, exitChan)

	var total = 0
	for _ = range exitChan {
		total++
		if total == 2 {
			break
		}
	}
	fmt.Println("ending......")

}
```

#### 2.使用`select`可以解决阻塞问题以及实现超时控制

2.1 使用`select`可以解决从管道取数据的阻塞问题（不需要关闭通道，解决阻塞)

```go
package main

import "fmt"

func main() {
	// 使用select可以解决从管道取数据的阻塞问题

	intChan := make(chan int, 10)

	for i := 0; i < 10; i++ {
		intChan <- i

	}

	strChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		strChan <- "hello" + fmt.Sprintf("%d", i)
	}

	// 传统方法在遍历管道时，如果不关闭会阻塞而导致deadlock
	// 问题 在实际开发中，我们可能不好确定什么时候关闭该管道
	// 可以使用select方式解决
	label:
	for {
		select {
		// 如果intChan一直没有关闭，不会一直阻塞而导致deadlock
		// 会自动到下一个case匹配
		case v := <-intChan:
			fmt.Printf("read intdata: %v\n", v)
		case v := <-strChan:
			fmt.Printf("read strdata: %v\n", v)
		default:
			fmt.Println("no data")
			break label
		}
	}
}
```

2.2 使用`select`实现超时控制

```go
select {
case ret := <-retCh:
  t.Logf("result %s", ret)
case <-time.After(time.Second * 1)
  t.Error("time out")
}
```

#### 3.goroutine 中使用`recover`，解决协程中出现`panic`，导致程序崩溃问题

```go
// 在goroutine中使用recover来捕获panic，进行处理，这样即使这个协程发生了问题，不会影响其他协程继续执行
defer func(){
    if err := recover(); err != nil {
        fmt.Printf("err: %v\n", err)
    }
}()
```

### 使用 WaitGroup 进行协程的同步

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var lock sync.Mutex
	var wg sync.WaitGroup
	count := 0

	for i := 0; i < 5000; i++ {
    // add one goroutine task
		wg.Add(1)
		go func() {
			lock.Lock()
			count++
			lock.Unlock()
      // end a goroutine task
			wg.Done()
		}()
	}

    // sync all tasks
	wg.Wait()
	fmt.Println(count)
}

```

## flag

可以用 flag 包对命令行参数进行解析

```go
package main

import (
	"flag"
	"fmt"
)

// os.Args[0] is the path to the executable.
// os.Args[1] is the first argument to the program.
// os.Args[2] is the second argument to the program.
// ...
// os.Args[n] is the nth argument to the program.

// go build -o test flag/main.go
// ./test -u root -p root -h localhost -port 3306
func main() {
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user, "u", "", "用户名 默认为空")
	flag.StringVar(&pwd, "p", "", "密码 默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名 默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口 默认为3306")

	flag.Parse()

	fmt.Printf("user=%v pwd=%v host=%v port=%v\n", user, pwd, host, port)
}
```
