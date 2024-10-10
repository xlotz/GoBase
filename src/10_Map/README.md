# 映射表

## 介绍
    映射表实现通常有两种，hash表（无序），搜索树（有序）
    map 基于hash表，无序
    引用型，零值或未初始化的map可以访问，但无法存放元素，必须分配内存（即分配容量）
    对不存在的键，返回对应类型的零值

## 初始化
    // 第一种
    map[keyType]valueType{}
    
    mp := map[int]string{
        0: "a",
        1: "a",
        2: "a",
    }    
    mp := map[string]int{
        "a": 0,
        "b": 1,
        "c": 2,
    }

    // 第二种 make, 包含 类型和初始容量
    mp := make(map[string]int,8)
    mp := make(map[string][]int, 10)

## 访问
    mp := map[string]int{
        "a":0,
        "b":1,
        "c":2,
    }
    fmt.Println(mp["a"])
    fmt.Println(len(mp))


## 存值
    类似数组存值，对已存在的值直接覆盖
    mp := make(map[string]int, 10)
    mp["a"] = 1
    mp["a"] = 2

## 删除
    func delete(m map[Type]Type1, key Type)
    delete(map, "a")
    如果值为NaN，无法删除该键值对

## 遍历
    使用for range 遍历，如果值为NaN 可以遍历
    
    for key, val := range mp {
        fmt.Println(key, val)
    }
## 清空
    使用clear
    clear(mp)
    

## set
    无序的，不包含重复元素的集合，go使用map 替代set
    set := make(map[int]struct{}, 10)
    for i:=0; i<10;i++{
        set[rand.Intn(100)] = struct{}{}
    }


## 注意
    map 不是并发安全的数据结构，不涉及高并发的场景，内部有读写检测机制，冲突时会触发fatal error.
    此时需要使用sync.Map替代。
    