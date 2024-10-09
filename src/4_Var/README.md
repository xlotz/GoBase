# 变量

## 概述
    用于保存一个值的存储位置，运行时可变化，每声明一个变量会分配一块内存存储对应类型的值。

## 声明
    只声明不赋值，则存储的是对应类型的零值
    var a int
    var a, b , c int
    var (
        name string
        age int
        addr string
    )


## 赋值
    # 第一种
    var name string
    name = "jack"
    # 第二种
    var name string="jack"
    # 第三种
    var name = "jack"
    # 第四种（短变量）
    name := "jack"  
    name, age := "jack", 1
    
### 赋值注意：
    1. 短变量初始化不能使用nil
    2. 短变量无法对一个已存在的变量赋值，不过可以在赋值一个旧变量的同时声明一个新变量
    var a int
    a:=1  // 错误

    a := 1
    a, b := 2, 2 // 正常

    3. 函数内的变量定义后必须被使用，函数外不受限制。
## 匿名

    使用下划线表示不使用某个变量
    file, _ := os.Open("readme.txt")

## 交换
    n1, n2 := 1, 2
    n1, n2 = n2, n1
    
    a, b, c := 0, 1, 1
    a, b, c = b, c, a+b // 先计算值在被赋值
    out: 1, 1, 1

## 比较
    类型必须相同, 支持所有的可比较的类型
    var a uint64
    var b int64
    fmt.Println(int64(a) == b)

    minV := min(1, 2, -1, 1.2)
    maxV := max(100, 22, -1, 1.12)

    可比较类型： 布尔、数字、字符串、指针、
               通道（判断是否相同）、
               数组（元素类型可比较的数组，切片不可以）
               结构体（元素类型可比较的结构体，仅支持是否相等）




