/**
 * 当前业务各种信号处理。
 * -- 启动。
 * -- 重启。
 * -- 退出。
 * @author fingerQin
 * @date 2019-01-02
 */

package Signal

import (
	"Utils"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"syscall"
)

const START = "start"
const STOP = "stop"
const RESTART = "restart"

func Run(command string) {
	conf := Utils.Conf{}
	pidSavePath := conf.GetString("server", "pid")
	switch command {
	case START:
		Start(pidSavePath)
		break
	case STOP:
		Stop(pidSavePath)
		break
	case RESTART:
		Restart(pidSavePath)
		break
	}
}

// 检验当前程序是否处于运行中。
func checkProgramRun(pidSavePath string) bool {
	pid := getFileSavePID(pidSavePath)
	if pid == 0 {
		return false
	}
	err := syscall.Kill(pid, 0)
	if err != nil {
		return false
	} else {
		return true
	}
}

// 获取保存在文件当中的父 ID
func getFileSavePID(pidSavePath string) int {
	b, err := ioutil.ReadFile(pidSavePath)
	if err != nil {
		return 0
	} else {
		str := string(b)
		ppid, err := strconv.Atoi(str)
		if err != nil {
			return 0
		} else {
			return ppid
		}
	}
}

// 保存父进程 ID 到文件
func savePIDToFile(pid int, pidSavePath string) {
	file, err := os.OpenFile(pidSavePath, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println(err)
		os.Exit(0) // 进程 PPID 写入失败
	}
	defer file.Close()
	data := strconv.Itoa(pid)
	file.WriteString(data)
}
