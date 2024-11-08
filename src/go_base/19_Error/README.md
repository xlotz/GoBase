# 错误

## 介绍
    三种级别的错误
    error: 部分流程出错，需要处理
    panic: 很严重的问题，程序需要在处理完问题后立即退出
    fatal: 致命的问题，程序应该立即退出

### error
    本身是一个预定义的接口，只有一个Error()方法，返回字符串，用于输出错误信息。
    type error interface{
        Error() string
    }

#### 创建
    err := errors.New("这是一个错误")
    err := fmt.Errorf("这是一个格式化参数的错误")

#### 自定义错误
    // 参考errors.errorString
    func New(text string) string{
        return &errorString{text}
    }
    type errorString struct {
        s string
    }
    func (e *errorString) Error() string {
        return e.s
    }
    
#### 传递
    将错误抛给上层调用者
    type wrapError struct {
        msg string
        err error
    }
    func (e *wrapError) Error() string {
        return e.msg
    }
    func (e *wrapError) Unwrap() error {
        return e.err
    }
    调用
    err := errors.New("这是一个原始错误")
    wrapErr := fmt.Errorf("错误， %w", err)

#### 处理
    errors 提供几个函数用于检查和处理错误
    errors.Unwrap() 解包一个错误链
    errors.Is 判断错误链中是否包含指定的错误
    errors.As() 在错误链中寻找第一个类型匹配的错误，并将值符给传入的err
    加强包：github.com/pkg/errors

### panic
    表示十分严重的程序问题，需要立即停止处理该问题。

#### 创建
    内置函数：
    func panic(v any)
    
    
#### 善后
    当放生panic时，会立即退出所在函数，并执行当前函数的善后工作，如defer,
    然后层层上抛，上游函数同样的进行善后，直到程序停止。

#### 恢复
    当发生panic时，使用内置函数recover() 可以及时的处理并保障程序继续运行，必须要在defer语句中运行。
    注意：
        recover必须在defer中使用
        多次使用也只能有一个恢复panic
        闭包recover不会恢复外部函数的任何panic
        panic的参数禁止使用nil


### fatal
    一种及其严重的问题，程序会立即停止，不会执行善后工作，通常会调用os.exit退出.
