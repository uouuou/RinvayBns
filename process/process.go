package main

import (
	"RinvayBns/utils"
	"github.com/shirou/gopsutil/process"
	"log"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	//检测是否使用管理员权限运行
	if !utils.IsAdmin() {
		utils.ShowMessage("Golang剑灵刺客辅助", "请使用管理员权限运行")
		time.Sleep(time.Second * 5)
		return
	}
	// 获取操作系统位数
	command := utils.GetRunPath() + "\\bns_" + runtime.GOARCH + ".exe"

	for {
		isRunning, err := isProcessRunning(command)
		if err != nil {
			utils.ShowMessage("Golang剑灵刺客辅助", "检查运行状态失败")
			time.Sleep(time.Second * 5)
			continue
		}
		if !isRunning {
			cmd := exec.Command("cmd", "/c", "start", "/b", command)
			err := cmd.Start()
			if err != nil {
				utils.ShowMessage("Golang剑灵刺客辅助", "启动失败")
				time.Sleep(time.Second * 5)
				continue
			}
			log.Println("进程保护启，启动成功")
		}
		time.Sleep(time.Millisecond * 300) // 每10秒检查一次
	}
}
func isProcessRunning(name string) (bool, error) {
	processes, err := process.Processes()
	if err != nil {
		return false, err
	}
	for _, p := range processes {
		exeName, err := p.Exe()
		if err != nil {
			continue
		}
		if exeName == name {
			return true, nil
		}
	}
	return false, nil
}
