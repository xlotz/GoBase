package main

import "fmt"

func main() {

	// 初始化
	var n1 []int            //值
	n2 := []int{1, 2, 3}    //值
	n3 := make([]int, 0, 0) //值
	n4 := new([]int)        // 指针
	fmt.Println(n1, n2, n3, n4)

	nums := make([]int, 0, 0)
	nums = append(nums, 1, 2, 3, 4, 5)
	fmt.Println(len(nums), cap(nums))

	// 从头插
	nums = append([]int{-1, 0}, nums...)
	fmt.Println(nums)

	// 从中间下标i 插
	i := 2
	nums = append(nums[:i+1], append([]int{999, 999}, nums[i+1:]...)...)
	fmt.Println("从下标", i, "插入", nums)

	// 从尾插入
	nums = append(nums, 99, 100)
	fmt.Println(nums)

	fmt.Println("删除元素")
	// 删除元素
	n := 1
	nums = nums[n:]
	fmt.Println(nums)
	nums = nums[:len(nums)-n]
	fmt.Println(nums)
	nums = append(nums[:i], nums[i+n:]...)
	fmt.Println(nums)

	// 删除所有
	fmt.Println("删除所有")
	nums = nums[:0]
	fmt.Println(nums)

	// 拷贝
	fmt.Println("拷贝")
	//dest := make([]int, 0)
	dest := make([]int, 10)
	src := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(src, dest)
	fmt.Println(copy(dest, src))
	fmt.Println(src, dest)

	// 遍历
	fmt.Println("遍历")
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, v := range slice {
		fmt.Println(i, v)
	}

	// 多维
	fmt.Println("多维")
	var nums1 [5][5]int
	for _, num := range nums1 {
		fmt.Println(num)
	}
	fmt.Println()
	slices := make([][]int, 5)
	for _, slice := range slices {
		fmt.Println(slice)
	}

	// 修复
	for i := 0; i < len(slices); i++ {
		slices[i] = make([]int, 5)
	}
	for _, slice := range slices {
		fmt.Println(slice)
	}

	// 拓展
	fmt.Println("拓展")
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} // cap = 9
	s2 := s1[3:4]                          // cap = 6
	//添加元素, 容量6，没有扩容直接修改底层数组
	fmt.Println(s2, s1)
	s2 = append(s2, 1)
	fmt.Println(s2, s1)

	// 修改
	s3 := s1[3:4:4]
	fmt.Println(s3, s1)
	s3 = append(s3, 1)
	fmt.Println(s3, s1)

	// clear
	fmt.Println("clear")
	s := []int{1, 2, 3, 4}
	clear(s)
	fmt.Println(s)

	fmt.Println("清空")
	s = s[:0:0]
	fmt.Println(s)
}
