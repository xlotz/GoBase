# 反射

## 介绍
    反射是在运行时检查语言自身结构的机制，go的反射API是由标准库reflect提供。

## 接口
    没有方法集的接口
    由方法集的接口

    静态：体现在对外表现的抽象的接口类型是不变的。
    动态：接口底层存储的具体实现类型可变化的。

    type iface struct {
        tab *itab // 包含数据类型，接口类型、方法集等
        data unsafe.Pointer // 指向值的指针
    }
    
    type eface struct {
        _type *_type // 类型
        data unsafe.Pointer //指向值的指针
    }

    type nonEmptyInterface struct {
        itab *struct {
            ityp *rtype //静态接口类型
            typ *rtype //动态具体类型
            hash uint32 //类型哈希
            _ [4]type
            fun [100000]unsafe.Pointer //方法集
        }
        word unsafe.Pointer //指向值的指针
    }

    type emptyInterface struct {
        typ *rtype //动态具体类型
        word unsafe.Pointer //指向指针的值
    }

## 核心
    1.反射将interface{}类型变量转换成发射对象
    2.发射将发射对象还原成interface{}类型变量
    3.要修改反射对象，其值必须可设置
    访问类型信息时，用到reflect.TypeOf
    修改反射值时，用到reflect.ValueOf

## 类型
    reflect.TypeOf()将变量转变成 reflect.Type 

### Kind
    表示基础类型（无符号整型uint）,reflect包使用kind枚举所有的基础类型。

### Elem
    判断类型为any的数据结构所存储的元素类型，可接收的底层参数类型必须是指针、切片、数组、通道、映射表，否则会panic。

### Size
    获取对应类型占的字节大小。（unsafe.Sizeof()同样的效果）

### Comparable
    判断一个类型是否可以比较。

### Implements
    判断一个类型是否实现某一接口。

### ConvertibleTo
    可以判断一个类型是否被转换为另一个指定的类型。

## 值
    reflect.Value代表着反射接口的值，使用reflect.ValueOf()函数将变量转换成reflect.Value

### Type
    获取一个反射值的类型

### Elem
    获取一个反射值的元素反射值

### 指针
    获取一个反射值的指针方式：
        // 返回一个表示v地址的指针反射值
        func (v Value) Addr() Value
        
        // 返回一个指向v的原始值的uinptr 等价于 uintptr(Value.Addr().UnsafePointer())
        func (v Value) UnsafeAddr() uintptr
        
        // 返回一个指向v的原始值的uintptr
        // 仅当v的Kind为 Chan, Func, Map, Pointer, Slice, UnsafePointer时，否则会panic
        func (v Value) Pointer() uintptr
        
        // 返回一个指向v的原始值的unsafe.Pointer
        // 仅当v的Kind为 Chan, Func, Map, Pointer, Slice, UnsafePointer时，否则会panic
        func (v Value) UnsafePointer() unsafe.Pointer

### 设置值
    通过反射修改反射值，其值必须是可取址, 通过指针来修改其元素值，

### 获取值
    通过Interface() 获取反射值原有的值


## 函数
    通过反射获取函数的一切信息，也可以反射调用函数

### 信息
    通过反射类型来获取函数的一切信息

### 调用
    通过反射值来调用函数

## 结构体
    type Person struct {
        Name    string `json:"name"`
        Age     int    `json:"age"`
        Address string `json:"address"`
        money   int
    }
    
    func (p Person) Talk(msg string) string {
        return msg
    }

### 访问字段
    reflect.StructField
    通过索引进行访问，通过名称访问。

### 修改字段
    修改结构体字段值，需传入一个结构体指针
    
### 访问Tag
    获取structField, 可以直接访问Tag

### 访问方法
    访问方法与访问字段类型类似

### 调用方法
    调用方法与调用函数的过程类似，不需要手动传入接入者

## 创建
    通过反射构建新的值，reflect包根据一些特殊的类型提供不同的函数。

### 基本类型
    // 返回指向反射值的指针反射值
    func New(typ Type) Value

### 结构体
    reflect.New 函数

### 切片
    reflect.MakeSlice 创建切片

### map
    reflect.MakeMapWithSize 创建map

### 管道
    reflect.MakeChan 创建管道

### 函数
    reflect.MakeFunc 创建函数

## 完全相等
    reflect.DeepEqual 反射包下提供的一个用于判断两个变量是否完全相等的函数。
    判断方式：
    * 数组：数组中的每一个元素都完全相等
    * 切片：都为nil时，判为完全相等，或者都不为空时，长度范围内的元素完全相等
    * 结构体：所有字段都完全相等
    * 映射表：都为nil时，为完全相等，都不为nil时，每一个键所映射的值都完全相等
    * 指针：指向同一个元素或指向的元素完全相等
    * 接口：接口的具体类型完全相等时
    * 函数：只有两者都为nil时才是完全相等，否则就不是完全相等
