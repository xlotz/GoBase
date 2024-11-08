# 条件控制

## if else
    // 
    if expr {
    }
    //
    if expr {
    }else {
    }
    
    // else if

    if expr {
    }else if expr2 {
    }else if expr3 {
    }else {
    }

## switch
    多分支的判断语句
    switch expr {
        case case1:
            value1
        case case2:
            value2
        default:
            default value3
    }


## label
    标签语句，给一个代码块打标签，可以是goto， break, continue
    func main(){
        A:
            a := 1
        B:
            b := 2
    }

## goto
    将控制权传递给在同一函数中对应标签的语句
    func main(){
        a := 1
        if a == 1{
            goto A
        }else{
            fmt.Println("b")
        }
    A:
        fmt.Println("a")
    }