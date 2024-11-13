# GoBase
go学习记录

## idea 配置go环境

    idea 设置

![image](images/01_go_root_sdk.png)


## go mod

    idea 不使用go path, 改使用 go mod

### idea 设置

![image](images/02_go_path.png)

![image](images/03_go_mod.png)

### 项目配置

    1. 进入项目: go_test
    2. 执行: 
        go mod init go_test
        go mod tidy
        会生成一个go.mod文件

    3. 包导入:
        import "go_test/目录名xxx/目录名xxx" 


## 函数传参

### 可变参数
#### 

    func main() {
        n := average(43, 56, 87, 12, 45, 57)
        fmt.Println(n)
    }
    
    func average(sf ...float64) float64 {
        fmt.Println(sf)
        fmt.Printf("%T \n", sf)
        var total float64
        for _, v := range sf {
            total += v
        }
        return total / float64(len(sf))
    }

#### map

    func main() {
        data := []float64{43, 56, 87, 12, 45, 57}
        n := average(data...)
        fmt.Println(n)
    }
    
    func average(sf ...float64) float64 {
        total := 0.0
        for _, v := range sf {
            total += v
        }
        return total / float64(len(sf))
    }

#### slice

    func main() {
        data := []float64{43, 56, 87, 12, 45, 57}
        n := average(data)
        fmt.Println(n)
    }
    
    func average(sf []float64) float64 {
        total := 0.0
        for _, v := range sf {
            total += v
        }
        return total / float64(len(sf))
    }

### struct
#### 特点
    1、封装
        状态（）
        行为（方法）
        出口、非出口
    2、可重用性
        继承
    3、多态
        接口
    4、重写、覆盖
#### 定义
    第一种：
    type person struct {
	    first string
	    last  string
	    age   int
    }
    调用：
    p1 := person{"James", "Bond", 20}

    第二种：方法
    type person struct {
	    first string
	    last  string
	    age   int
    }
    
    func (p person) fullName() string {
        return p.first + p.last
    }
    
    调用：
    p1 := person{"James", "Bond", 20}

#### 实现功能
    1、通过内部类型或嵌入类型实现复用
    type person struct {
        First string
        Last  string
        Age   int
    }
    
    type doubleZero struct {
        person
        LicenseToKill bool
    }
    调用：
    p1 := doubleZero{
		person: person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		LicenseToKill: true,
	}
    2、重写字段值
    type person struct {
        First string
        Last  string
        Age   int
    }
    
    type doubleZero struct {
        person
        First         string
        LicenseToKill bool
    }

    p1 := doubleZero{
        person: person{
        First: "James",
        Last:  "Bond",
        Age:   20,
        },
        First:         "Double Zero Seven",
        LicenseToKill: true,
    }
    3、重写方法
    
    type person struct {
        Name string
        Age  int
    }
    
    type doubleZero struct {
        person
        LicenseToKill bool
    }
    
    func (p person) Greeting() {
        fmt.Println("I'm just a regular person.")
    }
    
    func (dz doubleZero) Greeting() {
        fmt.Println("Miss Moneypenny, so good to see you.")
    }
    调用：
        p1 := person{
		Name: "Ian Flemming",
		Age:  44,
        }
    
        p2 := doubleZero{
            person: person{
                Name: "James Bond",
                Age:  23,
            },
            LicenseToKill: true,
        }
        p1.Greeting()
        p2.Greeting()
        p2.person.Greeting()

#### 结构体指针
    type person struct {
        name string
        age  int
    }
    
    func main() {
        p1 := &person{"James", 20}
        fmt.Println(p1)
        fmt.Printf("%T\n", p1)
        fmt.Println(p1.name)
        fmt.Println(p1.age)
    }

#### 编码解码（json <-> stream字节流 ->文件或网络IO）
    将数据结构编码为json字符串
    p1 := person{"James", "Bond", 20, 007}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)

    将json字符串解码为数据结构
    bs := []byte(`{"First":"James", "Last":"Bond", "Age":20}`)
	json.Unmarshal(bs, &p1)
#### 序列化和反序列化（json <-> object对象 ->映射）
    将数据结构转换成json字符串
    p1 := person{"James", "Bond", 20, 007}
	json.NewEncoder(os.Stdout).Encode(p1)
    将json字符串转换成数据结构
    rdr := strings.NewReader(`{"First":"James", "Last":"Bond", "Age":20}`)
	json.NewDecoder(rdr).Decode(&p1)
    
### interface
#### 

### go
### go
