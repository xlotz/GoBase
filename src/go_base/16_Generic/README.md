# 泛型

## 介绍
    解决执行逻辑与类型无关的问题
    
    func Sum[T int|float64](a, b T) T {
        return a + b
    }
    其中：
        T： 类型形参
        int | float64: 类型约束
        Sum[int](1, 2): 类型实参

    用法：
        Sum[int](1, 2)
        Sum(1, 2)

## 泛型结构
    泛型切片：
        定义：
        type GSlice[T int|int32|int64] []T
        使用：
        //不能省略类型
        GSlice[int]{1,2,3}
    泛型Hash:
        定义：
        type GMap[k comp, v int|string|byte] map[k]v
        // comp 是接口
        使用：
        g1 := GMap[int,string]{1: "hello"}
        g2 := make(GMap[string,byte], 0)
    泛型结构体:
        定义：
        type GStruct[T int|string] struct{
            Name string
            Id T
        }
        使用：
        GStruct[int]{
            Name: "jack",
            Id: 1024,
        }  
        GStruct[string]{
            Name: "jack",
            Id: "1024",
        }
    泛型切片形参：
        type Comp[T int|string, S []T] struct{
            Name string
            Id T
            Stuff S
        }
        或
        type Comp[T int|string, S []int|string] struct{
            Name string
            Id T
            Stuff S
        }

        使用：
        Comp[int, []int]{
            Name: "aaa",
            Id: 1,
            Stuff: []int{1},
        }

### 泛型结构注意
    泛型 不能作为一个类型的基本类型
        type G[T int|int32|int64] T
        type G[T int|int32|int64] int
    泛型无法使用类型断言
    匿名结构不支持泛型
    匿名函数不支持自定义泛型
    不支持泛型方法


## 类型集
    1.18后接口定义变成了类型集 type set, 含有类型集的接口又称为通用接口。
    类型集主要用于类型约束，不能作为类型声明。

### 并集
    接口类型SignedInt类型集
    type SingedInt interface {
        int8|int16|int|int32|int64
    }
    type UnSignedInt interface {
        uint8|uint16|uint32|uint64
    }
    type Integer interface {
        SignedInt|UnSignedInt
    }
    
### 交集
    非空接口的类型集
    type SignedInt interface {
        int8 | int16 | int | int32 | int64
    }
    
    type Integer interface {
        int8 | int16 | int | int32 | int64 | uint8 | uint16 | uint | uint32 | uint64
    }
    
    type Number interface {
        SignedInt
        Integer
    }

### 空集
    没有交集
    type SignedInt interface {
	    int8 | int16 | int | int32 | int64
    }
    
    type UnsignedInt interface {
        uint8 | uint16 | uint | uint32 | uint64
    }
    
    type Integer interface {
        SignedInt
        UnsignedInt
    }
    
### 空接口
    
    空接口时使用类型集的集合
    func Do[T interface{}](n T) T {
        return n
    }
    
    func main() {
        Do[struct{}](struct{}{})
        Do[any]("abc")
    }

### 底层类型

    使用type声明一个新的类型，底层类型包含在类型集内。
    

### 类型集注意
    带有方法集的接口无法并发类型集
    类型集无法当做类型实参使用
    类型集中的交接问题
    类型集不能直接或间接的并入自身
    comparable 接口无法并入类型集


## 使用
    

### 队列

    声明队列类型
    type Queue[T any] []T

### 堆
    特殊数据结构，可以在O(1)的时间内判断最大或最小值，可以排序的类型。
    type Comparator[T any] func(a,b T)int

## 小结
