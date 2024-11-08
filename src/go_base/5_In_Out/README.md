# 输入输出

## 文件描述符

    var (
        Stdin = NewFile(uintptr(syscall.Stdin), "/dev/sdtin.out")
        Stdout = NewFile(uintptr(syscall.Stdout), "/dev/sdtout.out")
        Stderr = NewFile(uintptr(syscall.Stderr), "/dev/sdterr.out")
    )
    os.Stdin 标准输入
    os.Stdout 标准输出
    os.Stderr 标准错误


##  输出
### stdout
    标准输出是一个文件，所以可以直接写入
    os.Stdout.WriteString("Hello")


### print
    内置函数 print\println, 输出到标准错误中
    print("hello\n")
    println("hello")
    
### fmt
    默认输出到标准输出中
    fmt.Println("hello")

### bufio
    可缓冲的输出，先写入到内存，累积到一定阈值在输出到指定的writer中，默认缓冲区4KB，网络IO建议。
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()
    writer.WriterString("hello")
    fmt.Println(writer, "hello")


### 格式化
    由fmt.Printf 函数提供

    | 格式化 | 描述            | 接收类型
    %%       输出%              任意
    %s       string/[]byte     string, []byte
    %q       输出包含""         string, []byte
    %d       输出十进制          整型
    %f       浮点               浮点
    %e、%E   科学计数，复数       浮点
    %g      输出%f/%e,去掉多余0   浮点
    %b      整型的二进制          数字
    %#b     完整的二进制          数字
    %o      整型的八进制          整型
    %#o     完整的八进制          整型
    %x      整型的小写十六进制     数字
    %#x     完整的小写十六进制     数字
    %X      整型的大写十六进制     数字
    %#X     完整的大写十六进制     数字
    %v      原本的形式，数据结构输出    任意
    %+v     结构体时加上字段名     任意
    %t      布尔值              布尔
    %T      输出对应的类型值       任意
    %c      Unicode对应的字符    int32
    %U      字符对应的Unicode码   rune,byte
    %p      指针指向的地址         指针
### 

## 输入

### read
    如读取文件一样读取输入的内容
    var buf [1024]byte
    n, _ := os.Stdin.Read(buf[:])
    os.Stdout.Write(buf[:n])

### fmt
    Scan扫描从os.Stdin读入的文本，空格分隔，换行也会当做空格
    Scanln 和Scan类似，但换行结束
    Scanf 格式化字符串扫描
    
    // 读取2个数字
    var a, b int
    fmt.Scanln(&a, &b)
    fmt.Printf("%d + %d = %d\n", a, b, a+b)

    // 读取固定长度的数组
    n := 10
    s := make([]int, n)
    for i := range n {
        fmt.Scan(&s[i])
    }
    fmt.Println(s)
    

### bufio
    适用于大量输入需要读取，bufio.Read
    reader := bufio.NewReader(os.Stdin)
    var a, b int
    fmt.Fscanln(reader, &a, &b)
    fmt.Printf("%d + %d = %d\n", a, b, a+b)

### scanner
    和bufio.Reader类似，不过是按行读取
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan(){
        line := scanner.Text()
        if line == "exit" {
            break
        }
        fmt.Println("scan", line)
    }

###

