# 循环

## for
    仅有 for, 可以当做while来用。

    for init; expr; post {
        exec state
    }
    
    // while
    for expr {
        exec state
    }

    // 死循环
    for {
        exec state
    }

## for range
    遍历一些可迭代的数据结构，如数组，切片，字符串，映射，通道
    for index, value := range iter {
        exec state
    }

## break
    终止最内层的for循环，结合标签可达到终止外层循环效果


## continue
    跳过最内层循环的迭代，直接进入下一次跌打，结合标签可调到外层循环
