# reflect

tags： golang

---
https://blog.csdn.net/kjfcpua/article/details/42389823
https://my.oschina.net/90design/blog/1614820
---
在计算机科学领域，反射是指一类应用，它们能够自描述和自控制。也就是说，这类应用通过采用某种机制来实现对自己行为的描述（self-representation）和监测（examination），并能根据自身行为的状态和结果，调整或修改应用所描述行为的状态和相关的语义。


>interface及其pair的存在，是Golang中实现反射的前提，理解了pair，就更容易理解反射。
>反射就是用来检测存储在接口变量内部(值value；类型concrete type) pair对的一种机制

---
#动静类型
##静态类型
静态类型就是变量声明时的赋予的类型
##动态类型
动态类型是在运行时给变量赋值时，这个值获得的的类型(如果值为nil的时候没有动态类型)。一个变量的动态类型在运行时可能改变，这主要依赖于它的赋值（前提是这个变量是接口类型）。
```
var A interface{} // 静态类型interface{}
A = 10            // 静态类型为interface{}  动态为int
A = "String"      // 静态类型为interface{}  动态为string
```

---
>reflect包定义了两个类型：Type和Value

reflect.ValueOf()   用来获取输入参数接口中的数据的值，如果接口为空则返回0
reflect.TypeOf()用来动态获取输入参数接口中的值的类型，如果接口为空则返回nil
```
var x float64 = 3.4
fmt.Println("type:", reflect.TypeOf(x))
```
```
func TypeOf(i interface{}) Type
```
reflect.Typeof里包含了一个空接口
>当我们调用reflect.Typeof(x)的时候，x首先被保存到一个空接口中，这个空接口然后被作为参数传递。reflect.Typeof 会把这个空接口拆包（unpack）恢复出类型信息


##什么情况下使用反射
>有时候你想在运行时使用变量来处理变量，这些变量使用编写程序时不存在的信息。
>也许你正试图将来自文件或网络请求的数据映射到变量中。
>也许创建一个适用于不同类型的tool

在这些情况下，你需要使用反射。反射使您能够在运行时检查类型。它还允许您在运行时检查，修改和创建变量，函数和结构

>>Type 接口：可以表示一个Go类型

- Kind() 将返回一个常量，表示具体类型的底层基础基础类型
- Elem()方法返回指针、数组、切片、map、通道的基类型；
- 可用反射提取struct tag，还能自动分解，常用于ORM映射、数据验证等；
- 辅助判断方法Implements()、ConvertibleTo()、AssignableTo()
```
type Type unterface {
    Kind() Kind //返回底层基础类型
    Name() string
    Elem() Type //// 返回该类型的元素类型，如果该类型的Kind不是Array、Chan、Map、Ptr或Slice，会panic
}
```


Value 结构体：可以持有一个任意类型的值

- Type()将返回具体类型所对应的 reflect.Type（静态类型）
- Kind() 将返回一个常量，表示具体类型的底层类型
- Interface方法是ValueOf方法的逆，它把一个reflect.Value恢复成一个接口值：把Value中保存的类型和值的信息打包成一个接口表示并返回；如：
y,ok := v.Interface().(float64) // y 的类型被断言为 float64
fmt.Println(y)
以上可简写为这样：fmt.Println(v.Interface())    //fmt.Println会把它恢复出来
- 通道类型的反射对象：有TrySend()、TryRecv()方法；
- IsNil()方法判断反射对象保存的值是否为nil；
- CanSet方法用于测试一个Value的Settablity性质,描述的是一个反射对象能够修改创造它的那个实际存储的值的能力,判断它是否可被取地址并可被修改

>读取一个变量的值，首先需要一个reflect.Valueof( var ) 实例（reflectVal := reflect.Valueof(var)）， 同时也可以获取变量的类型

>修改一个变量的值，那么必须通过该变量的指针地址 , 取消指针的引用  。通过refPtrVal := reflect.Valueof( &var )的方式获取指针类型，你使用refPtrVal.elem( ).set（一个新的reflect.Value）来进行更改，传递给set（）的值也必须是一个reflect.value
```
	s := "String字符串"
	sPtr := reflect.ValueOf(&s)
	sPtr.Elem().Set(reflect.ValueOf("修改值1"))
	sPtr.Elem().SetString("修改值2")
	// 修改指针指向的值，原变量改变
	fmt.Println(s)
	fmt.Println(sPtr) // 要注意这是一个指针变量，其值是一个指针地址
	
```

>创建一个值，那么使用NewPtrVal := reflect.New( vartype ) 传递一个reflect.Type类型。 返回的指针类型就可以使用以上修改的方式写入值
```
	type Foo struct {
		A int `tag1:"Tag1" tag2:"Second Tag"`
		B string
	}
	fo := Foo{}
	foType := reflect.TypeOf(fo)
	foVal := reflect.New(foType)
	// foVal.Elem().Field(0).SetString("A") // 引发panic
	foVal.Elem().Field(0).SetInt(1)
	foVal.Elem().Field(1).SetString("B")
	f2 := foVal.Elem().Interface().(Foo)
	fmt.Printf("%+v, %d, %s\n", f2, f2.A, f2.B)
```
##Call()

反射中 函数 和 方法 的类型（Type）都是 reflect.Func，如果要调用函数的话，可以通过 Value 的 Call() 方法

```
type MyStruct struct {
    name string
}
 
func (this *MyStruct) GetName() string {
    return this.name
}
 
type IStruct interface {
    GetName() string
}
 
func main() {
    // TypeOf
    s := "this is string"
    fmt.Println(reflect.TypeOf(s)) // output: "string"
 
    // object TypeOf
    a := new(MyStruct)
    a.name = "yejianfeng"
    typ := reflect.TypeOf(a)
    fmt.Println(typ)        // output: "*main.MyStruct"
    fmt.Println(typ.Elem()) // output: "main.MyStruct"
 
    // reflect.Type Base struct
    fmt.Println(typ.NumMethod())                   // 1
    fmt.Println(typ.Method(0))                     // {GetName  func(*main.MyStruct) string <func(*main.MyStruct) string Value> 0}
    fmt.Println(typ.Name())                        // ""
    fmt.Println(typ.PkgPath())                     // ""
    fmt.Println(typ.Size())                        // 8
    fmt.Println(typ.String())                      // *main.MyStruct
    fmt.Println(typ.Elem().String())               // main.MyStruct
    fmt.Println(typ.Elem().FieldByIndex([]int{0})) // {name main string  0 [0] false}
    fmt.Println(typ.Elem().FieldByName("name"))    // {name main string  0 [0] false} true
 
    fmt.Println(typ.Kind() == reflect.Ptr)                              // true
    fmt.Println(typ.Elem().Kind() == reflect.Struct)                    // true
    fmt.Println(typ.Implements(reflect.TypeOf((*IStruct)(nil)).Elem())) // true
 
    fmt.Println(reflect.TypeOf(12.12).Bits()) // 64, 因为是float64
 
    cha := make(chan int)
    fmt.Println(reflect.TypeOf(cha).ChanDir()) // chan
 
    var fun func(x int, y ...float64) string
    var fun2 func(x int, y float64) string
    fmt.Println(reflect.TypeOf(fun).IsVariadic())  // true
    fmt.Println(reflect.TypeOf(fun2).IsVariadic()) // false
    fmt.Println(reflect.TypeOf(fun).In(0))         // int
    fmt.Println(reflect.TypeOf(fun).In(1))         // []float64
    fmt.Println(reflect.TypeOf(fun).NumIn())       // 2
    fmt.Println(reflect.TypeOf(fun).NumOut())      // 1
    fmt.Println(reflect.TypeOf(fun).Out(0))        // string
 
    mp := make(map[string]int)
    mp["test1"] = 1
    fmt.Println(reflect.TypeOf(mp).Key()) //string
 
    arr := [1]string{"test"}
    fmt.Println(reflect.TypeOf(arr).Len()) // 1
 
    fmt.Println(typ.Elem().NumField()) // 1
 
    // MethodByName, Call
    b := reflect.ValueOf(a).MethodByName("GetName").Call([]reflect.Value{})
    fmt.Println(b[0]) // output: "yejianfeng"
}
```
<!--stackedit_data:
eyJoaXN0b3J5IjpbMTIwNDMyMTMyNV19
-->