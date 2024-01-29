package utils

import (
	"github.com/go-vgo/robotgo"
	"github.com/vcaesar/bitmap"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

// BnsKey 按键信息
type BnsKey struct {
	status bool
	key    string
	time   time.Duration
}

// NewBnsKey 创建一个 BnsKey
func NewBnsKey(key string, time time.Duration) *BnsKey {
	return &BnsKey{key: key, time: time}
}

// Down 按下 BnsKey
func (b *BnsKey) Down() {
	_ = robotgo.KeyTap(b.key)
	b.status = true
	// 计时器到期后自动将 AtomicBool 的值重置为 false
	time.AfterFunc(b.time, func() {
		b.status = false
	})
}

// Up 提前释放 BnsKey 的倒计时
func (b *BnsKey) Up() {
	b.status = false
}

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
		//isBc    AtomicBool
		num int
	)
	var wg sync.WaitGroup
	for {
		if !*status {
			return
		}
		num++
		log.Println("num:", num)
		x, y, w, h := robotgo.GetBounds(robotgo.GetPid())
		if w == 0 || h == 0 {
			log.Println("应用基础数据获取异常......")
			continue
		}
		bit := robotgo.CaptureScreen(int(float64(x)*Scale+float64(w)*Scale/3), int(float64(y)*Scale+float64(h)*Scale)/3*2, int(float64(w)*Scale)/3, int(float64(h)*Scale)/3)
		bitTop := robotgo.CaptureScreen(int(float64(x)*Scale+float64(w)*Scale/3), 0, int(float64(w)*Scale)/3, int(float64(h)*Scale)/3)
		if bitTop == nil {
			log.Println("截图失败，跳过本次循环......")
			continue
		}
		wg.Add(8)
		NewRoutine(func() {
			defer wg.Done()
			start := time.Now()

			if bit == nil {
				log.Println("截图失败，跳过本次循环......")
				return
			}
			l := robotgo.ImgToCBitmap(LJ)
			fx, fy := bitmap.Find(l, bit, Tolerance)
			if fx != -1 && fy != -1 {
				isLj.Set(true)
				log.Println("雷决启动ForBit")
			} else {
				isLj.Set(false)
			}
			end := time.Now()
			log.Printf("雷决检测：%v\n", end.Sub(start))
		})
		NewRoutine(func() {
			defer wg.Done()
			start := time.Now()
			yb := robotgo.ImgToCBitmap(YB)
			fx, fy := bitmap.Find(yb, bit, Tolerance)
			if fx != -1 && fy != -1 {
				isYB.Set(true)
				log.Println("可以影匕ForBit")
			} else {
				isYB.Set(false)
			}
			end := time.Now()
			log.Printf("影匕检测：%v\n", end.Sub(start))
		})
		NewRoutine(func() {
			defer wg.Done()
			start := time.Now()
			zd := robotgo.ImgToCBitmap(ZD)
			fx, fy := bitmap.Find(zd, bit, Tolerance)
			if fx != -1 && fy != -1 {
				isZD.Set(true)
				log.Println("可以掷毒ForBit")
			} else {
				isZD.Set(false)
			}
			end := time.Now()
			log.Printf("雷决检测：%v\n", end.Sub(start))
		})
		NewRoutine(func() {
			defer wg.Done()
			start := time.Now()
			zl := robotgo.ImgToCBitmap(ZL)
			fx, fy := bitmap.Find(zl, bit, Tolerance)
			if fx != -1 && fy != -1 {
				isZL.Set(true)
				log.Println("可以掷毒雷ForBit")
			} else {
				isZL.Set(false)
			}
			end := time.Now()
			log.Printf("掷毒雷检测：%v\n", end.Sub(start))
		})
		NewRoutine(func() {
			defer wg.Done()
			start := time.Now()
			ys := robotgo.ImgToCBitmap(YsTime)
			fx, fy := bitmap.Find(ys, bit, Tolerance)
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
		})
		NewRoutine(func() {
			defer wg.Done()
			start := time.Now()
			xd := robotgo.ImgToCBitmap(XD)
			fx, fy := bitmap.Find(xd, bit, Tolerance)
			if fx != -1 && fy != -1 {
				isXd.Set(true)
				log.Println("检测到毒镖可使用")
			} else {
				isXd.Set(false)
			}
			end := time.Now()
			log.Printf("毒镖检测：%v\n", end.Sub(start))
		})
		NewRoutine(func() {
			defer wg.Done()
			start := time.Now()
			xy := robotgo.ImgToCBitmap(XY)
			fx, fy := bitmap.Find(xy, bit, Tolerance)
			if fx != -1 && fy != -1 {
				isXY.Set(true)
				log.Println("检测到吸影可使用")
			} else {
				isXY.Set(false)
			}
			end := time.Now()
			log.Printf("吸影检测：%v\n", end.Sub(start))
		})
		NewRoutine(func() {
			defer wg.Done()
			start := time.Now()
			bzd := robotgo.ImgToCBitmap(BosZd)
			fx, fy := bitmap.Find(bzd, bitTop, Tolerance)
			if fx != -1 && fy != -1 {
				isBosZd.Set(true)
				log.Println("检测到BOS已中毒")
			} else {
				isBosZd.Set(false)
			}
			end := time.Now()
			log.Printf("BOS中毒检测：%v\n", end.Sub(start))
		})
		wg.Wait()
		if isYs.Get() && !isGl.Get() {
			robotgo.MilliSleep(BTime)
			_ = robotgo.KeyTap(robotgo.Key2)
			log.Println("挂雷成功")
			isGl.AfterFalse(time.Second * 20)
		}
		//if isYs.Get() {
		//	isBc.AfterFalse(time.Second * 5)
		//}
		//if !isBc.Get() && isYs.Get() {
		//	robotgo.MilliSleep(BTime)
		//	_ = robotgo.KeyTap(robotgo.KeyR)
		//	log.Println("背刺成功")
		//	isBc.Set(false)
		//}
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
			fx, fy := bitmap.Find(robotgo.ImgToCBitmap(B), bit, Tolerance)
			fxx, fyx := bitmap.Find(robotgo.ImgToCBitmap(XY), bit, Tolerance)
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
			robotgo.Click("right")
			robotgo.MilliSleep(BTime)
			_ = robotgo.KeyTap(F)
		} else {
			log.Printf("SS启动，不输出:%v\n", num)
			isSS.Set(false)
		}
	}
}
