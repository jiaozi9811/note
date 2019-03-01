# time

## 时间类型
  - Location    //代表一个地区和失去
  - Time        //代表一个纳秒精度的时间点
  - Duration    //代表两个时间点之间的时间，以纳秒为单位  type Duration int64
  - Time        //定时器  到达指定时间触发，且只触发一次
  - Ticker      //定时器  间隔指定时间触发
  - Weekday
  - Month


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

### Now with timezone  [↑ top](#time)

```go
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println(nowPST())
	fmt.Println(nowUTC())
	fmt.Println(todayPST())
	fmt.Println(todayUTC())
	fmt.Println(dateTag(nowPST()))
	fmt.Println(timeTag(nowPST()))
	/*
	   2015-08-09 21:17:50.508866464 -0700 PDT
	   2015-08-10 04:17:50.508917948 +0000 UTC
	   2015-08-09 00:00:00 -0700 PDT
	   2015-08-10 00:00:00 +0000 UTC
	   20150809
	   2015080921175050897
	*/
}

// Manipulate Time with:
// func (t Time) Add(d Duration) Time
// func (t Time) AddDate(years int, months int, days int) Time

func nowPST() time.Time {
	tzone, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Returning time.Now()")
		return time.Now()
	}
	return time.Now().In(tzone)
}

func nowUTC() time.Time {
	tzone, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Returning time.Now()")
		return time.Now()
	}
	return time.Now().In(tzone).UTC()
}

func todayPST() time.Time {
	tzone, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Returning time.Now()")
		return time.Date(
			time.Now().Year(),
			time.Now().Month(),
			time.Now().Day(),
			0,
			0,
			0,
			0,
			nil,
		)
	}
	pst := time.Now().In(tzone)
	return time.Date(
		pst.Year(),
		pst.Month(),
		pst.Day(),
		0,
		0,
		0,
		0,
		tzone)
}

func todayUTC() time.Time {
	tzone, err := time.LoadLocation("")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Returning time.Now()")
		return time.Date(
			time.Now().Year(),
			time.Now().Month(),
			time.Now().Day(),
			0,
			0,
			0,
			0,
			time.UTC,
		)
	}
	utc := time.Now().In(tzone).UTC()
	return time.Date(
		utc.Year(),
		utc.Month(),
		utc.Day(),
		0,
		0,
		0,
		0,
		tzone)
}

func stringByScale(t time.Time, scale string) string {
	switch scale {
	case "micro":
		// 2014-06-23 15:37:21.12311
		return t.String()[:25]
	case "second":
		return t.String()[:19]
	case "date":
		return t.String()[:10]
	default:
		return t.String()
	}
}

func dateTag(t time.Time) string {
	ts := stringByScale(t, "date")
	return strings.Replace(ts, "-", "", -1)
}

func timeTag(t time.Time) string {
	ts := stringByScale(t, "micro")
	ts = strings.Replace(ts, "-", "", -1)
	ts = strings.Replace(ts, ":", "", -1)
	ts = strings.Replace(ts, ".", "", -1)
	ts = strings.Replace(ts, " ", "", -1)
	return ts
}
```
---
Parse time[↑ top](#time)
```
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	ct1 := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	ct2 := parseTS("2009-11-10 23:00:00")
	if !ct1.Equal(ct2) {
		panic(fmt.Errorf("expected \n%v\n%v", ct1, ct2))
	}
	fmt.Println(ct2)
	fmt.Println(parseDate0("2016-07-07"))
	fmt.Println(parseDate1("November 7, 2017"))
	fmt.Println(parseDate2("06/07/2019"))
	fmt.Println(substractDate("2016-07-07", "2015-07-07"))
	/*
	   2009-11-10 23:00:00 +0000 UTC
	   2016-07-07 00:00:00 +0000 UTC
	   2017-11-07 00:00:00 -0800 PST
	   2019-06-07 00:00:00 -0700 PDT
	   366
	*/
}

// parseDate0 parses the string-format date time.
func parseDate0(stamp string) time.Time {
	rt, err := time.Parse("2006-01-02", stamp)
	if err != nil {
		panic(err)
	}
	return rt
}

func parseDate1(date string) time.Time {
	zoneSC, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		panic(err)
	}
	pt, err := time.Parse("January 2, 2006", date)
	if err != nil {
		panic(err)
	}
	rd := time.Date(
		pt.Year(),
		pt.Month(),
		pt.Day(),
		0,
		0,
		0,
		0,
		zoneSC)
	return rd
}

func parseDate2(date string) time.Time {
	zoneSC, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		panic(err)
	}
	pt, err := time.Parse("01/02/2006", date)
	if err != nil {
		panic(err)
	}
	rd := time.Date(
		pt.Year(),
		pt.Month(),
		pt.Day(),
		0,
		0,
		0,
		0,
		zoneSC)
	return rd
}

// substractDate returns the difference between two dates.
// (dt1 - dt2)
func substractDate(dt1, dt2 string) int {
	td1 := parseDate0(dt1)
	td2 := parseDate0(dt2)
	diff := td1.Sub(td2).Hours()
	return int(diff / 24)
}

// parseTS parses the string-format time stamp.
func parseTS(stamp string) time.Time {
	stamp = strings.TrimSpace(stamp)
	if len(strings.Split(stamp, ".")) > 1 {
		mc1 := strings.Split(stamp, ".")[0]
		mc2 := strings.Split(stamp, ".")[1]
		if len(mc2) == 1 {
			stamp = mc1 + "." + mc2 + "00000"
		}
		if len(mc2) == 2 {
			stamp = mc1 + "." + mc2 + "0000"
		}
		if len(mc2) == 3 {
			stamp = mc1 + "." + mc2 + "000"
		}
		if len(mc2) == 4 {
			stamp = mc1 + "." + mc2 + "00"
		}
		if len(mc2) == 5 {
			stamp = mc1 + "." + mc2 + "0"
		}
	} else if len(strings.Split(stamp, ".")) == 1 {
		stamp = stamp + ".000000"
	}

	rt, err := time.Parse("2006-01-02 15:04:05.000000", stamp)
	if err != nil {
		panic(err)
	}
	return rt
}
```
