tags： golang

- defer: Stack(堆)(Last-In-First-Out)
- goroutine: Queue(队列)(First-In-First-Out)

go的并发同步模型来自一个叫作通信顺序进程(communicating sequential processes,CSP)的范型
在goroutine之间同步和传递数据的数据类型是通道(channel)


>并发(concurrency)	同时管理很多事情
>并行(parallelism)	同时做很多事情
>并行是让不同的代码片段同时在不同的物理处理器执行


任何 Go 程序都必须有的 main() 函数也可以看做是一个协程，尽管它并没有通过 go 来启动。协程可以在程序初始化的过程中运行（在 init() 函数中）


>ch := make(chan int) //无缓冲通道(可读可写通道)
>ch := make(chan int, 10)//缓冲为10的通道（可读可写通道）
>ch := make(chan<- int)//无缓冲通道（只写通道）
>ch := make(<-chan int)//无缓冲通道（只读通道）


## 通道(channel)
unbuffered :=make(chan type)	无缓冲通道
buffered :=make(chan type,len)	有缓冲通道

var identifier chan datatype

通道是先进先出（FIFO）的结构所以可以保证发送给他们的元素的顺序
使用 make() 函数来给通道分配内存

无缓冲通道(unbuffered channel)
在接收前没有能力保存任何值的通道
这种类型的通道要求发送goroutine和接收goroutine同时准备好，才能完成发送和接收操作


向通道发送值或指针需要用到 <- c操作符
从通道接收一个值或指针时，<-运算符在要操作的通道变量的左侧

向通道发送数据
ch <- int1     //表示：用通道 ch 发送变量 int1
从通道发出数据
int2 := <- ch   //表示：变量 int2 从通道 ch（一元运算的前缀操作符，前缀 = 接收）接收数据（获取新值）
int3,ok :=<-ch
<- ch可以单独调用获取通道的(下一个)值，当前值会被丢弃，但是可以用来验证
if <- ch != 1000{   ...  }

defer close(ch)    //关闭通道
不要在接收端关闭channel，接收端通常无法判断发送端是否还会发送数值。
在发送端关闭channel
如果channel在关闭时还有值，可以使用接收表达式取出

v, ok := <-ch// ok is true if v received value   使用ok操作符检测通道是否被关闭
if !ok {  break  }
当关闭或者阻塞的时候使用 break


发送channel
chan<-T   只能向通道发送值，不能接受值
接收channel
<-chan T

func fun(ch chan<- int) <-chan int{...}
ch是只能写的channel
返回的chan是只能读的channel


使用 select 切换协程
select和switch控制语句非常相似，也被称作通信开关
select {
case u:= <- ch1:        ...
case v:= <- ch2:        ...
default:  	...	// no value ready to be received  
}
select语句需要放在一个单独的goruntine中，这样即使select语句阻塞了，也不会造成死锁





```
给程序1 秒时间输出上万个整数
import (	"fmt",	"time")
func main() {
	ch1 := make(chan int)
	go pump(ch1)
	go suck(ch1) // tons of numbers appear
	time.Sleep(1e9)
}
func pump(ch chan int) {
	for i := 0; ; i++ {		ch <- i}
}
func suck(ch chan int) {
	for {		fmt.Println(<-ch) }
}

-------------------------------------
c := make(chan int)
go func() {
	time.Sleep(15 * 1e9)
	x := <-c
	fmt.Println("received", x)
}()
fmt.Println("sending", 10)
c <- 10
fmt.Println("sent", 10)

/* Output:
sending 10
(15 s later):
received 10
sent 10
*/
```

```
ch:=make(chan int,10)
for i:=0;i<10;i++{
	ch<-i
}
fmt.Println(ch,len(ch),cap(ch))
sum:=len(ch)
//不能写成for i:=0;i<len(ch);i++
for i:=0;i<sum;i++{
	fmt.Println(<-ch)
}
```

**并发编程的难度在于协调，而协调就要通过交流。从这个角度看来，并发单元间的通信是最大的问题。**

----------
在工程上，有两种最常见的并发通信模型：共享数据和消息。

 - 共享数据是指多个并发单元分别保持对同一个数据的引用，实现对该数据的共享。被共享的数据可能有多种形式，比如内存数据块、磁盘文件、网络数据等。在实际工程应用中最常见的无疑是内存了，也就是常说的共享内存。
 - 消息机制认为每个并发单元是自包含的、独立的个体，并且都有自己的变量，但在不同并发单元间这些变量不共享。每个并发单元的输入和输出只有一种，那就是消息。这有点类似于进程的概念，每个进程不会被其他进程打扰，它只做好自己的工作就可以了。不同进程间靠消息来通信，它们不会共享内存。

###出让时间片
&nbsp;&nbsp;&nbsp;&nbsp;我们可以在每个goroutine中控制何时主动出让时间片给其他goroutine，这可以使用runtime包中的Gosched()函数实现。

###Goroutine(协程)
进程、线程、协程的关系和区别

 - 进程拥有自己独立的堆和栈，既不共享堆，亦不共享栈，进程由操作系统调度。
 - 线程拥有自己独立的栈和共享的堆，共享堆，不共享栈，线程亦由操作系统调度(标准线程是的)。
 - 协程和线程一样共享堆，不共享栈，协程由程序员在协程的代码里显示调度。

>&nbsp;&nbsp;&nbsp;&nbsp;goroutine是Go语言中的轻量级线程实现，由Go运行时(runtime)管理。在一个函数调用前加上go关键字，这次调用就会在一个新的goroutine中并发执行。当被调用的函数返回时，这个goroutine也自动结束了。
&nbsp;&nbsp;&nbsp;&nbsp;需要注意的是，如果这个函数有返回值，那么这个返回值会被丢弃


  [1]: https://upload.wikimedia.org/wikipedia/commons/2/25/Insertion_sort_animation.gif
