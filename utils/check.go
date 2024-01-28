package utils

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"github.com/vcaesar/bitmap"
	"image"
	"log"
	"time"
)

var LJ image.Image     //雷决
var YsTime image.Image //隐身条
var B image.Image      //背刺
var XD image.Image     //毒镖
var XY image.Image     //吸影
var YB image.Image     //影匕
var ZD image.Image     //掷毒
var ZL image.Image     //掷毒雷

// init 初始化取色图片
func init() {
	Bb, _ := robotgo.OpenImg("./static/B.png")
	Lb, _ := robotgo.OpenImg("./static/L.png")
	XDb, _ := robotgo.OpenImg("./static/XD.png")
	Yb, _ := robotgo.OpenImg("./static/Y.png")
	XYb, _ := robotgo.OpenImg("./static/XY.png")
	YBb, _ := robotgo.OpenImg("./static/YB.png")
	ZDb, _ := robotgo.OpenImg("./static/ZD.png")
	ZLb, _ := robotgo.OpenImg("./static/ZL.png")
	LJ, _ = robotgo.ByteToImg(Lb)
	YsTime, _ = robotgo.ByteToImg(Yb)
	B, _ = robotgo.ByteToImg(Bb)
	XD, _ = robotgo.ByteToImg(XDb)
	XY, _ = robotgo.ByteToImg(XYb)
	YB, _ = robotgo.ByteToImg(YBb)
	ZD, _ = robotgo.ByteToImg(ZDb)
	ZL, _ = robotgo.ByteToImg(ZLb)
}

// NewCheck 监听鼠标/键盘事件
func NewCheck() {
	// 监听事件
	evChan := hook.Start()
	defer hook.End()
	var start bool
	var colorNum int
	fileMap := make(map[int]string)
	fileMap[1] = "L"
	fileMap[2] = "Y"
	fileMap[3] = "B"
	fileMap[4] = "BosZd"
	fileMap[5] = "XD"
	fileMap[6] = "XY"

	for ev := range evChan {
		if ev.Kind == hook.MouseHold && ev.Button == uint16(RawCode) && IsMos {
			if !start {
				log.Println("Starting typing...")
				NewRoutine(func() {
					typeKeys(&start)
				})
				start = true
			}
		}
		if ev.Kind == hook.KeyDown && ev.Rawcode == uint16(RawCode) && !IsMos {
			if !start {
				log.Println("Starting typing...")
				NewRoutine(func() {
					typeKeys(&start)
				})
				start = true
			}
		}
		if ev.Kind == hook.MouseUp && ev.Button == uint16(RawCode) && IsMos {
			log.Println("Stopping typing...")
			start = false
		}
		if ev.Kind == hook.KeyDown && ev.Rawcode == uint16(RawCode) && !IsMos {
			log.Println("Stopping typing...")
			start = false
		}
		if ev.Kind == hook.MouseDown && ev.Button == uint16(RawCode) && IsMos {
			log.Println("Stopping typing...")
			start = false
		}
		if ev.Kind == hook.KeyDown && ev.Rawcode == uint16(RawCode) && !IsMos {
			log.Println("Stopping typing...")
			start = false
		}
		if IsCheckKey {
			if ev.Kind == hook.KeyDown {
				ShowMessage("按键打印", fmt.Sprintf("按键码: %d", ev.Rawcode))
			}
			if ev.Kind == hook.MouseDown && ev.Button != 1 && ev.Button != 2 && ev.Button != 3 {
				ShowMessage("按键打印", fmt.Sprintf("鼠标按键码: %d", ev.Button))
			}
		}
		if ev.Kind == hook.MouseDown && ev.Button == 4 {
			start := time.Now()
			x, y, w, h := robotgo.GetBounds(robotgo.GetPid())
			bit := robotgo.CaptureScreen(x+w/3, y+h/3*2, w/3, h/3)
			fmt.Println(x+w/3, y+h/3*2, w/3, h/3)
			fmt.Println(robotgo.GetPid())
			//x, y, w, h = robotgo.GetDisplayBounds(0)
			fmt.Println(robotgo.GetScaleSize(0))
			fmt.Println(robotgo.ScaleF(0))
			fmt.Println(robotgo.SysScale())
			fmt.Println(robotgo.IsValid(), x, y, w, h)
			//if x == 0 && y == 0 && w == 0 && h == 0 {
			//	log.Println("获取屏幕信息失败")
			//	continue
			//}
			FindBit := robotgo.ToCBitmap(robotgo.ImgToBitmap(LJ))
			fx, fy := bitmap.Find(FindBit, bit)
			if fx != -1 && fy != -1 {
				_ = robotgo.KeyTap(robotgo.Key4)
			}
			_ = robotgo.Save(robotgo.ToImage(bit), "test.png")
			robotgo.FreeBitmap(bit)
			end := time.Now()
			log.Printf("雷决检测测试：%v  %v\n", end.Sub(start), fx != -1 && fy != -1)
		}
		if ev.Kind == hook.KeyDown && ev.Rawcode == 89 && IsColor {
			log.Println("Starting Cursor...")
			colorNum++
			x, y := robotgo.Location()
			bit := robotgo.CaptureScreen(x-15, y-15, 30, 30)
			Jt := robotgo.ToImage(bit)
			_ = robotgo.Save(Jt, fileMap[colorNum]+".png")
			robotgo.FreeBitmap(bit)
			ShowMessage("BNS", "取图片成功")
		}
	}
}
