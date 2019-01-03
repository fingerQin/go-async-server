/**
 * 停止当前进程。
 * @author fingerQin
 * @date 2019-01-02
 */

package Signal

import (
	"fmt"
	"os"
	"syscall"
)

/**
 * 停止程序。
 * @param string pidSavePath PID 保存的位置。
 * @return void
 */
func Stop(pidSavePath string) {
	pid := getFileSavePID(pidSavePath)
	// [1] 检测进程是否启动。
	err := syscall.Kill(pid, 0)
	if err != nil {
		fmt.Println("服务未启动，无须退出")
		os.Exit(0)
	}
	// [2] 已启动,则发送停止信号。
	kerr := syscall.Kill(pid, syscall.SIGTERM)
	if kerr != nil {
		fmt.Println("服务停止失败")
		os.Exit(0)
	} else {
		fmt.Println("服务停止成功")
		os.Exit(0)
	}
}
