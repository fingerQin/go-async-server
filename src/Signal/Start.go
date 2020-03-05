/**
 * 启动当前进程。
 * @author fingerQin
 * @date 2019-01-02
 */

package Signal

import (
	"Utils"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"strconv"
	"syscall"
	"time"
)

// 是否退出当前程序。
// 信号收到之后设置此值。
var is_stop bool = false

/**
 * 启动程序。
 * @param string pidSavePath PID 保存的位置。
 * @return void
 */
func Start(pidSavePath string) {
	if runtime.GOOS != "linux" {
		fmt.Println("只支持 Linux 系统")
		os.Exit(0)
	}

	if checkProgramRun(pidSavePath) == true {
		fmt.Println("程序已经启动,请勿频繁启动")
		os.Exit(0)
	}

	if os.Getppid() != 1 { // 这是父进程。
		// 将命令行参数中执行文件路径转换成可用路径
		filePath, _ := filepath.Abs(os.Args[0])
		// 将其他命令传入生成出的进程
		cmd := exec.Command(filePath, "-c", "start")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Start() //开始执行新进程，不等待新进程退出
		os.Exit(0)
	} else {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		// 通过 goroutine 方式异步处理收到的信号。
		// 然后可以平滑重启程序了
		go func() {
			<-sigs
			is_stop = true
		}()
		savePIDToFile(os.Getpid(), pidSavePath)
		timer()
		return
	}
}

// 具体的业务处理。
func ProcessLogic(qos Utils.Qos, element string) {
	qos.Add(1)
	file, err := os.OpenFile("redis.log", os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_SYNC, 0766)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer file.Close()
	file.WriteString(element)
	qos.Cut(1)
}

// 暂时用这个函数名字来调用具体的业务。
func timer() {
	// [2]
	qos := Utils.Qos{0, 0}
	redisClient := Utils.Redis{}.Client()
	key := "test_go"
	for {
		if is_stop {
			waitServiceDone()
			break
		}
		val, err := redisClient.RPop(key).Result()
		if err == nil {
			go ProcessLogic(qos, val)
		}
	}
}

// 等待正在执行的业务正确执行完成。
// 通过休眠循环实现。
func waitServiceDone() {
	qos := Utils.Qos{}
	file, err := os.OpenFile("ingCount", os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println(err)
		os.Exit(0) // 进程 PPID 写入失败
	}
	defer file.Close()
	data := strconv.Itoa(qos.IngCount)
	file.WriteString(data)

	for qos.IngCount > 0 {

		//	time.Sleep(time.Duration(1) * time.Second)
		time.Sleep(300 * time.Millisecond)
	}
}
