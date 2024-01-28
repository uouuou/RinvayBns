package utils

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

var exitFuncs []func()

func init() {
	go func() {
		osc := make(chan os.Signal, 1)
		signal.Notify(osc, syscall.SIGTERM, syscall.SIGINT)
		<-osc
		log.Println("收到退出信号准备退出...")
		for _, fnExit := range exitFuncs {
			fnExit()
		}
	}()
}

// OnExit 注册退出处理函数，在接收到SIGTERM或SIGINT信号时执行
func OnExit(fnExit func()) {
	exitFuncs = append(exitFuncs, fnExit)
}

// WaitExit 同步等待到退出信号后退出
func WaitExit(fnExit func()) {
	osc := make(chan os.Signal, 1)
	signal.Notify(osc, syscall.SIGTERM, syscall.SIGINT)
	<-osc
	if fnExit != nil {
		fnExit()
	}
}

// IsAdmin 检测是否使用管理员权限运行
func IsAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		return false
	}
	return true
}
