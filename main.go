package main

import (
	"bns/utils"
	"log"
	"runtime"
	"time"
)

func main() {
	//检测是否使用管理员权限运行
	if !utils.IsAdmin() {
		utils.ShowMessage("Golang剑灵刺客辅助", "请使用管理员权限运行")
		time.Sleep(time.Second * 3)
		return
	}
	log.Println("--- Please press XButton1 (the extra mouse button). ---")
	utils.ReadIni()
	runtime.UnlockOSThread()
	utils.NewRoutine(func() {
		runtime.LockOSThread()
		utils.Systray()
	})
	//utils.Find()
	utils.NewCheck()
}
