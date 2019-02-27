## time
```
时间类型
  Location    //代表一个地区和失去
  Time        //代表一个纳秒精度的时间点
  Duration    //代表两个时间点之间的时间，以纳秒为单位  type Duration int64
  Time        //定时器  到达指定时间触发，且只触发一次
  Ticker      //定时器  间隔指定时间触发
  Weekday
  Month
```

>Location
>>  Local   表示当前系统时区
>>  UTC     通用协调时间
  
>Time        //定时器  到达指定时间触发，且只触发一次
>Ticker      //定时器  间隔指定时间触发  
  
Timer的实例必须通过NewTimer或AfterFunc获得
  
  
  
## Time    
  代表一个纳秒精度的时间点  
  程序中应使用 Time 类型值来保存和传递时间，而不是指针。就是说，表示时间的变量和字段，应为time.Time类型，而不是*time.Time.类型。一个Time类型值可以被多个go程同时使用。时间点可以使用Before、After和Equal方法进行比较。Sub方法让两个时间点相减，生成一个Duration类型值（代表时间段）。Add方法给一个时间点加上一个时间段，生成一个新的Time类型时间点
        var d time.Time
        fmt.Println(d)
        d=time.Date(2019, time.November, 10, 23, 0, 0, 0, time.Local)//func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
        fmt.Printf("Go launched at %s\n", d.Local())

        fmt.Println(d.Year())
        fmt.Println(d.Month())
        fmt.Println(d.Day())
        fmt.Println(d.Hour())
        fmt.Println(d.Minute())
        fmt.Println(d.Second())
        fmt.Println(d.Nanosecond()) //返回t对应的那一秒内的纳秒偏移量

        fmt.Println(d.UTC())
        fmt.Println(d.Local())

        fmt.Println(d.Clock()) //返回时、分、秒
        fmt.Println(d.Date())
        fmt.Println(d.Location())
        fmt.Println(d.Zone())//计算t所在的时区，返回该时区的规范名（如"CET"）和该时区相对于UTC的时间偏移量（单位秒）

        fmt.println(d.in(d.Local()))
        
        
func (t Time) Equal(u Time) bool //判断两个时间是否相同
func (t Time) Before(u Time) bool //如果t代表的时间点在u之前，返回真
func (t Time) After(u Time) bool  如果t代表的时间点在u之后，返回真
func (t Time) YearDay() int //返回时间点t对应的那一年的第几天
func (t Time) Weekday() Weekday  //返回时间点t对应的那一周的周几
func (t Time) Add(d Duration) Time //Add返回时间点t+d。
func (t Time) AddDate(years int, months int, days int) Time
func (t Time) Sub(u Time) Duration  //返回一个时间段t-u


Round 和 Truncate 方法
        


timer
type Timer struct { C<-chan Time }

func NewTimer(d Duration) *Timer

func (t *Timer) Reset(d Duration) bool
func (t *Timer) Stop() bool

func AfterFunc(d Duration, f func()) *Timer



```golang
ticker := time.NewTicker(5 * time.Second)
for _ = range ticker.C {
    fmt.Println(time.Now())
}
```
Ticker会在每隔一段时间执行，比如上面的例子中，每隔5秒打印一下当前时间

time包预定义了一些格式
```
const (
    ANSIC       = "Mon Jan _2 15:04:05 2006"
    UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
    RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
    RFC822      = "02 Jan 06 15:04 MST"
    RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
    RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
    RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
    RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
    RFC3339     = "2006-01-02T15:04:05Z07:00"
    RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
    Kitchen     = "3:04PM"
    // Handy time stamps.
    Stamp      = "Jan _2 15:04:05"
    StampMilli = "Jan _2 15:04:05.000"
    StampMicro = "Jan _2 15:04:05.000000"
    StampNano  = "Jan _2 15:04:05.000000000"
)
```
