package time_test

import (
	"fmt"
	"step-wx/lib"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	time := time.Time{}

	if now.After(time) {
		fmt.Println("yes")
	}
}

func TestTime2(t *testing.T) {
	now := time.Now()
	strTime := lib.FormatTime(now)
	strTime = strings.Replace(strTime, "-", "", -1)
	strTime = strings.Replace(strTime, " ", "", -1)
	strTime = strings.Replace(strTime, ":", "", -1)
	id := strTime + strconv.FormatInt(lib.RandInt64(100000000000, 1000000000000), 10)
	fmt.Println(id)
}

func TestGetTime(t *testing.T) {
	d1, _ := time.ParseDuration("-11h")

	t1 := time.Now()
	t2 := time.Now().Add(d1)
	d2 := t2.Sub(t1)
	time.Sleep(1 * time.Microsecond)
	fmt.Println(int(d2.Hours()))
}

func test() time.Duration {
	d1, _ := time.ParseDuration("-11h")

	t1 := time.Now()
	t2 := time.Now().Add(d1)
	d2 := t2.Sub(t1)
	time.Sleep(1 * time.Microsecond)
	return d2
	//fmt.Println(int(d2.Hours()))
}

func BenchmarkGetTime(t *testing.B) {
	for i := 0; i < t.N; i++ {
		test()
	}
}

func TestIsToday(t *testing.T) {
	d, _ := time.ParseDuration("-10h")
	fmt.Println(getTime(time.Now().Add(d)))
}

func getTime(t time.Time) string {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	if t.After(today) {
		return "今天 " + t.Format("15:04")
	}
	return t.Format("2006-01-02 15:04")
}

func TestFormat(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05"))
}

func TestParse(t *testing.T) {
	str := "2018-06-22 11:10:17"
	time, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(time)
}

func TestPrint(t *testing.T) {
	var time time.Time
	fmt.Println(time)
}

func TestGetInt(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Unix())

	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	fmt.Println(time.Unix(now.Unix(), now.UnixNano()))
}

// 计算两个时间差
func TestTimeNow(tt *testing.T) {
	time1 := time.Now()
	time.Sleep(100 * time.Millisecond)
	fmt.Println(time.Since(time1))
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second * 4)
	for _ = range ticker.C {
		fmt.Println("ticked")
	}
}

func TestTimeInt(t *testing.T) {
	t1 := time.Now()
	fmt.Println(t1.Unix())
	fmt.Println(t1.UnixNano() / 1000000)

	fmt.Println(time.Unix(0, t1.UnixNano()))
}

func TestFormatToDate(t *testing.T) {
	fmt.Println(time.Now().Hour())
}
