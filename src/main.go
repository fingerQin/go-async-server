package main

import "fmt"
import "time"

// import "github.com/go-sql-driver/mysql"
// import "github.com/go-redis/redis"
// import "github.com/uber-go/zap"
// import "github.com/axgle/mahonia"
// import "github.com/imroc/req"
// import "github.com/360EntSecGroup-Skylar/excelize"
// import "github.com/go-gomail/gomail"
// import "github.com/dchest/captcha"

import "Utils"

func ProcessLogic(qos Utils.Qos, element string) {
	v := qos.GetStatus()
	fmt.Println(v)
	qos.Add(1)
	time.Sleep(time.Duration(1) * time.Second)
	qos.Cut(1)
}

func main() {

	// [2]
	qos := Utils.Qos{0, 0}

	key := "test_go"

	for {
		val, err := client.RPop(key).Result()
		if err == nil {
			go ProcessLogic(qos, val)
		}
	}
}
