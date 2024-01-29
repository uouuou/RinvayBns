package utils

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"github.com/vcaesar/bitmap"
	"log"
	"strconv"
	"time"
)

var LJ robotgo.CBitmap     //雷决
var YsTime robotgo.CBitmap //隐身条
var B robotgo.CBitmap      //背刺
var XD robotgo.CBitmap     //毒镖
var XY robotgo.CBitmap     //吸影
var YB robotgo.CBitmap     //影匕
var ZD robotgo.CBitmap     //掷毒
var ZL robotgo.CBitmap     //掷毒雷

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
	lj, _ := robotgo.ByteToImg(Lb)
	LJ = robotgo.ImgToCBitmap(lj)
	ysTime, _ := robotgo.ByteToImg(Yb)
	YsTime = robotgo.ImgToCBitmap(ysTime)
	b, _ := robotgo.ByteToImg(Bb)
	B = robotgo.ImgToCBitmap(b)
	xd, _ := robotgo.ByteToImg(XDb)
	XD = robotgo.ImgToCBitmap(xd)
	xy, _ := robotgo.ByteToImg(XYb)
	XY = robotgo.ImgToCBitmap(xy)
	yb, _ := robotgo.ByteToImg(YBb)
	YB = robotgo.ImgToCBitmap(yb)
	zd, _ := robotgo.ByteToImg(ZDb)
	ZD = robotgo.ImgToCBitmap(zd)
	zl, _ := robotgo.ByteToImg(ZLb)
	ZL = robotgo.ImgToCBitmap(zl)
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
			_, _, ww, wh := robotgo.GetDisplayBounds(0)
			ws, hs := robotgo.GetScaleSize(0)
			fmt.Println(ww, wh)
			if ww == 0 || wh == 0 {
				log.Println("获取屏幕信息失败")
				continue
			}
			if w == 0 || h == 0 {
				log.Println("应用基础数据获取异常......")
				continue
			}
			fmt.Println(robotgo.GetDisplayBounds(0))
			value, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", float64(ww)/float64(ws)), 64)
			fmt.Println(value)
			value, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", float64(wh)/float64(hs)), 64)
			fmt.Println(value)
			bit := robotgo.CaptureScreen(int(float64(x)*value+float64(w)*value/3), int(float64(y)*value+float64(h)*value/3*2), int(float64(w)*value/3), int(float64(h)*value/3))
			fmt.Println(int(float64(x)*value), int(float64(y)*value), int(float64(w)*value), int(float64(h)*value))
			fmt.Println(int(float64(x)*value+float64(w)*value/3), int(float64(y)*value+float64(h)*value/3*2), int(float64(w)*value/3), int(float64(h)*value/3))
			fmt.Println(robotgo.GetPid())
			fmt.Println(robotgo.GetScaleSize(0))
			fmt.Println(robotgo.ScaleF(0))
			fmt.Println(robotgo.SysScale())
			fmt.Println(robotgo.IsValid(), x, y, w, h)
			//if x == 0 && y == 0 && w == 0 && h == 0 {
			//	log.Println("获取屏幕信息失败")
			//	continue
			//}
			fmt.Println(bit)
			if LJ == nil {
				log.Println("LJ is nil")
				continue
			}
			fx, fy := bitmap.Find(LJ, bit)
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
