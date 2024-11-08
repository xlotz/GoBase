# 字符串

## 介绍
    不可变的只读的字节数组，分配一片连续的内存空间

## 字面量

### 普通字符串
    由"" 表示，支持转义，不支持多行

### 原生字符串
    由`` 表示，不支持转义（包括换行和缩进），支持多行


## 访问
    字节数组，访问方式和数组切片一致
    str := "this is a string"
    str[0]
    str[0:4]
    // 不能修改，但可以覆盖
    str[2] = 'a'  // 报错
    str = "aaaa"

## 转换
    字符串转换为字节切片， 字节切片或字节数组 转换字符串
    字符串内容只读，不可变，无法修改，但字节切片可修改。

## 长度
    len(str)

## 拷贝
    使用内存函数 copy
    copy(dest, src)
    dest = strings.Clone(src)

## 拼接
    使用 + , str = str + "xxx"
    使用切片, append(bytes, "aaaa")
    （高性能） strings.Builder{}


## 遍历
    只读的字节切片，针对utf8，需要使用rune类型，可使用range

    


