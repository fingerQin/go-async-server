/**
 * Qos - 服务质量保证。
 * - 通过对实时消息的消费数量进行数量控制。
 * @author fingerQin
 */

package Utils

import "fmt"
import "os"

import "github.com/go-ini/ini"

/**
 * 当前正在执行的消息数量。
 */
var ingCountChan chan int

/**
 * 同一时刻允许最大处理的消息数量。
 */
var maxCount int

func init() {
	cfg, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	maxCount := cfg.Section("server").Key("qos").MustInt()
	ingCountChan = make(chan int, maxCount)
}

/**
 * Qos 类。
 */
type Qos struct {
	IngCount int // 当前正在处理的消费数量记数器。
	SlowSec  int // 慢处理阀值。
}

/**
 * 扣减。
 * -- 把当前正在消费的消息数量记数器减 1.
 * @param int val 扣减值。
 * @return bool
 */
func (qos Qos) Cut(val int) bool {
	<-ingCountChan
	qos.IngCount = len(ingCountChan)
	return true
}

/**
 * 增加。
 * -- 把当前正在消费的消息数量记数器加 1.
 * @return bool
 */
func (qos Qos) Add(val int) bool {
	ingCountChan <- 1
	qos.IngCount = len(ingCountChan)
	return true
}

/**
 * 获取当前 Qos 相关状态。
 * @return map[string]int
 */
func (qos Qos) GetStatus() map[string]int {
	var numbers = make(map[string]int)
	numbers["MaxCount"] = maxCount
	numbers["IngCount"] = len(ingCountChan)
	numbers["LastInterval"] = 30 // 最后一次与上次消费之间的间隔
	numbers["SlowTimes"] = 30    // 慢处理次数
	return numbers
}
