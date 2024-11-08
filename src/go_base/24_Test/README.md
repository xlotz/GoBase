# 测试

## 介绍
    go test
    示例测试
    单元测试
    基准测试
    模糊测试

## 编写规范
    测试包： 单独包，test
    测试文件： _test.go 结尾
    测试函数：示例测试 ExampleXXX, 单元测试 Testxxx, 
            基准测试 Benchmarkxxx, 模糊测试 Fuzzxxx

## 执行测试
    go test xxx.go  // 单独测试文件
    go test -run xxx  // 单独指定某个测试文件的某个测试用例
    go test ./      //执行目录下所有测试用例
    go test ./ ../  
    go test -bench .  //运行所有类型的测试
    go test -bench . -run ^$  // 只运行所有的基准测试
    ... -v // 输出更详细的结果
    
    常用参数

    -o file	指定编译后的二进制文件名称
    -c	只编译测试文件，但不运行
    -json	以json格式输出测试日志
    -exec xprog	使用xprog运行测试，等价于go run
    -bench regexp	选中regexp匹配的基准测试
    -fuzz regexp	选中regexp匹配的模糊测试
    -fuzztime t	模糊测试自动结束的时间，t为时间间隔，当单位为x时，表示次数，例如200x
    -fuzzminimizetime t	模式测试运行的最小时间，规则同上
    -count n	运行测试n次，默认1次
    -cover	开启测试覆盖率分析
    -covermode set,count,atomic	设置覆盖率分析的模式
    -cpu	为测试执行GOMAXPROCS
    -failfast	第一次测试失败后，不会开始新的测试
    -list regexp	列出regexp匹配的测试用例
    -parallel n	允许调用了t.Parallel的测试用例并行运行，n值为并行的最大数量
    -run regexp	只运行regexp匹配的测试用例
    -skip regexp	跳过regexp匹配的测试用例
    -timeout d	如果单次测试执行时间超过了时间间隔d，就会panic。d为时间间隔，例1s,1ms,1ns等
    -shuffle off,on,N	打乱测试的执行顺序，N为随机种子，默认种子为系统时间
    -v	输出更详细的测试日志
    -benchmem	统计基准测试的内存分配
    -blockprofile block.out	统计测试中协程阻塞情况并写入文件
    -blockprofilerate n	控制协程阻塞统计频率，通过命令go doc runtime.SetBlockProfileRate查看更多细节
    -coverprofile cover.out	统计覆盖率测试的情况并写入文件
    -cpuprofile cpu.out	统计cpu情况并写入文件
    -memprofile mem.out	统计内存分配情况并写入文件
    -memprofilerate n	控制内存分配统计的频率，通过命令go doc runtime.MemProfileRate查看更多细节
    -mutexprofile mutex.out	统计锁竞争情况并写入文件
    -mutexprofilefraction n	设置统计n个协程竞争一个互斥锁的情况
    -trace trace.out	将执行追踪情况写入文件
    -outputdir directory	指定上述的统计文件的输出目录，默认为go test的运行目录


### 示例测试
    展示某一功能的使用方法，起到文档作用。

### 单元测试
    对软件中的最小可测试单元进行测试。
    t.Helper()

### 基准测试
    性能测试，测试程序的内存、CPU使用情况，执行耗时等性能指标。
    benchstat 性能测试分析工具

### 模糊测试