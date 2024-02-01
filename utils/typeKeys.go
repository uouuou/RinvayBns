package utils

import (
	"github.com/go-vgo/robotgo"
	"github.com/uouuou/bitmap"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

// KeyPresser 模拟键盘输入
type KeyPresser struct {
	mu sync.Mutex
}

// NewKeyPresser 创建一个 KeyPresser
func NewKeyPresser() *KeyPresser {
	return &KeyPresser{}
}

// KeyTap 模拟键盘输入
func (k *KeyPresser) KeyTap(key string, sleep ...int) {
	k.mu.Lock()
	defer k.mu.Unlock()
	_ = robotgo.KeyDown(key)
	if sleep != nil && len(sleep) > 0 && sleep[0] > 0 {
		log.Println("按键独立 sleep:", sleep[0])
		robotgo.MilliSleep(sleep[0])
	} else {
		robotgo.MilliSleep(BTime)
	}
	_ = robotgo.KeyUp(key)
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
		isZdUse AtomicBool
		isYB    AtomicBool
		isZD    AtomicBool
		isZL    AtomicBool
		isGl    AtomicBool
		isBc    AtomicBool
		isSS    AtomicBool
		num     int
	)
	keyPresser := NewKeyPresser()
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
		if bitTop == nil || bit == nil {
			log.Println("截图失败，跳过本次循环......")
			continue
		}
		allStart := time.Now()

		// 检测是否可以雷决
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
		// 检测是否可以使用影匕
		start = time.Now()
		yb := robotgo.ImgToCBitmap(YB)
		fx, fy = bitmap.Find(yb, bit, Tolerance)
		if fx != -1 && fy != -1 {
			isYB.Set(true)
			log.Println("可以影匕ForBit")
		} else {
			isYB.Set(false)
		}
		end = time.Now()
		log.Printf("影匕检测：%v\n", end.Sub(start))
		// 检测是否可以掷毒
		start = time.Now()
		zd := robotgo.ImgToCBitmap(ZD)
		fx, fy = bitmap.Find(zd, bit, Tolerance)
		if fx != -1 && fy != -1 {
			isZD.Set(true)
			log.Println("可以掷毒ForBit")
		} else {
			isZD.Set(false)
		}
		end = time.Now()
		log.Printf("掷毒检测：%v\n", end.Sub(start))
		// 检测是否可以掷毒雷
		start = time.Now()
		zl := robotgo.ImgToCBitmap(ZL)
		fx, fy = bitmap.Find(zl, bit, Tolerance)
		if fx != -1 && fy != -1 {
			isZL.Set(true)
			log.Println("可以掷毒雷ForBit")
		} else {
			isZL.Set(false)
		}
		end = time.Now()
		log.Printf("掷毒雷检测：%v\n", end.Sub(start))
		// 检测是否可以隐身
		start = time.Now()
		ys := robotgo.ImgToCBitmap(YsTime)
		fx, fy = bitmap.Find(ys, bit, Tolerance)
		if !isYs.Get() && fx != -1 && fy != -1 {
			isYs.Set(true)
			log.Println("开始隐身")
		}
		if isYs.Get() && fx == -1 && fy == -1 {
			isYs.Set(false)
			log.Println("结束隐身")
		}
		end = time.Now()
		log.Printf("隐身检测：%v\n", end.Sub(start))
		// 检测是否可以吸影
		start = time.Now()
		xy := robotgo.ImgToCBitmap(XY)
		fx, fy = bitmap.Find(xy, bit, Tolerance)
		if fx != -1 && fy != -1 {
			isXY.Set(true)
			log.Println("检测到吸影可使用")
		} else {
			isXY.Set(false)
		}
		end = time.Now()
		log.Printf("吸影检测：%v\n", end.Sub(start))
		// 检测是否可以背刺
		start = time.Now()
		bc := robotgo.ImgToCBitmap(B)
		fx, fy = bitmap.Find(bc, bit, Tolerance)
		if fx != -1 && fy != -1 {
			isBc.Set(true)
			log.Println("检测到背刺")
		} else {
			isBc.Set(false)
		}
		end = time.Now()
		log.Printf("背刺检测：%v\n", end.Sub(start))
		// 检测是否已中毒
		start = time.Now()
		bzd := robotgo.ImgToCBitmap(BosZd)
		fx, fy = bitmap.Find(bzd, bitTop, Tolerance)
		if fx != -1 && fy != -1 {
			isBosZd.Set(true)
			log.Println("检测到BOS已中毒")
		} else {
			isBosZd.Set(false)
		}
		end = time.Now()
		log.Printf("BOS中毒检测：%v\n", end.Sub(start))
		allEnd := time.Now()
		log.Printf("检测耗时：%v\n", allEnd.Sub(allStart))
		if isYs.Get() && !isGl.Get() {
			robotgo.MilliSleep(BTime)
			keyPresser.KeyTap(robotgo.Key2)
			log.Println("挂雷成功")
			isGl.AfterFalse(time.Second * 5)
		}
		if !isBosZd.Get() {
			if isZdUse.Get() {
				isZdUse.Set(false)
			}
			if isZD.Get() && !isZdUse.Get() {
				keyPresser.KeyTap(robotgo.Key4)
				isZD.Set(false)
				isZdUse.AfterFalse(time.Second * 9)
				log.Println("掷毒启动ForBit")
			}
			if isZL.Get() && !isZdUse.Get() {
				keyPresser.KeyTap(robotgo.KeyZ)
				isZL.Set(false)
				isZdUse.AfterFalse(time.Second * 9)
				log.Println("掷毒雷启动ForBit")

			}
			if isYB.Get() && !isZdUse.Get() {
				keyPresser.KeyTap(robotgo.KeyX)
				isYB.Set(false)
				isZdUse.AfterFalse(time.Second * 9)
				log.Println("影匕启动ForBit")
			}
		}
		if !isYs.Get() && isXY.Get() {
			isXY.AfterFalse(50)
			keyPresser.KeyTap(robotgo.Key1, 80)
			log.Println("吸影成功")
			isXY.Set(false)
			robotgo.MilliSleep(BTime)
			printTime(allStart)
			continue
		}
		if !isYs.Get() && !isXY.Get() && !isSS.Get() {
			isSS.AfterFalse(50)
			_ = robotgo.KeyTap(robotgo.KeyS)
			robotgo.MilliSleep(15)
			_ = robotgo.KeyTap(robotgo.KeyS)
			_ = robotgo.KeyTap(robotgo.Key1)
			log.Println("SS1执行成功")
			printTime(allStart)
			continue
		}
		if isLj.Get() {
			keyPresser.KeyTap(robotgo.Key4, 50)
			isLj.Set(false)
			log.Println("雷决启动ForBit")
			robotgo.MilliSleep(BTime)
			printTime(allStart)
			continue
		}
		if isXd.Get() {
			keyPresser.KeyTap(robotgo.KeyX)
			isXd.Set(false)
			log.Println("毒镖启动")
		}
		if isYs.Get() && isBc.Get() && !isXY.Get() {
			keyPresser.KeyTap(robotgo.KeyR, 80)
			log.Println("背刺启动ForBit")
			robotgo.MilliSleep(BTime)
			printTime(allStart)
			continue
		}
		if !isSS.Get() {
			keyPresser.KeyTap(T)
			keyPresser.KeyTap(F)
		}
		//销毁数据
		robotgo.FreeBitmap(bit)
		robotgo.FreeBitmap(bitTop)
		printTime(allStart)
	}
}

// printTime 打印耗时
func printTime(allStart time.Time) {
	allEnd := time.Now()
	log.Printf("单次循环耗时：%v\n", allEnd.Sub(allStart))
	log.Println("----------------------")
}
