# 反射

## 反射的基本介绍

1. 反射可以在运行时动态获取变量的各种信息，比如变量的类型(type)，类别(kind)

- 如果是基本数据类型，type==kind
- 如果是结构体，type!=kind

2. 如果是结构体变量，还可以获取到结构体本身的信息（包括结构体的字段、方法）
3. 通过反射，可以修改变量的值，可以调用关联的方法。
4. 使用反射，需要 import("reflect")

通过 reflect.typeOf(i interface{})返回 reflect.Type 类型。使用 reflect.Type 接口关联的方法，可以对 i 进行修改
通过 reflect.valueOf(i inerface{})返回 reflect.Value 类型。使用 reflect.Value 接口关联的方法，甚至调用结构体 i 的方法。

- reflect.Value 也可以调用方法，返回 reflect.Type 类型

## 反射重要的函数和概念

1. reflect.TypeOf(变量名)，获取变量的类型，返回 reflect.Type 类型。
2. reflect.ValueOf(变量名)，获取变量的值，返回 reflect.Value 类型。

- reflect.Value 是一个结构体类型。通过 reflect.Value，可以获取到关于该变量的很多信息

3. 变量、interface{}和 reflect.Value 是可以相互转换的，这点在实际开发中，会经常使用到。

```go
var student Stu

// 通过参数传递，将Stu转换为interface{}
func test(b interface{}){

    // 1. 将interface{}转换为reflect.Value
    rVal := reflect.ValueOf(b)

    // 2. 将reflect.Value转换为interface{}
    iVal := rVal.Interface()

    // 3. 将interface{}转换为原来的变量类型，使用类型断言
    v := iVal.(Stu)
}

```

> tips: 利用 switch x.(type)进行类型断言
>
> ```go
> switch x.(type){
> case Student:
> case Person:
> }
> ```

## 反射常见的应用场景

1. 不知道接口调用哪个函数，根据传入参数在运行时确定调用的具体接口，这种需要对函数活方法反射。
2. 对结构体序列化时，如果结构体有指定 Tag，也会使用到反射生成对应的字符串

## 反射注意事项和细节说明

1. reflect.Value.Kind，获取变量的类型，返回的是一个常量
2. Type 是类型，Kind 是类别，Type 和 Kind 可能是相同的，也可能是不同的。

- 比如：var num int = 10 num 的 Type 是 int，Kind 也是 int
- 比如：type MyInt int; var num MyInt = 10 num 的 Type 是 MyInt，Kind 是 int
- 比如：var stu Student stu 的 Type 是 module.Student，Kind 是 struct

3. 通过反射可以让变量在 interface{}和 reflect.Value 之间相互转换
4. 使用反射的方法来获取变量的值（并返回对应的数据），要求数据类型匹配，比如 x 是 int，那么就应该使用 reflect.Value(x).Int()，而不能使用其他的，否则会 panic
5. 通过反射来修改变量，主要当使用SetXxx方法来设置需要通过对应的指针类型来完成，这样此才能传入变量的值，同时需要使用reflect.Value.Elem()来获取指针的值



