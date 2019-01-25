testing

测试用例有四种形式： 
TestXxxx(t *testing.T) // 基本测试用例 
BenchmarkXxxx(b *testing.B) // 压力测试(基准测试)的测试用例 
Example_Xxx() // 测试控制台输出的例子 
TestMain(m *testing.M) // 测试Main函数

单元测试
创建名称为_test.go结尾的文件，包含TestXxx函数。该文件被排除在正常的程序之外，运行go test时被包含
func TestXxx(t *testing.T)
Xxx 可以是任何字母数字字符串，但是第一个字母不能是小些字母
在函数中，使用Error，Fail或相关方法发出失败信号

T类型
单元测试中，传递给测试函数的参数是*testing.T类型。它用于管理测试状态并指出格式化测试日志。测试日志在测试执行中不断累积，在测试完成时转储至标准输出
当测试调用FailNow,Fatal,Fatalf,SkipNow,Skip或Skipf时，测试结束
Fail: 测试失败，测试继续，也就是之后的代码依然会执行
FailNow: 测试失败，测试中断
SkipNow: 跳过测试，测试中断
Log: 输出信息
Logf: 输出格式化的信息
Skip: 相当于 Log + SkipNow
Skipf: 相当于 Logf + SkipNow
Error: 相当于 Log + Fail
Errorf: 相当于 Logf + Fail
Fatal: 相当于 Log + FailNow
Fatalf: 相当于 Logf + FailNow


并行测试(Parallel)
func (t *T) Parallel() {
当前测试只会与其他带有 Parallel 方法的测试并行进行测试


go help testflag

基准测试(压力测试)
func BenchmarkXxx(b *testing.B)
通过go test -bench flag执行 go test -bench=.||go test -bench . -count=5
测试函数会允许目标代码b.N次。测试期间，会调整b.N直到测试函数持续足够长的时间

计时
StartTimer：开始对测试进行计时。该方法会在基准测试开始时自动被调用，我们也可以在调用 StopTimer 之后恢复计时；
StopTimer：停止对测试进行计时。当你需要执行一些复杂的初始化操作，并且你不想对这些操作进行测量时，就可以使用这个方法来暂时地停止计时。
ResetTimer：对已经逝去的基准测试时间以及内存分配计数器进行清零。对于正在运行中的计时器，这个方法不会产生任何效果。本节开头有使用示例。


并行(RunParallel)
RunParallel 会创建出多个 goroutine，并将 b.N 分配给这些 goroutine 执行，其中 goroutine 数量的默认值为 GOMAXPROCS


内存统计
ReportAllocs 方法用于打开当前基准测试的内存统计功能， 与 go test 使用 -benchmem 标志类似，但 ReportAllocs 只影响那些调用了该函数的基准测试



子测试与子基准测试
func (b *B) Run(name string, f func(b *B)) bool
子测试，又叫 命名测试(named tests)，它意味着您现在可以拥有嵌套测试，这对于自定义（和过滤）给定测试的示例非常有用

go test -run ''      # Run 所有测试。
go test -run Foo     # Run 匹配 "Foo" 的顶层测试，例如 "TestFooBar"。
go test -run Foo/A=  # 匹配顶层测试 "Foo"，运行其匹配 "A=" 的子测试。
go test -run /A=1    # 运行所有匹配 "A=1" 的子测试。


示例测试(ExampleXxx)
示例函数以 Example 开头，如果示例函数包含以 "Output" 开头的行注释，在运行测试时，go 会将示例函数的输出和 "Output" 注释中的值做比较
如果示例函数没有 "Output" 注释，该示例函数只会被编译而不会被运行


对示例函数的命名有如下约定：
包级别的示例函数，直接命名为 func Example() { ... }
函数 F 的示例，命名为 func ExampleF() { ... }
类型 T 的示例，命名为 func ExampleT() { ... }
类型 T 上的 方法 M 的示例，命名为 func ExampleT_M() { ... }

通常，示例代码会放在单独的示例文件中，命名为 example_test.go。可以查看 io 包中的 example_test.go 了解示例的编写


TestMain
func TestMain(m *testing.M)
如果测试文件中包含该函数，那么生成的测试将调用 TestMain(m)，而不是直接运行测试。TestMain 运行在主 goroutine 中, 可以在调用 m.Run 前后做任何设置和拆卸。注意，在 TestMain 函数的最后，应该使用 m.Run 的返回值作为参数调用 os.Exit


测试覆盖率
使用的是 cover 相关的工具（go test -cover、go tool cover）。虽然 testing 包提供了 cover 相关函数，不过它们是给 cover 的工具使用的



httptest-HTTP测试辅助工具
