# 方法

## 介绍
    方法和函数的区别： 
        方法有接收者，函数没有
        只有自定义类型拥有方法

    先声明，再初始化，再调用

    type IntSlice []int

    func (i IntSlice) Get(index int) int {
        return i[index]
    }
    func (i IntSlice) Set(index, val int) {
        i[index] = val
    }
    func (i IntSlice) Len() int {
        return len(i)
    }
    先声明一个类型IntSlice,底层类型[]int, 在声明三个方法Get, Set, Len
    其中i 是接收者类似this\self，IntSlice是接收者的类型。


## 接收者
    值接收者
    
    指针接收者