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
	conf := Utils.Conf{}
	fmt.Println(conf.GetString("server", "qos"))

	// [2]
	qos := Utils.Qos{0, 0}

	// [3]
	redisHost := conf.GetString("redis", "default.host")
	redisPort := conf.GetString("redis", "default.port")
	redisIndex := conf.GetInt("redis", "default.index")
	redisPwd := conf.GetString("redis", "default.auth")
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
