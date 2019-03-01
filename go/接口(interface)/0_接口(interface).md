# 接口(interface)

tags： golang

---
一个接口表示一组确定的方法（method）集合。一个接口变量能存储任意的具体值（这里的具体concrete就是指非接口的non-interface），只要这个具体值所属的类型实现了这个接口的所有方法


一个接口类型变量存储了一个pair：**赋值给这个接口变量的具体值，以及这个值的类型描述符(value, type)**。更进一步的讲，这个”值”是实现了这个接口的底层具体数据项（underlying concrete data item),而这个“类型”是描述了那个项（item）的全类型（full type）
```golang
var r io.Reader
tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
if err != nil {
    return nil, err
}
r = tty
```
>r就包含了(value, type)对，即(tty, *os.File)
```
var w io.Writer
w = r.(io.Writer)
```

```
var x interface{}  // x is nil and has static type interface{}
var v *T           // v has value nil, static type *T
x = 42             // x has value 42 and dynamic type int
x = v              // x has value (*T)(nil) and dynamic type *T
```
<!--stackedit_data:
eyJoaXN0b3J5IjpbMTQzNjg3NzQ5NV19
-->