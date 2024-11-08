# 常量

## 定义
    值无法再运行时改变，一旦赋值无法修改
    来源：字面量、其他常量标识符、常量表达式、结果是常量的类型转换，iota
    常量只能是基础类型，不能是结构体、接口、切片、数组、函数的返回值等。

## 初始化
    const, 声明时必须初始化一个值，类型可省略。
    const name string = "jack"
    const num = 1
    const sum = 1 +2 +3

    const (
        Count = 1
        Name = "jack"
    )

    同一个常量分组中，后面的常量可不用赋值，使用前一个的值
    const (
        A = 1
        B //1
        C //1
    )

## iota
    iota: 内置的常量标识符，表示一个常量声明中的无类型整数序数，
            一般括号中使用，递增。
    const (
        Num = iota //0
        Num1 //1
        Num2 //2
        Num3 //3
    )
    const (
        Num = iota * 2 //0
        Num1 //2
        Num2 //4
    )
    
## 枚举
    使用 自定义类型 + const + iota 实现
    
