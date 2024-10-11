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
    
### 交集
### 空集
### 空接口
### 底层类型

### 类型集注意

## 使用

### 队列

### 堆

## 小结
