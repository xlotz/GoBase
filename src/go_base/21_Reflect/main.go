package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func main() {

	// 声明一个any类型的变量
	var eface1, eface2, eface3 any

	fmt.Println("Kind ...")
	// 赋值
	eface1 = 100
	// 获取类型
	fmt.Println(reflect.TypeOf(eface1).Kind())

	eface2 = map[string]int{}
	rType1 := reflect.TypeOf(eface2)
	// key() 返回map的键反射类型
	fmt.Println(rType1.Key().Kind())
	fmt.Println(rType1.Elem().Kind())

	//赋值指针
	fmt.Println("Elem ...")
	eface3 = new(strings.Builder)
	rType2 := reflect.TypeOf(eface3)
	// 拿到指针指向元素的反射类型
	vType := rType2.Elem()
	// 输出包路径
	fmt.Println(vType.PkgPath())
	// 输出器名称
	fmt.Println(vType.Name())

	// Size 获取对应类型所占的字节大小
	fmt.Println("Size ...")
	fmt.Println(reflect.TypeOf(0).Size())
	fmt.Println(reflect.TypeOf("").Size())
	fmt.Println(reflect.TypeOf(complex(0, 0)).Size())
	fmt.Println(reflect.TypeOf(0.1).Size())
	fmt.Println(reflect.TypeOf([]string{}).Size())

	// Comparable 判断一个类型是否可以被比较
	fmt.Println("Comparable ...")
	fmt.Println(reflect.TypeOf("Hello world!").Comparable())
	fmt.Println(reflect.TypeOf(1024).Comparable())
	fmt.Println(reflect.TypeOf([]int{}).Comparable())
	fmt.Println(reflect.TypeOf(struct{}{}).Comparable())

	// Implements 判断类型是否实现某一个接口
	fmt.Println("Implements ...")
	rIface3 := reflect.TypeOf(new(MyInterface)).Elem()
	fmt.Println(reflect.TypeOf(new(MyStruct)).Elem().Implements(rIface3))
	fmt.Println(reflect.TypeOf(new(HisStruct)).Elem().Implements(rIface3))

	// ConvertibleTo判断类型是否实现某一个接口
	fmt.Println("ConvertibleTo ...")
	fmt.Println(reflect.TypeOf(new(MyStruct)).Elem().ConvertibleTo(rIface3))
	fmt.Println(reflect.TypeOf(new(HisStruct)).Elem().ConvertibleTo(rIface3))

	// 值
	fmt.Println("reflect.Value")
	str := "Hello world!"
	reflectV := reflect.ValueOf(str)
	fmt.Println(reflectV)

	// Type 类型
	fmt.Println(reflectV.Type())

	fmt.Println("获取值的元素反射值...")
	num := new(int)
	*num = 123123
	rValue := reflect.ValueOf(num).Elem()
	fmt.Println(rValue.Interface())

	fmt.Println("获取一个反射值的指针...")
	num2 := 1023
	ele := reflect.ValueOf(&num2).Elem()
	fmt.Println("&num2", &num2)
	fmt.Println("Addr", ele.Addr())
	fmt.Println("UnsafeAddr", unsafe.Pointer(ele.UnsafeAddr()))
	fmt.Println("Pointer", unsafe.Pointer(ele.Addr().Pointer()))
	fmt.Println("UnsafePointer", ele.Addr().UnsafePointer())

	dic := map[string]int{}
	ele2 := reflect.ValueOf(&dic).Elem()
	println(dic)
	fmt.Println("Addr", ele2.Addr())
	fmt.Println("UnsafeAddr", *(*unsafe.Pointer)(unsafe.Pointer(ele2.UnsafePointer())))
	fmt.Println("Pointer", unsafe.Pointer(ele2.Pointer()))
	fmt.Println("UnsafePointer", ele2.UnsafePointer())

	fmt.Println("设置 num 值")

	rValue2 := reflect.ValueOf(num)
	// 获取指针指向的元素
	ele3 := rValue2.Elem()
	fmt.Println(ele3.Interface())
	ele3.SetInt(111)
	fmt.Println(ele3.Interface())

	fmt.Println("获取 reflectV 值")
	if v, ok := reflectV.Interface().(string); ok {
		fmt.Println(v)
	}

	fmt.Println("通过反射获取函数的所有信息")
	rType3 := reflect.TypeOf(Max)
	fmt.Println(rType3.Name(), rType3.NumIn(), rType3.NumOut())
	rParamType := rType3.In(0)
	fmt.Println(rParamType.Kind())
	rResType := rType3.Out(0)
	fmt.Println(rResType.Kind())

	//fmt.Println("通过反射调用函数")
	//rResValue := rType3.
	rType4 := reflect.TypeOf(new(Person)).Elem()
	fmt.Println(rType4.NumField())
	for i := 0; i < rType4.NumField(); i++ {
		StructField := rType4.Field(i)
		fmt.Println(StructField.Index, StructField.Name, StructField.Type, StructField.Offset, StructField.IsExported())
	}

	fmt.Println("修改字段")
	rValue5 := reflect.ValueOf(&Person{
		Name:    "",
		Age:     0,
		Address: "",
		money:   0,
	}).Elem()

	name := rValue5.FieldByName("Name")
	if (name != reflect.Value{}) {
		name.SetString("jack")
	}
	fmt.Println(rValue5.Interface())

	money := rValue5.FieldByName("money")
	if (money != reflect.Value{}) {
		p := reflect.NewAt(money.Type(), money.Addr().UnsafePointer())
		field := p.Elem()
		field.SetInt(164)
	}
	fmt.Printf("%+v\n", rValue5.Interface())

	fmt.Println("访问方法")
	rType6 := reflect.TypeOf(new(Person)).Elem()
	fmt.Println(rType6.NumMethod())
	for i := 0; i < rType6.NumMethod(); i++ {
		method := rType6.Method(i)
		fmt.Println(method.Index, method.Name, method.Type, method.IsExported())
		fmt.Println("方法参数")
		for i := 0; i < method.Func.Type().NumIn(); i++ {
			fmt.Println(method.Func.Type().In(i).String())
		}
		fmt.Println("方法返回值")
		for i := 0; i < method.Func.Type().NumOut(); i++ {
			fmt.Println(method.Func.Type().Out(i).String())
		}
	}

	// 调用方法
	talk := rValue5.MethodByName("Talk")
	if (talk != reflect.Value{}) {
		res := talk.Call([]reflect.Value{reflect.ValueOf("hello, reflect!")})
		for _, re := range res {
			fmt.Println(re.Interface())
		}
	}

	// 返回指定反射值的指针反射值
	fmt.Println("返回指定反射值的指针反射值")
	rValue6 := reflect.New(reflect.TypeOf(*new(string)))
	rValue6.Elem().SetString("hello world!")
	fmt.Println(rValue6.Elem().Interface())

	// 创建结构体反射值
	fmt.Println("创建结构体反射值")
	rType5 := reflect.TypeOf(new(Person)).Elem()
	person := reflect.New(rType5).Elem()
	fmt.Println(person.Interface())

	fmt.Println("反射创建切片")
	rValue7 := reflect.MakeSlice(reflect.TypeOf(*new([]int)), 10, 10)
	for i := 0; i < 10; i++ {
		rValue7.Index(i).SetInt(int64(i))
	}
	fmt.Println(rValue7.Interface())

	fmt.Println("反射创建Map")
	rValue8 := reflect.MakeMapWithSize(reflect.TypeOf(*new(map[string]int)), 10)
	rValue8.SetMapIndex(reflect.ValueOf("a"), reflect.ValueOf(1))
	fmt.Println(rValue8.Interface())

	fmt.Println("反射创建管道")
	makeChan := reflect.MakeChan(reflect.TypeOf(new(chan int)).Elem(), 0)
	fmt.Println(makeChan.Interface())

	fmt.Println("反射创建函数")
	fn := reflect.MakeFunc(reflect.TypeOf(new(func(int))).Elem(), func(args []reflect.Value) (results []reflect.Value) {
		for _, arg := range args {
			fmt.Println(arg.Interface())
		}
		return nil
	})
	fmt.Println(fn.Type())
	fn.Call([]reflect.Value{reflect.ValueOf(1024)})

	fmt.Println("完全相等: 切片")
	a := make([]int, 100)
	b := make([]int, 100)
	fmt.Println(reflect.DeepEqual(a, b))

	fmt.Println("完全相等: 切片")

	mike := Person2{
		Name:   "mike",
		Age:    39,
		Father: nil,
	}

	jack := Person2{
		Name:   "jack",
		Age:    30,
		Father: &mike,
	}
	tom := Person2{
		Name:   "tom",
		Age:    18,
		Father: &mike,
	}

	fmt.Println(reflect.DeepEqual(mike, jack))
	fmt.Println(reflect.DeepEqual(tom, jack))
	fmt.Println(reflect.DeepEqual(jack, jack))

}

type MyInterface interface {
	My() string
}
type MyStruct struct {
}

func (m MyStruct) My() string {
	return "my"
}

type HisStruct struct {
}

func (h HisStruct) String() string {
	return "his"
}

// 通过反射获取函数的一切信息
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	money   int
}

func (p Person) Talk(msg string) string {
	return msg
}

type StructField struct {
	Name    string
	PkgPath string
	Type    Type
	//Tag 	StructTag
	Offset    uintptr
	Index     []int
	Anonymous bool
}

type Type interface {
	Field(i int) StructField
}

type Method struct {
	Name    string
	PkgPath string
	Type    Type
	//Func Value
	Index int
}

type Person2 struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Father any
}
