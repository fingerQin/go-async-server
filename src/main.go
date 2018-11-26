package main

import "fmt"
import "os"
import "time"

// import "github.com/go-sql-driver/mysql"
// import "github.com/go-redis/redis"
// import "github.com/uber-go/zap"
// import "github.com/axgle/mahonia"
// import "github.com/imroc/req"
// import "github.com/360EntSecGroup-Skylar/excelize"
// import "github.com/go-gomail/gomail"
// import "github.com/dchest/captcha"

import "github.com/go-ini/ini"
import "github.com/go-redis/redis"

import "Utils"

func ProcessLogic(qos Utils.Qos, element string) {
	v := qos.GetStatus()
	fmt.Println(v)
	qos.Add(1)
	time.Sleep(time.Duration(1) * time.Second)
	qos.Cut(1)
}

func main() {
	// [1]
	cfg, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	fmt.Println(cfg.Section("server").Key("qos").MustInt())

	// [2]
	qos := Utils.Qos{0, 0}

	// [3]
	redisHost := cfg.Section("redis").Key("default.host").String()
	redisPort := cfg.Section("redis").Key("default.port").String()
	redisIndex := cfg.Section("redis").Key("default.index").MustInt()
	redisPwd := cfg.Section("redis").Key("default.auth").String()
	addr := redisHost + ":" + redisPort
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: redisPwd,   // no password set
		DB:       redisIndex, // use default DB
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	key := "test_go"

	for {
		val, err := client.RPop(key).Result()
		if err == nil {
			go ProcessLogic(qos, val)
		}
	}

	fmt.Printf("Hello, world or 你好，世界 or καλημ ́ρα κóσμ or こんにちはせかい\n")
}
