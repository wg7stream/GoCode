package main
import(
	"fmt"
	"time"
)

func main() {
	// 1. 获取当前时间
	now := time.Now()
	fmt.Printf("now = %v now type = %T", now, now)

	// 2. 通过now可以获取到年月日，时分秒
	//fmt.Printf("年 = %v\n", now.Year())
	//fmt.Printf("月 = %v\n", now.Month())

	// 3.格式化日期时间
	//fmt.Printf("当前年月日 %d-%d-5d %d:%d:%d \n", now.Year(),
	//now.Month(), now.Day(), now.Hour(). now.Minute(), now.Second())

	//dateStr := fmt.Sprintf("当前年月日 %d-%d-5d %d:%d:%d \n",now.Year(),
	//now.Month(), now.Day(), now.Hour(). now.Minute(), now.Second())

	//fmt.Printf("dateStr = %v\n", dateStr)

	// 格式化日期的第二种方式
	// 必须用这个时间
	fmt.Printf(now.Format("2006-01-02 15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()

	// 每个1秒打印一个数字
	// 每个0.1秒打印一个数字
	i := 0
	for {
		i ++
		fmt.Println(i)
		// 休眠
		//time.Sleep(time.Second)
		time.Sleep(time.Millisecond * 100)
		if i == 100{
			break
		}
	}

	// Unix和UnixNano的使用
	fmt.Printf("unix时间戳 = %v  unixnano时间戳 = %v\n", now.Unix(), now.UnixNano())
}