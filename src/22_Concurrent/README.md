# 并发

## 协程
    轻量级的线程，或用户态的线程，不受操作系统调度，由GO调度器调度。
    上下文切换开销小。
### 创建

    go hello()
    go func(){
    }()
    
    注意：
        具有返回值的内置函数不允许跟在go关键字后面
        协程是并发执行，执行顺序不确定
        为避免不确定协程的退出状态，需要控制并发，并发控制方法包括：
            管道 channel：适合协程间通信
            信号量 waitgroup: 适合动态控制一组指定数量的协程。
            上下文 context：适合子孙协程嵌套层级更深的情况
        锁控制：
            互斥锁 Mutex
            读写互斥锁 RWMutex
    
## 管道
    通过消息进行内存共享，一种协程间通信的解决方案，也可以用于并发控制。
    基本语法：
        var Name chan Type
    

### 创建
    只能使用内置函数make
    intCh := make(chan int)
    strCh := make(chan string, 1) // 带缓冲区
    close(intCh) // 关闭管道
    

### 读写
    ch <- 对一个管道写入数据
    <- ch 对一个管道读取数据
    管道里的数据流动方式和队列一样，先进先出（FIFO），对于管道的操作，
    同一个时刻，只有一个协程写入同时也只有一个协程能够读取。

    ints, ok := <-intCh

### 无缓冲
    缓冲区为0，不能临时存放任何数据。写入数据必须立即读取，否则会阻塞等待（读取也是同理）。
    

### 有缓冲
    写入数据时，会先将数据放入缓冲区里，只有当缓冲区容量满了才会阻塞的等待协程来读取管道中的数据。
    同样的，读取有缓冲管道时，会先从缓冲区中读取数据，
    直到缓冲区没数据了，才会阻塞的等待协程来向管道中写入数据

### 注意点
    使用不当时会导致管道阻塞：
    1. 读写无缓冲管道
    2. 读取空缓冲区管道
    3. 写入满缓冲区管道
    4. 管道为nil
    5. 写入已关闭的管道
    6. 关闭已关闭的管道

### 单向管道
    双向管道：可以在管道的两边进行读写操作。
    单向管道：只能读或只能写的管道。用于限制通道的行为，会在函数的形参和返回值中出现。
    
    <- chan int 只读管道
    chan <- string 只写管道
    
    双向管道可以转换成单向管道，反之不行。

### for range
    遍历读取缓冲管道中的数据，有且只有一个返回值

### select
    管道多路复用的控制结构。（某一时刻同时检测多个元素是否可用，被检测的可以是网络IO、文件IO）
    select 检测的只能是管道。
    语法和switch 类型
    select {
    case xxx:
        xxx
    default:
        xxx
    }

    永久阻塞：当select中什么都没有时，会永久阻塞
    

## WaitGroup

    sync.WaitGroup 是sync包下提供的结构体，等待执行，实现等待一组协程的效果。 
    适用于可动态调整协程数量的场景。
    提供三个方法：
    Add 用于指明要等待的协程的数量
        func (wg *WaitGroup) Add(delta int)
    Done 表示当前协程已经执行完毕
        func (wg *WaitGroup) Done()
    Wait 等待子协程介绍，否则阻塞
        func (wg *WaitGroup) Wait()
    内部实现：
        计数器+信号量，程序开始时调用Add 初始化计数，当一个协程执行完调用Done,
        计数 -1，只到为0，在此期间主协程调用Wait阻塞等待只到计数为0，开会唤醒。


## Context
    上下文，一种并发控制的解决方案，对比管道 和 WaitGroup ，可以更好的控制子孙协程以及更深层级的协程。
    提供：
    emptyCtx
    cancelCtx
    timerCtx
    valueCtx

### Context
    type Context interface {

        Deadline() (deadline time.Time, ok bool)
        
        Done() <-chan struct{}
        
        Err() error
        
        Value(key any) any
    }
    Deadline 返回 截止时间（上下文应该取消的时间）和 是否设置deadline,默认 false.
    Done 返回一个空结构体类型的只读管道，起到通知作用，不传递任何数据。
    Err 返回一个error, 表示上下文关闭的原因，没有关闭时返回nil.
    Value 返回对应的键值，如果key 不存在或不支持，返回nil


### emptyCtx
    空的上下文，可以通过context.BackGround、context.TODO创建。

    var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
    )
    
    func Background() Context {
    return background
    }
    
    func TODO() Context {
    return todo
    }

    只能返回指针，需要有不同的内存地址，没法被取消，没有deadline,不能取值。
    实现的方法都返回零值。

    type emptyCtx int
    
    func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
    return
    }
    
    func (*emptyCtx) Done() <-chan struct{} {
    return nil
    }
    
    func (*emptyCtx) Err() error {
    return nil
    }
    
    func (*emptyCtx) Value(key any) any {
    return nil
    }
    

### valueCtx
    内部只包含一对键值对，和一个内嵌的Context类型的字段。
    type valueCtx struct {
        Context
        key, val any
    }
    
### cancelCtx
    type canceler interface {
        // removeFromParent 表示是否从父上下文中删除自身
        // err 表示取消的原因
        cancel(removeFromParent bool, err error)
        // Done 返回一个管道，用于通知取消的原因
        Done() <-chan struct{}
    }
    在创建上下文时通过闭包将其包装为返回值以供外界调用

### timerCtx
    在concelCtx 的基础之上增加超时机制，context包提供两种创建函数，分别
    WithDeadline: 指定一个具体的超时时间
    WithTimeout: 指定一个超时的时间间隔


## 锁

### 互斥锁
    
    适合读操作与写操作频率都差不多情况
    sync.Locker
    
    type Locker interface {
        // 加锁
        Lock()
        // 解锁
        Unlock()
    }

### 读写锁
    sync.RWMutex
    适合一些读多写少的数据：
        如果获得了读锁，其他协程进行写操作会阻塞，其他协程进行读操作时不会阻塞
        如果获得了写锁，其他协程进行写操作会阻塞，其他协程进行读操作时会阻塞
    
    // 加读锁
    func (rw *RWMutex) RLock()
    
    // 尝试加读锁
    func (rw *RWMutex) TryRLock() bool
    
    // 解读锁
    func (rw *RWMutex) RUnlock()
    
    // 加写锁
    func (rw *RWMutex) Lock()
    
    // 尝试加写锁
    func (rw *RWMutex) TryLock() bool
    
    // 解写锁
    func (rw *RWMutex) Unlock()
    

### 条件变量
    通讯机制：
    func NewCond(l Locker) *Cond
    
    // 阻塞等待条件生效，直到被唤醒
    func (c *Cond) Wait()
    
    // 唤醒一个因条件阻塞的协程
    func (c *Cond) Signal()
    
    // 唤醒所有因条件阻塞的协程
    func (c *Cond) Broadcast()



## sync

### Once
    懒加载，当在使用一些数据结构时，数据结构过大可以使用懒加载，使用时才会初始化该数据结构。


### Pool
    用于存储临时对象便于复用。
    方法：
    // 申请一个对象
    func (p *Pool) Get() any
    
    // 放入一个对象
    func (p *Pool) Put(x any)

    注意：
    1. 临时对象：只适合存放临时对象，不适合存放网络连接、数据库连接。
    2. 不可预知：申请对象时，无法预知这个对象是新创建还是复用的。
    3. 并发安全：New函数由使用者传入，其并发安全性也由自己维护。

### Map
    并发安全Map的实现。

    // 根据一个key读取值，返回值会返回对应的值和该值是否存在
    func (m *Map) Load(key any) (value any, ok bool)
    
    // 存储一个键值对
    func (m *Map) Store(key, value any)
    
    // 删除一个键值对
    func (m *Map) Delete(key any)
    
    // 如果该key已存在，就返回原有的值，否则将新的值存入并返回，当成功读取到值时，loaded为true，否则为false
    func (m *Map) LoadOrStore(key, value any) (actual any, loaded bool)
    
    // 删除一个键值对，并返回其原有的值，loaded的值取决于key是否存在
    func (m *Map) LoadAndDelete(key any) (value any, loaded bool)
    
    // 遍历Map，当f()返回false时，就会停止遍历
    func (m *Map) Range(f func(key, value any) bool)

## 原子
    不可在细化分割的操作，执行结果要么成功要么失败。

### 类型
    atomic.Bool{}
    atomic.Pointer[]{}
    atomic.Int32{}
    atomic.Int64{}
    atomic.Uint32{}
    atomic.Uint64{}
    atomic.Uintptr{}
    atomic.Value{}

### 使用
    每一个原子类型都会提供以下三个方法：

    Load()：原子的获取值
    Swap(newVal type) (old type)：原子的交换值，并且返回旧值
    Store(val type)：原子的存储值

### CAS
    乐观锁，是一种并发条件下无锁化并发控制方式。

### Value
    atomic.Value 结构体，存储任何类型的值