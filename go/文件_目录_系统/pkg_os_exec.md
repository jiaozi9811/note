[TOC]

## exec进程
os 包及其子包 os/exec 提供了创建进程的方法
应该优先使用 os/exec 包。因为 os/exec 包依赖 os 包中关键创建进程的 API


### 创建进程
Go 语言中，Linux 下创建进程使用的系统调用是 clone


### 系统调用fork、execve、wait和exit
fork：允许一进程（父进程）创建一新进程（子进程）。具体做法是，新的子进程几近于对父进程的翻版：子进程获得父进程的栈、数据段、堆和执行文本段的拷贝。可将此视为把父进程一分为二。
exit(status)：终止一进程，将进程占用的所有资源（内存、文件描述符等）归还内核，交其进行再次分配。参数 status 为一整型变量，表示进程的退出状态。父进程可使用系统调用 wait() 来获取该状态。
wait(&status)目的有二：其一，如果子进程尚未调用 exit() 终止，那么 wait 会挂起父进程直至子进程终止；其二，子进程的终止状态通过 wait 的 status 参数返回。
execve(pathname, argv, envp) 加载一个新程序（路径名为 pathname，参数列表为 argv，环境变量列表为 envp）到当前进程的内存。这将丢弃现存的程序文本段，并为新程序重新创建栈、数据段以及堆。通常将这一动作称为执行一个新程序。

Go 语言中，没有直接提供 fork 系统调用的封装，而是将 fork 和 execve 合二为一，提供了 syscall.ForkExec。如果想只调用 fork，得自己通过 syscall.Syscall(syscall.SYS_FORK, 0, 0, 0) 实现


## os.Process 及其相关方法
```go
type Process struct {
	Pid    int
	handle uintptr      // handle is accessed atomically on Windows
	isdone uint32       // process has been successfully waited on, non zero if true
	sigMu  sync.RWMutex // avoid race between wait and signal
}
```

## StartProcess 创建的进程的相关信息
>func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)   //程序名、命令行参数、属性
>func FindProcess(pid int) (*Process, error)通过 pid 查找一个运行中的进程。即使 pid 对应的进程不存在


## os/exec运行外部命令
exec包装了 os.StartProcess 函数以便更容易的重定向标准输入和输出，使用管道连接I/O，以及作其它的一些调整

>func LookPath(file string) (string, error)
>在环境变量PATH指定的目录中搜索可执行文件，如file中有斜杠，则只在当前目录搜索。返回完整路径或者相对于当前目录的一个相对路径

```go
//Cmd类型
type Cmd struct {
	Path string // 将要执行的命令的路径
	Args []string  //保管命令的参数，包括命令名作为第一个参数；如果为空切片或者nil，相当于无参数命令
	Env []string  //指定进程的环境，如为nil，则是在当前进程的环境下执行
	Dir string    //指定命令的工作目录。如为空字符串，会在调用者的进程当前目录下执行

	Stdin io.Reader
	Stdout io.Writer
	Stderr io.Writer
	ExtraFiles []*os.File //指定额外被新进程继承的已打开文件流
	SysProcAttr *syscall.SysProcAttr  //保管可选的、各操作系统特定的sys执行属性
	Process *os.Process   //是底层的，只执行一次的进程
	ProcessState *os.ProcessState //包含一个已经存在的进程的信息，只有在调用Wait或Run后才可用
    
	finished        bool            // when Wait was called

	ctx             context.Context  nil means none
	lookPathErr     error           // LookPath error, if any.
	childFiles      []*os.File
	closeAfterStart []io.Closer
	closeAfterWait  []io.Closer
	goroutine       []func() error
	errch           chan error // one send per goroutine
	waitDone        chan struct{}
}
``` 

## exec.Command
func Command(name string, arg ...string) *Cmd
一般通过exec.Command函数产生 Cmd 实例

>接下来一般有两种写法：
>	- exec.Output(),返回([]byte, error)
>	- 调用 Start()，接着调用 Wait()，然后会阻塞直到命令执行完成；只返回error
>	- 调用 Run()，它内部会先调用 Start()，接着调用 Wait(),只返回error
 
 func (c *Cmd) Start() error
 func (c *Cmd) Wait() error
 
 func (c *Cmd) Run() error  //执行命令，并阻塞直到完成
 
 func (c *Cmd) Output() ([]byte, error)   //执行命令并返回标准输出的切片
 
 func (c *Cmd) CombinedOutput() ([]byte, error) //执行命令并返回标准输出和错误输出合并的切片
 
 
 ## StdoutPipe、StderrPipe 和 StdinPipe
 
 >func (c *Cmd) StdoutPipe() (io.ReadCloser, error)
 
 >func (c *Cmd) StderrPipe() (io.ReadCloser, error)
 
 >func (c *Cmd) StdinPipe() (io.WriteCloser, error)
 
 
 
 ## 进程的属性和控制
 os.Getpid()    //获得进程ID 
 os.Getppid()   //获得父进程ID
 
 os.Getuid()  //获取用户id
 os.Getgid()  //获取用户组
 
 os.Getgroups()
