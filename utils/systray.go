package utils

import (
	_ "embed"
	"github.com/energye/systray"
	"golang.org/x/sys/windows"
	"log"
	"os"
	"runtime/debug"
	"syscall"
	"unsafe"
)

// Systray 系统托盘
func Systray() {
	systray.Run(onReady, onExit)
}

//go:embed icon.ico
var iconData []byte

var Version string
var pidFileHandle *os.File
var sigfile = "./bns.pid"

func onReady() {
	_ = os.Remove(sigfile)
	_, err := os.Stat(sigfile)
	if err == nil {
		//pid文件存在-进程已经存在
		log.Println("PID file exist.running...")
		os.Exit(0)
	}

	// 创建当前进程的pid文件
	pidFileHandle, _ = os.OpenFile(sigfile, os.O_RDONLY|os.O_CREATE, os.ModePerm)

	systray.SetIcon(iconData)
	systray.SetTitle("BNS")
	systray.SetTooltip("BNS")
	systray.AddMenuItem("异常反馈", "异常反馈").Click(func() {
		ShowMessage("异常反馈", "请在群内直接联系我本人，我会看情况解决")
	})
	systray.AddMenuItem("检查更新", "检查更新").Click(func() {
		ShowMessage("检查更新", "没有使用线上更新功能")
	})
	systray.AddMenuItem("关于我们", "关于我们").Click(func() {
		ShowMessage("关于我们", "BNS\n版本号："+Version+"\n作者：Rinvay\n联系方式：uouuou@foxmail.com")
	})
	systray.AddMenuItem("使用说明", "使用说明").Click(func() {
		ShowMessage("使用说明", "")
	})
	ysColor := systray.AddMenuItem("隐身取色", "隐身取色")
	ysColor.Click(func() {
		if !IsColor {
			ShowMessage("隐身取色", "开启取色后鼠标移动到需要取的位置按下键盘Y键即可将颜色和XY坐标信息写入配置文件")
			ysColor.SetTitle("关闭取图")
			ysColor.SetTooltip("关闭取图")
			IsColor = true
		} else {
			ysColor.SetTitle("取图测试")
			ysColor.SetTooltip("取图测试")
			IsColor = false
		}
	})

	CheckKey := systray.AddMenuItem("按键打印", "按键打印")
	CheckKey.Click(func() {
		if !IsCheckKey {
			ShowMessage("按键打印", "开启按键打印后，按下的按键会使用弹窗的方式显示出来\n鼠标左键是1，右键是2，中键是3")
			CheckKey.SetTitle("关闭打印")
			CheckKey.SetTooltip("关闭打印")
			IsCheckKey = true
		} else {
			CheckKey.SetTitle("按键打印")
			CheckKey.SetTooltip("按键打印")
			IsCheckKey = false
		}
	})
	systray.AddSeparator()
	systray.AddMenuItem("退出程序", "退出程序").Click(func() { // 1 定义当前进程PID文件
		// 执行完毕
		err = pidFileHandle.Close()
		if err != nil {
			log.Println(err)
		}
		// 删除该文件
		err = os.Remove(sigfile)
		if err != nil {
			log.Println(err)
		}
		os.Exit(0)
	})
	OnExit(func() {
		// 执行完毕
		err = pidFileHandle.Close()
		if err != nil {
			log.Println(err)
		}
		// 删除该文件
		err = os.Remove(sigfile)
		if err != nil {
			log.Println(err)
		}
	})
	systray.SetOnClick(func(menu systray.IMenu) {
		_ = menu.ShowMenu()
	})
	systray.SetOnDClick(func(menu systray.IMenu) {
		_ = menu.ShowMenu()
	})
	systray.SetOnRClick(func(menu systray.IMenu) {
		_ = menu.ShowMenu()
	})
}

func onExit() {
	// 执行完毕
	err := pidFileHandle.Close()
	if err != nil {
		log.Println(err)
	}
	// 删除该文件
	err = os.Remove(sigfile)
	if err != nil {
		log.Println(err)
	}
}

func intPtr(n int) uintptr {
	return uintptr(n)
}
func strPtr(s string) uintptr {
	utf, _ := syscall.UTF16PtrFromString(s)
	return uintptr(unsafe.Pointer(utf))
}

func strUint16(s string) *uint16 {
	utf, _ := syscall.UTF16PtrFromString(s)
	return utf
}

// ShowMessage windows下的另一种DLL方法调用
func ShowMessage(tittle, text string) {
	if !FileExist(".env") {
		NewRoutine(func() {
			_, _ = windows.MessageBox(0, strUint16(text), strUint16(tittle), 0)
		})
	}
	return
}

// NewRoutine 采用提前recover的方式终止因为goroutine错误导致的整体崩溃(如果在For循环的入参需要var一个变量的话,需要注意变量的作用域)
func NewRoutine(f func()) {
	go func() {
		defer func() {
			if info := recover(); info != nil {
				stack := string(debug.Stack())
				log.Printf("[RunConcurrently] panic recover: %v\n%s", info, stack)
			}
		}()
		f()
	}()
}

// NewRoutineArgs 采用提前recover的方式终止因为goroutine错误导致的整体崩溃(可入参)
func NewRoutineArgs(f func(args ...any), args ...any) {
	go func() {
		defer func() {
			if info := recover(); info != nil {
				stack := string(debug.Stack())
				log.Printf("[RunConcurrently] panic recover: %v\n%s", info, stack)
			}
		}()
		f(args...)
	}()
}
