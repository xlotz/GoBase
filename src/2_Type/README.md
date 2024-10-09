# 数据类型

## 布尔
    bool // 0不代表真假，true/false

## 整型

    序号	类型和描述
    uint8	无符号 8 位整型
    uint16	无符号 16 位整型
    uint32	无符号 32 位整型
    uint64	无符号 64 位整型
    int8	有符号 8 位整型
    int16	有符号 16 位整型
    int32	有符号 32 位整型
    int64	有符号 64 位整型
    uint	无符号整型 至少32位
    int	整型 至少32位
    uintptr	等价于无符号64位整型，但是专用于存放指针运算，用于存放死的指针地址。

## 浮点型

    分单精度和双精度
    float32  32位浮点数
    float64  64位浮点数

## 复数类型

    complex128  64位实数和虚数
    complex64   32位实数和虚数

## 字符类型
    go字符串兼容UTF-8
    byte // == uinnt8 可表达ANSCII字符
    rune // == int32 可表达Unicode字符
    string // 字符序列，可转换[]byte类型（字节切片）

## 派生类型
    类型	例子
    数组	[5]int，长度为5的整型数组
    切片	[]float64，64位浮点数切片
    映射表	map[string]int，键为字符串类型，值为整型的映射表
    结构体	type Gopher struct{}，Gopher结构体
    指针	*int，一个整型指针。
    函数	type f func()，一个没有参数，没有返回值的函数类型
    接口	type Gopher interface{}，Gopher接口
    通道	chan int，整型通道
## 零值
    一个类型的空值或默认值
    数字类型    0
    布尔类型    false
    字符串类型   ""
    数组      固定长度的对应类型的零值集合
    结构体     内部字段都是零值的结构体
    切片、映射、函数、接口、通道、指针   nil

## nil
    一个变量，nil == nil 编译不通过