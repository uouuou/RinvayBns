package utils

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var RawCode int       //快捷键值
var IsMos bool        //是否开启鼠标
var T string          //刺心按键
var F string          //雷电杀按键
var BTime int         //按键延迟时间
var IsColor bool      //是否开启取色
var IsCheckKey bool   //是否开启检测按键
var Tolerance float64 //容差值

// ReadIni 读取配置文件
func ReadIni() {
	cfg, err := ini.Load(GetRunPath() + "/config.ini")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}
	RawCode, _ = cfg.Section("快捷键").Key("快捷键值").Int()
	IsMos, _ = cfg.Section("快捷键").Key("使用鼠标").Bool()
	Tolerance, _ = cfg.Section("快捷键").Key("容差值").Float64()
	BTime, _ = cfg.Section("快捷键").Key("按键延时").Int()
	T = cfg.Section("快捷键").Key("刺心按键").String()
	F = cfg.Section("快捷键").Key("雷电杀按键").String()
	if Tolerance == 0 {
		Tolerance = 0.08
	}
}

// FileExist 获取某个文件是否存在
func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

// GetRunPath 获取程序运行位置
func GetRunPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	return filepath.Dir(path)
}
