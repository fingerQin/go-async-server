/**
 * 缓存操作（Redis）。
 * @author fingerQin
 * @date 2018-11-30
 */

package Utils

import (
	"fmt"
	"os"
)

import "github.com/go-redis/redis"

type Redis struct {
}

/**
 * 返回客户端连接。
 * @param  string  key  缓存键名。
 * @return string
 */
func (r Redis) Client() *redis.Client {
	conf := Conf{}
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
	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println("Failed to open Redis connection!")
		os.Exit(1)
	}
	return client
}
