# 迭代器

## 介绍
    作用于以下的几个数据结构
    数组、切片、字符串、map、chan、整型值

## 推送式迭代器
    由迭代器控制迭代的逻辑，用户被动获取元素


## 拉取式迭代器
    由用户控制迭代器逻辑，祖东获取序列元素，使用特定的函数
    next()、stop() 控制迭代的开始或结束。


## 错误处理
    迭代时发送错误，可将其传递给yield让for range 返回，让调用者处理。
    
## 标准库
    slices
        slices.All 将切片转换成切片迭代器
        slices.Values 将切片转换成切片迭代器，不带索引
        slices.Chunk 返回一个迭代器，将以n个元素为切片推送给调用者
        slices.Collect 将切片迭代器收集成一个切片
        
    maps
        maps.Keys 返回迭代map所有键的迭代器，配合slices.Collect 直接收集成一个切片
        maps.Values 返回迭代map所有值的迭代器，配合slices.Collect 直接收集成一个切片
        maps.All 将一个map转换成功一个map迭代器
        maps.Collect 将一个map迭代器收集成一个map
