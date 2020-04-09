package time_test

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	fmt.Println(time.Unix(1551427400, 0))
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
	a := time.Now().UnixNano()
	time.Sleep(3 * time.Second)
	b := time.Now().UnixNano()
	fmt.Println(b - a)
}

func ReturnInt() []int {
	return []int{}
}

func TestFormatToDate(t *testing.T) {
	fmt.Println(0x1A)

}

func dive100(i int64) string {
	z := i / 100
	y := i % 100
	if y < 10 {
		return fmt.Sprintf("%d.0%d", z, y)
	} else {
		return fmt.Sprintf("%d.%d", z, y)
	}
}
