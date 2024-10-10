# 指针

## 创建
    取地址符 &， 对一个变量取地址，会返回对应类型的指针
    解引用符 *

    num := 2
    p := &num
    fmt.Println(p)

    r := *p
    fmt.Println(r)

    指针需要声明并初始化(分配内存)，才能正常使用。
    var n *int
    n = new(int)
    或
    n := new(int)

    指针禁止运算（无法偏移）
    
## * new 和 make

    func new(Type) *Type
    说明：
        返回值是类型指针
        接收参数是类型
        用于给指针分配内存空间

    func make(t Type, size ...InterType) Type
    说明:
        返回值是值，不是指针
        参数：类型、不定长参数根据传入的类型不同而不同
        用于给切片、映射表、通道 分配内存

    