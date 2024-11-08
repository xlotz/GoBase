# 数组

## 介绍
    定长的数据结构，长度被指定后不能被改变，值类型，值传递（整个数组拷贝）

## 初始化
    数组在声明是长度只能是一个常量，分配固定大小的内存

    var a [5]int
    nums :=[5]int{1,2,3}
    nums := new([5]int) //值是指针

## 使用

    fmt.Println(nums[0])
    nums[0] = 1
    len(nums)  //数组元素的数量
    cap(nums)  //数组容量

## 切割

    nums := [5]int{1,2,3,4,5}
    nums[1:]
    nums[:5]
    nums[1:3]

    arr := [5]int{1, 2, 3, 4, 5}
	slice := slices.Clone(arr[:])  // 拷贝，切片修改不影响数组
	slice[0] = 0
	fmt.Printf("array: %v\n", arr)
	fmt.Printf("slice: %v\n", slice)

    
    
    