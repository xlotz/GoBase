# 切片

## 介绍
    存放不定长的数据，后续使用过程中可以修改和删除，自动扩张
    引用类型

## 初始化

    var nums []int //值
    nums := []int{1,2,3} // 值
    nums := make([]int, 0, 0) // 值（包含类型，长度，容量）
    nums := new([]int) // 指针

## 使用
    nums := make([]int, 0,0)
    nums := append(nums, 1,2,3,4,5,6,7)
    fmt.Println(len(nums), cap(nums))
    
    // cap 超过256， newcap = old + (old + 3*256)/4

## 插入元素
    nums := []int{1,2,3,4,5,6,7,8,9,10}
    // 从头插入
    nums = append([]int{-1,0}, nums...)
    // 从中间插入
    nums = append(nums[:i+1], append([]int{999,999}, nums[i+1:]...)...)
    // 从尾插入
    nums = append(nums, 99, 100)

## 删除元素
    nums := []int{1,2,3,4,5,6,7,8,9,10}
    // 从头删
    nums = nums[n:]
    // 从尾删
    nums = nums[:len(nums)-n]
    // 从中间删
    nums = append(nums[:i], nums[i+n:]...)
    // 删除所有
    nums = nums[:0]

## 拷贝
    使用copy
    dest := make([]int, 0)
    src := []int{1,2,3,4,5,6}
    copy(dest, src)


## 遍历
    使用for
    slice := []int{1,2,3,4,5,6,7}
    for i:=0;i<len(slice);i++{
        fmt.Println(slice[i])
    }

    for i, v := range slice {
        fmt.Println(i, v)
    }
    

## 多维切片
    slices := make([][]int, 5)

## 拓展
    slice[low:high:max]
    满足: low <= high <= max <= cap

## clear
    将切片所有值置0
    s := []int{1,2,3,4}
    clear(s) // [0 0 0 0]

    清空
    s = s[:0:0]

    