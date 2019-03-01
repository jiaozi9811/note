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
        
## 比较时间
func (t Time) Equal(u Time) bool //判断两个时间是否相同
func (t Time) Before(u Time) bool //如果t代表的时间点在u之前，返回真
func (t Time) After(u Time) bool  如果t代表的时间点在u之后，返回真


func (t Time) YearDay() int //返回时间点t对应的那一年的第几天
func (t Time) Weekday() Weekday  //返回时间点t对应的那一周的周几
func (t Time) Add(d Duration) Time //Add返回时间点t+d。
func (t Time) AddDate(years int, months int, days int) Time
func (t Time) Sub(u Time) Duration  //返回一个时间段t-u
func Since(t Time) Duration	//返回从t到现在经过的时间，等价于time.Now().Sub(t)
func (t Time) Format(layout string) string
func Parse(layout, value string) (Time, error)


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

calendar[top](#time)
```
else { 
package main

import (
	"fmt"
	"time"
)

func main() {
	ct := time.Date(2014, time.July, 10, 07, 0, 0, 0, time.UTC)
	if weekDay(ct) != "Thu" {
		panic(fmt.Errorf("expected Thu but %v", weekDay(ct)))
	}

	ct1 := time.Date(2014, time.July, 11, 07, 0, 0, 0, time.UTC)
	str1, bl1 := isWeekend(ct1)
	if str1 != "Fri" || bl1 != false {
		panic(fmt.Errorf("expected %v\n%v", str1, bl1))
	}

	ct2 := time.Date(2014, time.July, 12, 03, 0, 0, 0, time.UTC)
	str2, bl2 := isWeekend(ct2)
	if str2 != "Sat" || bl2 != true {
		panic(fmt.Errorf("expected %v\n%v", str2, bl2))
	}

	ct3 := time.Date(2014, time.July, 13, 17, 0, 0, 0, time.UTC)
	str3, bl3 := isWeekend(ct3)
	if str3 != "Sun" || bl3 != true {
		panic(fmt.Errorf("expected %v\n%v", str3, bl3))
	}

	newyearsday := time.Date(2014, time.January, 01, 07, 0, 0, 0, time.UTC)
	if str, bl := isUSHoliday(newyearsday); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	martinluther := time.Date(2014, time.January, 20, 07, 0, 0, 0, time.UTC)
	if str, bl := isUSHoliday(martinluther); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	valentine := time.Date(2014, time.February, 14, 07, 0, 0, 0, time.UTC)
	if str, bl := isUSHoliday(valentine); bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	if str, bl := isHalloweenValentine(valentine); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	washington := time.Date(2014, time.February, 17, 07, 0, 0, 0, time.UTC)
	if str, bl := isUSHoliday(washington); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	mother := time.Date(2014, time.May, 11, 07, 0, 0, 0, time.UTC)
	if str, bl := isMotherFather(mother); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	memorial := time.Date(2014, time.May, 26, 07, 0, 0, 0, time.UTC)
	if str, bl := isUSHoliday(memorial); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	father := time.Date(2014, time.June, 15, 07, 0, 0, 0, time.UTC)
	if str, bl := isMotherFather(father); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	independence := time.Date(2014, time.July, 4, 07, 0, 0, 0, time.UTC)
	if str, bl := isUSHoliday(independence); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	labor := time.Date(2014, time.September, 1, 07, 0, 0, 0, time.UTC)
	if str, bl := isUSHoliday(labor); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	columbus := time.Date(2014, time.October, 13, 07, 0, 0, 0, time.UTC)
	if str, bl := isUSHoliday(columbus); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	halloween := time.Date(2014, time.October, 31, 07, 0, 0, 0, time.UTC)
	if str, bl := isHalloweenValentine(halloween); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	veteran := time.Date(2014, time.November, 11, 07, 0, 0, 0, time.UTC)
	if str, bl := isUSHoliday(veteran); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	thanksgiving := time.Date(2014, time.November, 27, 07, 0, 0, 0, time.UTC)
	if str, bl := isUSHoliday(thanksgiving); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	christmas := time.Date(2014, time.December, 25, 07, 0, 0, 0, time.UTC)
	if str, bl := isUSHoliday(christmas); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	christmaseve := time.Date(2014, time.December, 24, 07, 0, 0, 0, time.UTC)
	if str, bl := isSemiHoliday(christmaseve); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	newyeareve := time.Date(2014, time.December, 31, 07, 0, 0, 0, time.UTC)
	if str, bl := isSemiHoliday(newyeareve); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}
}

func weekDay(t time.Time) string {
	const layout = "Mon"
	return t.Format(layout)
}

func isWeekend(t time.Time) (string, bool) {
	day := weekDay(t)
	if day == "Sat" || day == "Sun" {
		return day, true
	}
	return day, false
}

func isUSHoliday(t time.Time) (string, bool) {
	switch t.Month() {
	case time.January:
		if t.Day() == 1 {
			return "New Year's Day", true
		}
		// third Monday
		if weekDay(t) == "Mon" &&
			t.AddDate(0, 0, -7*2).Month() == time.January &&
			t.AddDate(0, 0, -7*3).Month() == time.December {
			return "Martin Luther King, Jr. Day", true
		}
	case time.February:
		// third Monday
		if weekDay(t) == "Mon" &&
			t.AddDate(0, 0, -7*2).Month() == time.February &&
			t.AddDate(0, 0, -7*3).Month() == time.January {
			return "Washington's Birthday", true
		}
	case time.May:
		// last Monday
		if weekDay(t) == "Mon" &&
			t.AddDate(0, 0, 7*1).Month() == time.June {
			return "Memorial Day", true
		}
	case time.July:
		if t.Day() == 4 {
			return "Independence Day", true
		}
	case time.September:
		// first Monday
		if weekDay(t) == "Mon" &&
			t.AddDate(0, 0, -7*1).Month() == time.August {
			return "Labor Day", true
		}
	case time.October:
		// second Monday
		if weekDay(t) == "Mon" &&
			t.AddDate(0, 0, -7*1).Month() == time.October &&
			t.AddDate(0, 0, -7*2).Month() == time.September {
			return "Columbus Day", true
		}
	case time.November:
		if t.Day() == 11 {
			return "Veterans Day", true
		}
		// fourth Thursday
		if weekDay(t) == "Thu" &&
			t.AddDate(0, 0, -7*3).Month() == time.November &&
			t.AddDate(0, 0, -7*4).Month() == time.October {
			return "Thanksgiving Day", true
		}
	case time.December:
		if t.Day() == 25 {
			return "Christmas Day", true
		}
	}
	return "None", false
}

// isSemiHoliday returns true if it is American Holiday Eve,
// or the day after Thanksgiving day.
func isSemiHoliday(t time.Time) (string, bool) {
	// the day after Thanksgiving Day
	st, _ := isUSHoliday(t.AddDate(0, 0, -1))
	if st == "Thanksgiving Day" {
		return "Day After Thanksgiving Day", true
	}
	tm := t.AddDate(0, 0, 1)
	sh, bl := isUSHoliday(tm)
	return "Before " + sh, bl
}

// isHalloweenValentine returns true if it is Halloween or Valentine's Day.
func isHalloweenValentine(t time.Time) (string, bool) {
	if t.Month() == time.October {
		if t.Day() == 31 {
			return "Halloween", true
		}
	}
	if t.Month() == time.February {
		if t.Day() == 14 {
			return "Valentine's Day", true
		}
	}
	return "None", false
}

// isMotherFather returns true if it is Mother's day or Father's day.
func isMotherFather(t time.Time) (string, bool) {
	if t.Month() == time.May {
		// second Sunday
		if weekDay(t) == "Sun" &&
			t.AddDate(0, 0, -7*1).Month() == time.May &&
			t.AddDate(0, 0, -7*2).Month() == time.April {
			return "Mother's Day", true
		}
	}
	if t.Month() == time.June {
		// third Sunday
		if weekDay(t) == "Sun" &&
			t.AddDate(0, 0, -7*2).Month() == time.June &&
			t.AddDate(0, 0, -7*3).Month() == time.May {
			return "Father's Day", true
		}
	}
	return "None", false
}
```
---
## format[top](#time)
```
package main

import (
	"bytes"
	"fmt"
	"time"
)

func main() {
	{
		st := time.Now()
		for i := 0; i < 10000; i++ {
			buf := bytes.NewBuffer(nil)
			now := time.Now()
			ts := now.Format("2006-01-02 15:04:05")
			buf.WriteString(ts)
			ms := now.Nanosecond() / 1000
			buf.WriteString(fmt.Sprintf(".%06d", ms))
			_ = buf.String()
		}
		fmt.Println("took", time.Since(st))
	}

	{
		st := time.Now()
		for i := 0; i < 10000; i++ {
			_ = time.Now().String()[:26]
		}
		fmt.Println("took", time.Since(st))
	}
}

/*
took 8.934964ms
took 5.888227ms
*/
```
