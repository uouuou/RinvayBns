package utils

import (
	"github.com/go-vgo/robotgo"
	"github.com/vcaesar/bitmap"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

// AtomicBool 原子布尔值
type AtomicBool struct {
	val int32
}

func (b *AtomicBool) Set(value bool) {
	var i int32 = 0
	if value {
		i = 1
	}
	atomic.StoreInt32(&(b.val), i)
}

func (b *AtomicBool) Get() bool {
	return atomic.LoadInt32(&(b.val)) != 0
}

// AfterFalse 以定时器设定 AtomicBool 的值，并在指定时长后自动清除
func (b *AtomicBool) AfterFalse(duration time.Duration) {
	b.Set(true)

	// 计时器到期后自动将 AtomicBool 的值重置为 false
	time.AfterFunc(duration, func() {
		b.Set(false)
	})
}

// typeKeys 模拟键盘输入
func typeKeys(status *bool) {
	if cname, err := GetActiveWindowClassName(); err == nil {
		if cname == "LaunchUnrealUWindowsClient" {

		} else {
			log.Println("NO")
			return
		}
	} else {
		log.Println(err.Error())
		return
	}
	// 一些运行参数信息
	var (
		isLj    AtomicBool
		isYs    AtomicBool
		isXY    AtomicBool
		isXd    AtomicBool
		isBosZd AtomicBool
		isYB    AtomicBool
		isZD    AtomicBool
		isZL    AtomicBool
		isSc    AtomicBool
		isSS    AtomicBool
		isGl    AtomicBool
		num     int
	)
	var wg sync.WaitGroup
	for {
		if !*status {
			return
		}
		num++
		log.Println("num:", num)
		x, y, w, h := robotgo.GetBounds(robotgo.GetPid())
		bit := robotgo.CaptureScreen(x+w/3, y+h/3*2, w/3, h/3)
		wg.Add(7)
		NewRoutineArgs(func(args ...any) {
			defer wg.Done()
			start := time.Now()
			FindBit := robotgo.ToCBitmap(robotgo.ImgToBitmap(LJ))
			fx, fy := bitmap.Find(FindBit, args[0].(robotgo.CBitmap), Tolerance)
			if fx != -1 && fy != -1 {
				isLj.Set(true)
				log.Println("雷决启动ForBit")
			} else {
				isLj.Set(false)
			}
			end := time.Now()
			log.Printf("雷决检测：%v\n", end.Sub(start))
		}, bit)
		NewRoutineArgs(func(args ...any) {
			defer wg.Done()
			start := time.Now()
			FindBit := robotgo.ToCBitmap(robotgo.ImgToBitmap(YB))
			fx, fy := bitmap.Find(FindBit, args[0].(robotgo.CBitmap), Tolerance)
			if fx != -1 && fy != -1 {
				isYB.Set(true)
				log.Println("可以影匕ForBit")
			} else {
				isYB.Set(false)
			}
			end := time.Now()
			log.Printf("影匕检测：%v\n", end.Sub(start))
		}, bit)
		NewRoutineArgs(func(args ...any) {
			defer wg.Done()
			start := time.Now()
			FindBit := robotgo.ToCBitmap(robotgo.ImgToBitmap(ZD))
			fx, fy := bitmap.Find(FindBit, args[0].(robotgo.CBitmap), Tolerance)
			if fx != -1 && fy != -1 {
				isZD.Set(true)
				log.Println("可以掷毒ForBit")
			} else {
				isZD.Set(false)
			}
			end := time.Now()
			log.Printf("雷决检测：%v\n", end.Sub(start))
		}, bit)
		NewRoutineArgs(func(args ...any) {
			defer wg.Done()
			start := time.Now()
			FindBit := robotgo.ToCBitmap(robotgo.ImgToBitmap(ZL))
			fx, fy := bitmap.Find(FindBit, args[0].(robotgo.CBitmap), Tolerance)
			if fx != -1 && fy != -1 {
				isZL.Set(true)
				log.Println("可以掷毒雷ForBit")
			} else {
				isZL.Set(false)
			}
			end := time.Now()
			log.Printf("掷毒雷检测：%v\n", end.Sub(start))
		}, bit)
		NewRoutineArgs(func(args ...any) {
			defer wg.Done()
			start := time.Now()
			FindBit := robotgo.ToCBitmap(robotgo.ImgToBitmap(YsTime))
			fx, fy := bitmap.Find(FindBit, args[0].(robotgo.CBitmap), Tolerance)
			if !isYs.Get() && fx != -1 && fy != -1 {
				isYs.Set(true)
				log.Println("开始隐身")
			}
			if isYs.Get() && fx == -1 && fy == -1 {
				isYs.Set(false)
				log.Println("结束隐身")
			}
			end := time.Now()
			log.Printf("隐身检测：%v\n", end.Sub(start))
		}, bit)
		NewRoutineArgs(func(args ...any) {
			defer wg.Done()
			start := time.Now()
			FindBit := robotgo.ToCBitmap(robotgo.ImgToBitmap(XD))
			fx, fy := bitmap.Find(FindBit, args[0].(robotgo.CBitmap), Tolerance)
			if fx != -1 && fy != -1 {
				isXd.Set(true)
				log.Println("检测到毒镖可使用")
			} else {
				isXd.Set(false)
			}
			end := time.Now()
			log.Printf("毒镖检测：%v\n", end.Sub(start))
		}, bit)
		NewRoutineArgs(func(args ...any) {
			defer wg.Done()
			start := time.Now()
			FindBitXY := robotgo.ToCBitmap(robotgo.ImgToBitmap(XY))
			fx, fy := bitmap.Find(FindBitXY, args[0].(robotgo.CBitmap), Tolerance)
			if fx != -1 && fy != -1 {
				isXY.Set(true)
				log.Println("检测到吸影可使用")
			} else {
				isXY.Set(false)
			}
			end := time.Now()
			log.Printf("吸影检测：%v\n", end.Sub(start))
		}, bit)
		wg.Wait()
		if isYs.Get() && !isGl.Get() {
			robotgo.MilliSleep(BTime)
			_ = robotgo.KeyTap(robotgo.Key2)
			log.Println("挂雷成功")
			isGl.AfterFalse(time.Second * 20)
		}
		if !isBosZd.Get() {
			if isZD.Get() && !isBosZd.Get() {
				robotgo.MilliSleep(BTime)
				_ = robotgo.KeyTap(robotgo.Key4)
				isZD.Set(false)
				isBosZd.AfterFalse(time.Second * 9)
				log.Println("掷毒启动ForBit")
			}
			if isZL.Get() && !isBosZd.Get() {
				robotgo.MilliSleep(BTime)
				_ = robotgo.KeyTap(robotgo.KeyZ)
				isZL.Set(false)
				isBosZd.AfterFalse(time.Second * 9)
				log.Println("掷毒雷启动ForBit")

			}
			if isYB.Get() && !isBosZd.Get() {
				robotgo.MilliSleep(BTime)
				_ = robotgo.KeyTap(robotgo.KeyX)
				isYB.Set(false)
				isBosZd.AfterFalse(time.Second * 9)
				log.Println("影匕启动ForBit")
			}
		}
		if !isYs.Get() && isXY.Get() {
			_ = robotgo.KeyTap(robotgo.Key1)
			log.Println("吸影成功")
			isXY.Set(false)
			robotgo.MilliSleep(BTime)
		}
		if !isYs.Get() && !isXY.Get() && isSc.Get() && !isSS.Get() {
			isSS.AfterFalse(8 * time.Second)
			_ = robotgo.KeyTap(robotgo.KeyS)
			robotgo.MilliSleep(30)
			_ = robotgo.KeyTap(robotgo.KeyS)
			robotgo.MilliSleep(50)
			_ = robotgo.KeyTap(robotgo.Key1)
			log.Println("SS1执行成功")
			robotgo.MilliSleep(BTime)
		}
		if isLj.Get() {
			_ = robotgo.KeyTap(robotgo.Key4)
			isLj.Set(false)
			log.Println("雷决启动ForBit")
			robotgo.MilliSleep(BTime)
		}
		if isXd.Get() {
			_ = robotgo.KeyTap(robotgo.KeyX)
			isXd.Set(false)
			log.Println("毒镖启动")
			robotgo.MilliSleep(BTime)
		}
		if isYs.Get() {
			start := time.Now()
			FindBit := robotgo.ToCBitmap(robotgo.ImgToBitmap(B))
			fx, fy := bitmap.Find(FindBit, bit, Tolerance)
			FindBitXY := robotgo.ToCBitmap(robotgo.ImgToBitmap(XY))
			fxx, fyx := bitmap.Find(FindBitXY, bit, Tolerance)
			if fx != -1 && fy != -1 && fxx == -1 && fyx == -1 {
				_ = robotgo.KeyTap(robotgo.KeyR)
				isSc.Set(true)
				log.Println("背刺启动ForBit")
				robotgo.MilliSleep(BTime)
			}
			end := time.Now()
			log.Printf("背刺检测：%v\n", end.Sub(start))
		}
		if !isSS.Get() {
			log.Printf("SS未启动，正常输出:%v\n", num)
			_ = robotgo.KeyTap(T)
			robotgo.MilliSleep(BTime)
			_ = robotgo.KeyTap(F)
			robotgo.MilliSleep(BTime)
		} else {
			log.Printf("SS启动，不输出:%v\n", num)
			isSS.Set(false)
		}
		robotgo.FreeBitmap(bit)
	}
}
