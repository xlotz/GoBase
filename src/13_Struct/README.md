# 结构体

## 简介
    存储一组不同类型的数据，是一种复合类型。
    
    type Person struct {
        Name string
        Age int
        Job string
        Language []string
    }


## 声明
    结构体及其内部字段遵守大小写命名的暴露方式
    类型相邻的字段不需要重复声明类型
    字段名不能和方法名重复

    type Person struct{
        name string
        age int
    }

    type Rect struct{
        height, width, area int
        color string
    }

    

## 实例化
    prog := prog {
        Name: "jack",
        Age: 19,
        Job: "coder",
        Language: []string{"Go", "Python"},
    }
    或者
    type Person struct{
        Name string
        Age int
        Address string
        Salary float64
    }
    func NewPerson(name string, age int, address string, salary float64) *Person {
        return &Person{
            Name: name,
            Age: age,
            Address: address,
            Salary: salary,
        }
    }

    选项模式
    
    type Person struct{
        Name string
        Age int
        Address string
        Salary float64
    }
    type PersonOptions func(p *Person)
    func WithName(name string) PersonOptions{
        return func(p *Person){
            p.Name = name
        }
    }
    func WithAge(name age) PersonOptions{
        return func(p *Person){
            p.Age = age
        }
    } 
    func WithAddress(address string) PersonOptions {
	    return func(p *Person) {
		    p.Address = address
	    }
    }


## 组合
    type Person struct {
        name string
        age  int
    }
        
    type Student struct {
        p      Person
        school string
    }
        
    type Employee struct {
        p   Person
        job string
    }

## 指针
    对于结构体指针，不需要解引用可以直接访问结构体内容
    p := &Person{
        name: "jack",
        age: 18,
    }
    fmt.Println(p.age, p.name)

## 标签
    键值对形式，使用空格分隔
    type Programmer struct {
        Name     string `json:"name"`
        Age      int `yaml:"age"`
        Job      string `toml:"job"`
        Language []string `properties:"language"`
    }

## 内存对齐
    结构体字段的内存分布遵循内存对齐的规则

## 空指针
    没有字段，不占用内存，使用 unsafe.SizeOf函数计算占用的字节大小
    type Empty struct {}
    fmt.Println(unsafe.SizeOf(Empty{}))