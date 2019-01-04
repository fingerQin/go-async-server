package main

import (
	"Signal"
	"flag"
	"fmt"
	"os"
)

func main() {
	cc := flag.String("c", "help", "请输入指令：\n -c start 启动程序\n -c restart 重启程序\n -c stop 停止程序\n")
	flag.Parse()
	command := *cc // 指针转换为值

	switch command {
	case "start", "restart", "stop":
		Signal.Run(command)
	default:
		fmt.Println("指令有误,请执行 --help 查看帮助信息")
		os.Exit(0)
	}
}
