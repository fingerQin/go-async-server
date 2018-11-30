/**
 * 配置文件读取工具类。
 * -- 通过该类实现配置文件加载一次。
 * -- 避免每次调用打开一次配置文件。
 * -- 通过使用 init 方法实现只读取一次配置文件。
 * @author fingerQin
 * @date 2018-11-30
 */

package Utils

import (
	"fmt"
	"os"
)

import "github.com/go-ini/ini"

type Conf struct {
}

/**
 * 保存配置文件。
 */
var onceConf *ini.File

/**
 * init 方法实现加载一次文件。
 */
func init() {
	var err error
	onceConf, err = ini.Load("./config/config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
}

/**
 * 返回一个 string 值。
 * @param  string  section  配置节。
 * @param  string  name     配置值。
 * @return string
 */
func (c Conf) GetString(section string, name string) string {
	return onceConf.Section(section).Key(name).String()
}

/**
 * 返回一个 int 值。
 * @param  string  section  配置节。
 * @param  string  name     配置值。
 * @return int
 */
func (c Conf) GetInt(section string, name string) int {
	return onceConf.Section(section).Key(name).MustInt()
}

/**
 * 返回一个 int 值。
 * @param  string  section  配置节。
 * @param  string  name     配置值。
 * @return bool
 */
func (c Conf) GetBool(section string, name string) bool {
	return onceConf.Section(section).Key(name).MustBool()
}
