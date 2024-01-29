package utils

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"github.com/vcaesar/bitmap"
	"image"
	"log"
	"os"
	"strconv"
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
var BosZd image.Image  //BOS中毒

// init 初始化取色图片
func init() {
	B = loadImage("./static/B.png")
	LJ = loadImage("./static/L.png")
	XD = loadImage("./static/XD.png")
	YsTime = loadImage("./static/Y.png")
	XY = loadImage("./static/XY.png")
	YB = loadImage("./static/YB.png")
	ZD = loadImage("./static/ZD.png")
	ZL = loadImage("./static/ZL.png")
	BosZd = loadImage("./static/BosZd.png")
}

// loadImage 读取图片
func loadImage(imgPath string) image.Image {
	imageFile, err := os.Open(imgPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func(imageFile *os.File) {
		err := imageFile.Close()
		if err != nil {
			return
		}
	}(imageFile)
	img, _, _ := image.Decode(imageFile)
	return img
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
	fileMap[7] = "YB"
	fileMap[8] = "ZD"
	fileMap[9] = "ZL"
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
			bit := robotgo.CaptureScreen(int(float64(x)*Scale+float64(w)*Scale/3), int(float64(y)*Scale+float64(h)*Scale)/3*2, int(float64(w)*Scale)/3, int(float64(h)*Scale)/3)
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
			fx, fy := bitmap.Find(robotgo.ImgToCBitmap(LJ), bit)
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
