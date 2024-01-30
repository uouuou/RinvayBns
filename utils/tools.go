package utils

import (
	"errors"
	"github.com/go-vgo/robotgo"
	"syscall"
	"unsafe"
)

// checkColor 检查指定屏幕位置的颜色
func checkColor(x, y int, color string) bool {
	// 获取指定屏幕位置的颜色
	colorNow := robotgo.GetPixelColor(x, y)
	// 比较当前颜色与目标颜色
	if color == colorNow {
		return true
	}
	return false
}

// getScreenAtCursor 获取鼠标所在屏幕索引
func getScreenAtCursor() int {
	// 获取鼠标当前位置
	mx, my := robotgo.Location()

	// 获取屏幕数量
	numScreens := robotgo.DisplaysNum()
	for i := 0; i < numScreens; i++ {
		// 获取每个屏幕的起始坐标和尺寸
		sx, sy, sw, sh := robotgo.GetDisplayBounds(i)
		// 判断鼠标位置是否在当前屏幕内
		if mx >= sx && mx < sx+sw && my >= sy && my < sy+sh {
			return i // 返回屏幕索引
		}
	}
	return -1 // 如果鼠标不在任何已知屏幕上，返回-1
}

// 一些窗口相关的函数
var (
	user32        = syscall.NewLazyDLL("user32.dll")
	getClassNameW = user32.NewProc("GetClassNameW")
)

// GetActiveWindowClassName 获取当前活动窗口的类名
func GetActiveWindowClassName() (string, error) {
	winHandle := robotgo.GetHandle()
	className := make([]uint16, 256)

	ret, _, _ := getClassNameW.Call(
		uintptr(winHandle),
		uintptr(unsafe.Pointer(&className[0])),
		uintptr(len(className)),
	)

	if ret == 0 {
		return "", errors.New("failed to get class name")
	}
	return syscall.UTF16ToString(className), nil
}
