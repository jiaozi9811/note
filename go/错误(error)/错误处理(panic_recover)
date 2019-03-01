defer-panic-and-recover

 Go 检查和报告错误条件的惯有方式：
产生错误的函数会返回两个变量，一个值和一个错误码；如果后者是 nil 就是成功，非 nil 就是发生了错误
为了防止发生错误时正在执行的函数（如果有必要的话甚至会是整个程序）被中止，在调用函数后必须检查错误

预先定义的 error 接口类型
type error interface {
    Error() string
}


定义错误
err:=errors.New("error_info")

func New(text string) error
errors.Errorf("errorinfo")




panic
当错误条件（我们所测试的代码）很严苛且不可恢复，程序不能继续运行时，可以使用 panic 函数产生一个中止程序的运行时错误
panic 接收一个做任意类型的参数，通常是字符串，在程序死亡时被打印出来。在Go 运行时负责中止程序并给出调试信息



panic 会导致栈被展开直到 defer 修饰的 recover() 被调用或者程序中止

recover
recover 只能在 defer 修饰的函数中使用：用于取得 panic 调用中传递过来的错误值，如果是正常执行，调用 recover 会返回 nil，且没有其它效果
